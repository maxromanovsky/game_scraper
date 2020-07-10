// Code is a modified version of the quickstart code:
// https://developers.google.com/gmail/api/quickstart/go

package scrape

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

const user = "me"
const maxResults = 100

type GMailScraper struct {
	client *gmail.Service
	wg     sync.WaitGroup
	delay  time.Duration
}

func NewMailScraper(delay time.Duration) GMailScraper {
	return GMailScraper{client: getGmailClient(), delay: delay}
}

func (s *GMailScraper) Scrape(filters []EmailFilter, messages chan<- EmailMessage) {
	s.wg = sync.WaitGroup{}
	for _, f := range filters {
		filter := createFilter(f)
		pageToken := ""
		i := 1

		for {
			r, err := s.client.Users.Messages.List(user).Q(filter).MaxResults(maxResults).PageToken(pageToken).Do()
			if err != nil {
				log.Fatalf("Unable to retrieve messages: %v", err)
			}

			pageToken = r.NextPageToken
			log.Printf("page #%d, filter: '%s', nextPageToken '%s'", i, filter, pageToken)
			i++

			if len(r.Messages) == 0 {
				// No messages
				break
			}

			for _, m := range r.Messages {
				s.wg.Add(1)
				go s.getMessage(m.Id, messages)
				// Preventing googleapi: Error 429: Too many concurrent requests for user, rateLimitExceeded
				// Better approach would be exponential backoff
				// https://developers.google.com/gmail/api/v1/reference/quota#concurrent_requests
				time.Sleep(s.delay)
			}

			if pageToken == "" {
				// Last page
				break
			}
		}
	}

	s.wg.Wait()
	close(messages)
}

func (s *GMailScraper) getMessage(id string, messages chan<- EmailMessage) {
	defer s.wg.Done()

	full, err := s.client.Users.Messages.Get(user, id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve message %s: %v", id, err)
	}

	raw, err := s.client.Users.Messages.Get(user, id).Format("raw").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve message %s: %v", id, err)
	}

	if msg := toEmailMessage(full, raw); msg != nil {
		messages <- *msg
	}
}

func createFilter(f EmailFilter) string {
	var filters []string
	if len(f.From) > 0 {
		//Assumption: all values are not null
		filters = append(filters, fmt.Sprintf("from:(%s)", strings.Join(f.From, " OR ")))
	}
	return strings.Join(filters, " AND ")
}

func toEmailMessage(full, raw *gmail.Message) *EmailMessage {
	m := EmailMessage{Id: full.Id, Raw: raw.Raw, Source: GOG}

	m.DateReceived = time.Unix(0, full.InternalDate*int64(time.Millisecond))

	extractMessageHeaders(full, &m)
	m.Parts = append(m.Parts, *convertBodyPart(full.Payload))
	for _, p := range full.Payload.Parts {
		m.Parts = append(m.Parts, *convertBodyPart(p))
	}
	return &m
}

func extractMessageHeaders(full *gmail.Message, m *EmailMessage) {
	for _, h := range full.Payload.Headers {
		switch h.Name {
		case "From":
			m.From = h.Value
		case "To":
			m.To = h.Value
		case "Subject":
			m.Subject = h.Value
		}
	}
}

func convertBodyPart(mp *gmail.MessagePart) *BodyPart {
	part := BodyPart{PartId: mp.PartId, MimeType: mp.MimeType, Filename: mp.Filename}
	part.Headers = convertBodyPartHeaders(mp.Headers)

	body, err := base64.URLEncoding.DecodeString(mp.Body.Data)
	if err != nil {
		log.Fatalf("Base64 decode error: %v", err)
	}
	part.Body = string(body)

	return &part
}

func convertBodyPartHeaders(src []*gmail.MessagePartHeader) map[string]string {
	headers := make(map[string]string)
	for _, h := range src {
		headers[h.Name] = h.Value
	}
	return headers
}

func getGmailClient() *gmail.Service {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	ctx := context.Background()
	tokenOpt := getTokenClientOption(ctx, config)
	srv, err := gmail.NewService(ctx, tokenOpt)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}
	return srv
}

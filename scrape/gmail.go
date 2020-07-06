// Code is a modified version of the quickstart code:
// https://developers.google.com/gmail/api/quickstart/go

package scrape

import (
	"encoding/base64"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

const user = "me"

type GMailScraper struct {
	client *gmail.Service
	wg     sync.WaitGroup
}

func NewMailScraper() GMailScraper {
	return GMailScraper{getGmailClient(), sync.WaitGroup{}}
}

func (s *GMailScraper) Scrape(messages chan<- EmailMessage) {
	//from:(gog.com) -newsletter@email.gog.com -newsletter@email2.gog.com -do-not-reply@email.gog.com -do_not_reply@gog.com
	//from:(do_not_reply@gog.com OR do-not-reply@email.gog.com)
	//"GOG.com Team" <do_not_reply@gog.com>
	//"GOG.com Team" <do-not-reply@email.gog.com>
	filter := "from:(do_not_reply@gog.com OR do-not-reply@email.gog.com)"
	r, err := s.client.Users.Messages.List(user).Q(filter).MaxResults(1).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}

	for _, m := range r.Messages {
		s.wg.Add(1)
		go s.getMessage(m.Id, messages)
	}

	s.wg.Wait()
	close(messages)
}

func (s *GMailScraper) getMessage(id string, messages chan<- EmailMessage) {
	defer s.wg.Done()
	m, err := s.client.Users.Messages.Get(user, id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve message %s: %v", id, err)
	}

	if msg, ok := toEmailMessage(m); ok {
		messages <- *msg
	}
}

func toEmailMessage(src *gmail.Message) (*EmailMessage, bool) {
	m := EmailMessage{Id: src.Id, Raw: src.Raw}

	m.DateReceived = time.Unix(0, src.InternalDate*int64(time.Millisecond))

	for _, h := range src.Payload.Headers {
		switch h.Name {
		case "From":
			m.From = h.Value
		case "To":
			m.To = h.Value
		case "Subject":
			m.Subject = h.Value
		}
	}

	body, err := base64.URLEncoding.DecodeString(src.Payload.Body.Data)
	if err != nil {
		log.Printf("Base64 decode error: %v", err)
		return nil, false
	}
	m.Body = string(body)

	//todo check
	log.Println(src.Id, src.Raw, src.Payload.Body.Data)

	//todo error handling
	return &m, true
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

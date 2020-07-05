// Code is a modified version of the quickstart code:
// https://developers.google.com/gmail/api/quickstart/go

package scrape

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
	"sync"
)

const user = "me"

type GMailScraper struct {
	client *gmail.Service
	wg     sync.WaitGroup
}

func NewMailScraper() GMailScraper {
	return GMailScraper{getGmailClient(), sync.WaitGroup{}}
}

func (s *GMailScraper) Scrape(messages chan<- Message) {
	r, err := s.client.Users.Messages.List(user).MaxResults(5).Do()
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

func (s *GMailScraper) getMessage(id string, messages chan<- Message) {
	defer s.wg.Done()
	m, err := s.client.Users.Messages.Get(user, id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve message %s: %v", id, err)
	}
	messages <- Message{Id: m.Id}
	//for _, h := range m.Payload.Headers {
	//	log.Println(h.Name, h.Value)
	//}
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

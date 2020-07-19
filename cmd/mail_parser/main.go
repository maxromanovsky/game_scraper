package main

import (
	"github.com/maxromanovsky/game_scraper/domain/entity"
	"github.com/maxromanovsky/game_scraper/domain/repository/avro"
	"github.com/maxromanovsky/game_scraper/filter"
	"log"
)

func main() {
	messages := make(chan *entity.EmailMessage)
	done := make(chan struct{})

	repo := avro.NewEmailMessageRepository("email_messages.avro") //todo: configurable via CLI
	go repo.Load(messages, nil)

	filters := filter.NewChain(true, filter.NewGogSubjectRegex())

	go filterMessages(filters, messages, done)

	<-done
}

func filterMessages(filter filter.Filter, messages <-chan *entity.EmailMessage, done chan<- struct{}) {
	i := 1
	for m := range messages {
		if res, err := filter.Filter(m); err == nil && res {
			log.Printf("%d -> %s, %s, %s", i, m.Id, m.DateReceived, m.Subject)
		}
		i++
	}
	close(done)
}

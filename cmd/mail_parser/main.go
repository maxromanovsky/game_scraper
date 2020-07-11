package main

import (
	"github.com/maxromanovsky/game_scraper/domain/entity"
	"github.com/maxromanovsky/game_scraper/domain/repository/avro"
	"log"
)

func main() {
	messages := make(chan *entity.EmailMessage)
	done := make(chan struct{})

	repo := avro.NewEmailMessageRepository("email_messages.avro") //todo: configurable via CLI
	go repo.Load(messages, nil)

	go printMessages(messages, done)

	<-done
}

func printMessages(messages <-chan *entity.EmailMessage, done chan<- struct{}) {
	i := 1
	for m := range messages {
		log.Printf("%d -> %s, %s, %s", i, m.Id, m.DateReceived, m.Subject)
		i++
	}
	close(done)
}

package main

import (
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/maxromanovsky/game_scraper/domain/entity"
	"github.com/maxromanovsky/game_scraper/domain/repository/avro"
	"github.com/maxromanovsky/game_scraper/scrape"
	"log"
	"os"
	"time"
)

func main() {
	scraper := scrape.NewMailScraper(20 * time.Millisecond) //todo: make configurable through CLI param
	messages := make(chan entity.EmailMessage)

	done := make(chan struct{})

	go func() {
		defer close(done)

		fileWriter, err := os.Create("email_messages.avro")
		if err != nil {
			log.Fatalf("Error opening file writer: %v", err)
		}
		defer fileWriter.Close()

		containerWriter, err := avro.NewEmailMessageSchemaWriter(fileWriter, container.Null, 10)
		if err != nil {
			log.Fatalf("Error opening container writer: %v", err)
		}

		i := 1
		for m := range messages {
			log.Printf("%d -> %s, %s", i, m.From, m.Id)
			i++
			err = containerWriter.WriteRecord(avro.ToEmailMessageSchema(&m))
			if err != nil {
				log.Fatalf("Error writing record to file: %v", err)
			}
		}

		err = containerWriter.Flush()
		if err != nil {
			log.Fatalf("Error flushing last block to file: %v", err)
		}
	}()
	scraper.Scrape([]scrape.EmailFilter{scrape.GmailEmailFilter}, messages)
	<-done
}

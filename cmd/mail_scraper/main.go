package main

import (
	"github.com/maxromanovsky/game_scraper/scrape"
	"log"
)

func main() {
	scraper := scrape.NewMailScraper()
	messages := make(chan scrape.EmailMessage)

	done := make(chan struct{})

	go func() {
		for m := range messages {
			log.Println(m)
		}
		close(done)
	}()
	scraper.Scrape(messages)
	<-done
}

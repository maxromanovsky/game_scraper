package main

import (
	"github.com/maxromanovsky/game_scraper/scrape"
	"log"
	"time"
)

func main() {
	scraper := scrape.NewMailScraper(20 * time.Millisecond) //todo: make configurable through CLI param
	messages := make(chan scrape.EmailMessage)

	done := make(chan struct{})

	go func() {
		i := 1
		for m := range messages {
			log.Printf("%d -> %s, %s", i, m.From, m.Id)
			i++
		}
		close(done)
	}()
	scraper.Scrape([]scrape.EmailFilter{scrape.GmailEmailFilter}, messages)
	<-done
}

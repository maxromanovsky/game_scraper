package main

import (
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/maxromanovsky/game_scraper/domain/entity"
	"github.com/maxromanovsky/game_scraper/domain/repository/avro"
	"github.com/maxromanovsky/game_scraper/scrape"
	"time"
)

func main() {
	messages := make(chan *entity.EmailMessage)
	done := make(chan struct{})

	scraper := scrape.NewMailScraper(20 * time.Millisecond) //todo: make configurable through CLI param
	go scraper.Scrape([]scrape.EmailFilter{scrape.GmailEmailFilter}, messages)

	repo := avro.NewEmailMessageRepository("email_messages.avro") //todo configurable via CLI
	go repo.Save(container.Null, 10, messages, func() { close(done) })

	<-done
}

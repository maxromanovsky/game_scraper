package main

import (
	"fmt"
	"github.com/maxromanovsky/game_scraper/scrape"
)

func main() {
	fmt.Println("Test")
	scraper := scrape.MailScraper{}
	scraper.Scrape()
}

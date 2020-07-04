package main

import (
	"github.com/maxromanovsky/game_scraper/scrape"
)

func main() {
	scraper := scrape.NewMailScraper()
	scraper.Scrape()
}

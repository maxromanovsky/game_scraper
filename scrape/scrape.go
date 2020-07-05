package scrape

import "time"

type Scraper interface {
	Scrape(messages chan<- Message)
}

type Message struct {
	Id, From, Body string
	DateReceived   time.Time
}

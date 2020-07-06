package scrape

import "time"

type Scraper interface {
	Scrape(messages chan<- EmailMessage)
}

type EmailMessage struct {
	Id, From, To, Subject, Body, Raw string
	DateReceived                     time.Time
}

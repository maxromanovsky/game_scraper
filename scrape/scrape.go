package scrape

import (
	"time"
)

type Scraper interface {
	Scrape(messages chan<- EmailMessage)
}

type EmailMessage struct {
	Id, From, To, Subject, Raw string
	Source                     EmailSource
	DateReceived               time.Time
	Parts                      []BodyPart
}

type BodyPart struct {
	PartId, MimeType, Filename, Body string
	Headers                          map[string]string //Assumption: no headers with duplicate names
}

type EmailSource int

const (
	GOG EmailSource = iota
	AppleAppStore
)

package scrape

import (
	"time"
)

type Scraper interface {
	Scrape(filters []EmailFilter, messages chan<- EmailMessage)
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

type EmailFilter struct {
	From []string
}

//from:(gog.com) -newsletter@email.gog.com -newsletter@email2.gog.com -do-not-reply@email.gog.com -do_not_reply@gog.com
//from:(do_not_reply@gog.com OR do-not-reply@email.gog.com)
var GmailEmailFilter = EmailFilter{From: []string{"do_not_reply@gog.com", "do-not-reply@email.gog.com"}}

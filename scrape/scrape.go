package scrape

import "github.com/maxromanovsky/game_scraper/domain/entity"

type Scraper interface {
	Scrape(filters []EmailFilter, messages chan<- entity.EmailMessage)
}

type EmailFilter struct {
	From []string
}

//from:(gog.com) -newsletter@email.gog.com -newsletter@email2.gog.com -do-not-reply@email.gog.com -do_not_reply@gog.com
//from:(do_not_reply@gog.com OR do-not-reply@email.gog.com)
var GmailEmailFilter = EmailFilter{From: []string{"do_not_reply@gog.com", "do-not-reply@email.gog.com"}}

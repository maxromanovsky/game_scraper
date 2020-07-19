package filter

import (
	"fmt"
	"github.com/maxromanovsky/game_scraper/domain/entity"
	"log"
	"regexp"
)

var gogSubjects []*regexp.Regexp

// Skipping the following subjects:
// Items on your wishlist are now discounted!
// Produkty z Twojej listy życzeń zostały przecenione!
// A game you’ve been waiting for is now available on GOG.com
// A gift for you!
// Otrzymujesz prezent!
// Hey, your free game has arrived!

// Currently it does make sense to keep it global, but "canonically" it might belong to struct
func init() {
	var patterns = [...]string{
		`Your order #([a-z0-9]+) is complete!`,
		`Twoje zamówienie nr ([a-z0-9]+) zostało zrealizowane!`,
		`Free items added to your GOG\.com library\.`,
		`Do Twojej biblioteki GOG\.com dodano darmowe produkty\.`,
	}

	gogSubjects = make([]*regexp.Regexp, len(patterns))

	for i, p := range patterns {
		r, err := regexp.Compile(fmt.Sprintf("^%s$", p))
		if err != nil {
			log.Fatalf("Can't compile Gog regexp '%s'", p)
		}
		gogSubjects[i] = r
	}
}

type GogSubjectRegex struct {
}

func NewGogSubjectRegex() *GogSubjectRegex {
	return &GogSubjectRegex{}
}

func (s *GogSubjectRegex) IsSupported(message *entity.EmailMessage) bool {
	return message.Source == entity.GOG
}

func (s *GogSubjectRegex) Filter(message *entity.EmailMessage) (bool, error) {
	for _, r := range gogSubjects {
		if r.MatchString(message.Subject) {
			return true, nil
		}
	}

	return false, nil
}

package filter

import "github.com/maxromanovsky/game_scraper/domain/entity"

type Filter interface {
	// True if message should be retained, False if message should be filtered out
	Filter(message *entity.EmailMessage) (bool, error)
}

type Chain struct {
	matchAny bool
	filters  []Filter
}

func NewChain(matchAny bool, filters ...Filter) *Chain {
	return &Chain{matchAny: matchAny, filters: filters}
}

// Lazy evaluation is used. Not all filters are executed if conditions are met.
func (cf *Chain) Filter(message *entity.EmailMessage) (bool, error) {
	res := true
	var err error = nil
	for _, f := range cf.filters {
		res, err = f.Filter(message)
		switch {
		case err != nil:
			return false, err
		case res == true && cf.matchAny:
			return true, nil
		case res == false && !cf.matchAny:
			return false, nil
		}
	}
	return res, nil
}

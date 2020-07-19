package filter

import "github.com/maxromanovsky/game_scraper/domain/entity"

type Filter interface {
	// Checks if message is supported. Typically checks the source
	IsSupported(message *entity.EmailMessage) bool

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

func (c *Chain) IsSupported(message *entity.EmailMessage) bool {
	// All messages are supported. If none underlying filter supports the message, then it will be filtered out
	return true
}

// Lazy evaluation is used. Not all filters are executed if conditions are met.
func (c *Chain) Filter(message *entity.EmailMessage) (bool, error) {
	res := false
	var err error = nil
	for _, f := range c.filters {
		if !f.IsSupported(message) {
			continue
		}

		res, err = f.Filter(message)
		switch {
		case err != nil:
			return false, err
		case res == true && c.matchAny:
			return true, nil
		case res == false && !c.matchAny:
			return false, nil
		}
	}
	return res, nil
}

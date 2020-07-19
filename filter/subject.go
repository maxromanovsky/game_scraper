package filter

import (
	"github.com/maxromanovsky/game_scraper/domain/entity"
)

type SubjectRegexFilter struct {
}

func NewSubjectRegexFilter() *SubjectRegexFilter {
	return &SubjectRegexFilter{}
}

func (s *SubjectRegexFilter) Filter(message *entity.EmailMessage) (bool, error) {
	panic("implement me")
}

package filter

import (
	"errors"
	"github.com/maxromanovsky/game_scraper/domain/entity"
	"testing"
)

type Stub struct {
	res bool
	err error
}

func (s *Stub) Filter(message *entity.EmailMessage) (bool, error) {
	return s.res, s.err
}

func NewStub(res bool, err error) *Stub {
	return &Stub{res: res, err: err}
}

func TestChainTable(t *testing.T) {
	dummyMessage := entity.EmailMessage{Id: "dummy"}
	stubTrue := NewStub(true, nil)
	stubFalse := NewStub(false, nil)
	stubErrTrue := NewStub(true, errors.New("Foo"))
	stubErrFalse := NewStub(false, errors.New("Foo"))

	tests := []struct {
		name     string
		matchAny bool
		filters  []Filter
		message  *entity.EmailMessage
		res      bool
		err      bool
	}{
		{"nil filters nil message match all", false, nil, nil, true, false},
		{"nil filters nil message match any", true, nil, nil, true, false},
		{"empty filters nil message match all", false, []Filter{}, nil, true, false},
		{"empty filters nil message match any", true, []Filter{}, nil, true, false},
		{"nil filters dummy message match all", false, nil, &dummyMessage, true, false},
		{"nil filters dummy message match any", true, nil, &dummyMessage, true, false},
		{"empty filters dummy message match all", false, []Filter{}, &dummyMessage, true, false},
		{"empty filters dummy message match any", true, []Filter{}, &dummyMessage, true, false},
		{"single err filter, message match any - all true", true, []Filter{stubErrTrue}, &dummyMessage, false, true},
		{"single err filter, message match any - all false", true, []Filter{stubErrFalse}, &dummyMessage, false, true},
		{"single filter, message match any - all true", true, []Filter{stubTrue}, &dummyMessage, true, false},
		{"single filter, message match any - all false", true, []Filter{stubFalse}, &dummyMessage, false, false},
		{"multi err filter, message match any - 1st true 2nd err", true, []Filter{stubTrue, stubErrTrue}, &dummyMessage, true, false},
		{"multi err filter, message match any - 1st false 2nd err", true, []Filter{stubFalse, stubErrTrue}, &dummyMessage, false, true},
		{"multi err filter, message match all - 1st true 2nd err", false, []Filter{stubTrue, stubErrTrue}, &dummyMessage, false, true},
		{"multi err filter, message match all - 1st false 2nd err", false, []Filter{stubFalse, stubErrFalse}, &dummyMessage, false, false},
		{"multi err filter, message match any - 1st true err", true, []Filter{stubErrTrue, stubTrue}, &dummyMessage, false, true},
		{"multi err filter, message match any - 1st err", true, []Filter{stubErrFalse, stubTrue}, &dummyMessage, false, true},
		{"multi filter, message match any - some true", true, []Filter{stubFalse, stubTrue, stubFalse}, &dummyMessage, true, false},
		{"multi filter, message match any - all false", true, []Filter{stubFalse, stubFalse}, &dummyMessage, false, false},
		{"multi filter, message match all - some true", false, []Filter{stubFalse, stubTrue, stubFalse}, &dummyMessage, false, false},
		{"multi filter, message match all - all false", false, []Filter{stubFalse, stubFalse}, &dummyMessage, false, false},
		{"multi filter, message match all - all true", false, []Filter{stubTrue, stubTrue}, &dummyMessage, true, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewChain(test.matchAny, test.filters...)
			res, err := c.Filter(test.message)

			if err != nil {
				if res == true {
					t.Errorf(`Res is true besides the error %v`, err)
				}

				if !test.err {
					t.Errorf(`Error is not nil: %v`, err)
				}
				return
			}

			if test.err {
				t.Error(`Expected error was not thrown`)
			}

			if res != test.res {
				t.Errorf(`Expected: %v, Actual: %v`, test.res, res)
			}
		})
	}
}

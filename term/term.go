package term

import "time"

const (
	Daily   = "d"
	Monthly = "m"
	Yearly  = "y"
)

var FileTerms []string = []string{Daily, Monthly, Yearly}

type NextTermer interface {
	NextTerm() time.Time
}

type timeTerm struct {
	term   time.Time
	period string
}

func (t *timeTerm) NextTerm() time.Time {
	t.term = subtractDate(t.term, t.period)
	return t.term
}

func subtractDate(d time.Time, term string) time.Time {
	switch term {
	case Daily:
		return d.AddDate(0, 0, -1)
	case Monthly:
		return d.AddDate(0, -1, 0)
	case Yearly:
		return d.AddDate(-1, 0, 0)
	}
	return d
}

func NewTimeTerm(t time.Time, p string) NextTermer {
	tt := &timeTerm{term: t, period: p}
	return tt
}

package assessment

import (
	"fmt"
	"time"
)

/**
 * Constants
 */
const TimeLayout = "2006-Jan-02"

const (
	Male int = iota
	Female
)

/**
 * Interfaces
 */

// Measurer is the interface to be implemented by any struct that should convey
// a measurement in physical assessment.
type Measurer interface {
	// Proxy to retrieve the measurement name.
	GetName() string
	// Result shows the meaningful information available from the measurement
	// itself. It can also show an error with the measurement process is
	// somehow invalid.
	Result() ([]string, error)
}

/**
 * Assessment
 */

// Assessment represents a collection of measurements made in a given date.
// Also knows how to represent the collection of measurements through Measurer
// interface.
type Assessment struct {
	Date     time.Time
	Measures []Measurer
}

// NewAssessment returns an Assessment instance
func NewAssessment(date string, measurements ...Measurer) (*Assessment, error) {
	a := &Assessment{Measures: measurements}
	d, err := time.Parse(TimeLayout, date)
	if err != nil {
		return a, err
	}
	a.Date = d
	return a, nil
}

func (a *Assessment) String() string {
	return fmt.Sprintf("Assessment made in %s", a.Date.Format(TimeLayout))
}

func (a *Assessment) GetName() string {
	return a.String()
}

// Result aggregates all measures results into one representation. If one
// measure can't has error, then the given error is returned.
func (a *Assessment) Result() ([]string, error) {
	accum := []string{}
	for _, measure := range a.Measures {
		rs, err := measure.Result()
		if err != nil {
			return []string{}, fmt.Errorf("Measure _%s_ failed: %s", measure.GetName(), err)
		}
		accum = append(accum, measure.GetName())
		accum = append(accum, rs...)
	}
	return accum, nil
}

// AddMeasure allow to add a new measure to the ones available in a given
// assessment.
func (a *Assessment) AddMeasure(m Measurer) {
	a.Measures = append(a.Measures, m)
}

/**
 * Person
 */

// Person is the common information from the individual being measured
type Person struct {
	FullName string
	Birthday time.Time
	Gender   int
}

func NewPerson(fullName string, birth string, gender int) (*Person, error) {
	p := &Person{}
	b, err := time.Parse(TimeLayout, birth)
	if err != nil {
		return p, err
	}
	p.FullName = fullName
	p.Birthday = b.UTC()
	p.Gender = gender
	return p, nil
}

func (p *Person) String() string {
	return fmt.Sprintf("Name: %s\nGender: %s\nAge: %.0f", p.FullName, p.genderRepr(), p.Age())
}

func (p *Person) Age() float64 {
	return p.AgeFromDate(time.Now().UTC())
}

func (p *Person) AgeInMonths() float64 {
	return p.AgeInMonthsFromDate(time.Now().UTC())
}

func (p *Person) AgeFromDate(t time.Time) float64 {
	age := t.Year() - p.Birthday.Year()
	if t.Month() < p.Birthday.Month() || (t.Month() == p.Birthday.Month() && t.Day() < p.Birthday.Day()) {
		age -= 1
	}
	return float64(age)
}

func (p *Person) AgeInMonthsFromDate(t time.Time) float64 {
	return elapsedFromDateIn(p.Birthday, t, secondsInMonth)
}

func (p *Person) genderRepr() string {
	choices := map[int]string{
		Male:   "Male",
		Female: "Female",
	}
	return choices[p.Gender]
}

/**
 * Private methods
 */

// elapsedFromDateIn calculates the amont of time passsed between two time.Time
// instances and convert it to a common time frame.
func elapsedFromDateIn(from time.Time, to time.Time, in float64) float64 {
	return to.Sub(from).Seconds() / in
}

// these variables make it easier the conversion between different times frames
var (
	daysInYear      = 365.25                                       // approximate days in a year, assuming no leap year
	daysInMonth     = daysInYear / 12                              // approximate days in month, assuming constant distribution
	secondsInMinute = 60.0                                         // seconds in a minute
	minutesInHour   = 60.0                                         // minutes in a hour
	hoursInDay      = 24.0                                         // hours in a day
	secondsInDay    = hoursInDay * minutesInHour * secondsInMinute // seconds in a days
	secondsInMonth  = daysInMonth * secondsInDay                   // approximate seconds in a month
)

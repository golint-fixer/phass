package phass

import "fmt"

/**
 * Constants
 */

// Skinfold constants.
const (
	// SKFSubscapular: subscapular skinfold.
	SKFSubscapular = iota
	// SKFTriceps: triceps skinfold.
	SKFTriceps
	// SKFBiceps: biceps skinfold.
	SKFBiceps
	// SKFChest: chest skinfold.
	SKFChest
	// SKFMidaxillary: mid-axillary skinfold.
	SKFMidaxillary
	// SKFSuprailiac: suprailiac skinfold.
	SKFSuprailiac
	// SKFAbdominal: abdominal skinfold.
	SKFAbdominal
	// SKFThigh: thigh skinfold.
	SKFThigh
	// SKFCalf: calf skinfold.
	SKFCalf
)

/*
 * Skinfolds
 */

// Skinfolds represent a collection of skinfold measurements.
type Skinfolds struct {
	Measures map[int]float64
}

// NewSkinfolds return a new Skinfolds instance.
func NewSkinfolds(measures map[int]float64) *Skinfolds {
	return &Skinfolds{Measures: measures}
}

func (s *Skinfolds) String() string {
	return fmt.Sprintf("Sum %d skinfolds: %.2f mm", len(s.Measures), s.Sum())
}

// GetName returns name for this measurement.
func (s *Skinfolds) GetName() string {
	return "Skinfolds"
}

// Result returns relevant information about this measurement.
func (s *Skinfolds) Result() ([]string, error) {
	rs := []string{}
	for k, v := range s.Measures {
		rs = append(rs, fmt.Sprintf("Skinfold %s: %.2f mm", NamedSkinfold(k), v))
	}
	rs = append(rs, fmt.Sprintf("Sum skinfolds: %.2f mm", s.Sum()))
	return rs, nil
}

// Sum all skinfolds values.
func (s *Skinfolds) Sum() float64 {
	accum := 0.0
	for _, v := range s.Measures {
		accum += v
	}
	return accum
}

// SumSpecific set of skinfolds measurements.
func (s *Skinfolds) SumSpecific(skinfolds []int) float64 {
	accum := 0.0
	for _, skf := range skinfolds {
		accum += s.Measures[skf]
	}
	return accum
}

// NamedSkinfold returns the name for a given skinfold constant.
func NamedSkinfold(name int) string {
	named := map[int]string{
		SKFSubscapular: "subscapular",
		SKFTriceps:     "triceps",
		SKFBiceps:      "biceps",
		SKFChest:       "chest",
		SKFMidaxillary: "mid-axillary",
		SKFSuprailiac:  "suprailiac",
		SKFAbdominal:   "abdominal",
		SKFThigh:       "thigh",
		SKFCalf:        "calf",
	}

	return named[name]
}

package sensor

import (
	"fmt"
	"time"
)

type State uint

const (
	Inactive State = iota + 1
	Active
	Outdated
)

type ISensor interface {
	Reading() (float64, State, error)
	ReadingAge() (time.Duration, error)
	Name() string
}

type Options struct {
	Name          string
	MaxReadingAge time.Duration
}

type Sensor struct {
	value       float64
	State       State
	UpdatedTime time.Time
	options     Options
}

func (s State) String() string {
	switch s {
	case Inactive:
		return "inactive"
	case Active:
		return "active"
	case Outdated:
		return "outdated"
	}
	return "unknown"
}

func (s *Sensor) IsReadingValid() bool {

	age, err := s.ReadingAge()

	if err != nil {
		return false
	}

	outdated := age - s.options.MaxReadingAge

	if outdated.Seconds() > 0 {
		return false
	}

	return true
}

func (s *Sensor) Reading() (float64, State, error) {

	age, err := s.ReadingAge()

	if err != nil {
		return 0, s.State, err
	}

	outdated := age - s.options.MaxReadingAge

	if outdated.Seconds() > 0 {
		s.State = Outdated
		// returns value anyway but with error message
		return s.value, s.State, fmt.Errorf("time value outdated %d seconds for %s sensor", outdated.Seconds(), s.options.Name)
	}

	return s.value, s.State, nil
}

func (s *Sensor) ReadingAge() (time.Duration, error) {
	if s.UpdatedTime.IsZero() {
		return 0, fmt.Errorf("no sensor reading available (yet?) for sensor '%s'", s.options.Name)
	}

	return time.Now().Sub(s.UpdatedTime), nil
}

func (s *Sensor) Name() string {
	return s.options.Name
}

func (s *Sensor) SetValue(value float64) *Sensor {
	s.UpdatedTime = time.Now()
	s.State = Active
	s.value = value
	return s
}

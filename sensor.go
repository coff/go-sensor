package sensor

import (
	"fmt"
	"time"
)

type iSensor interface {
	Reading() (float64, error)
}

type Options struct {
	name            string
	readingValidity time.Duration
}

type Sensor struct {
	value       float64
	UpdatedTime time.Time
	options     Options
}

func (s *Sensor) IsReadingValid() bool {
	if s.UpdatedTime.IsZero() {
		return false
	}

	outdated := time.Now().Sub(s.UpdatedTime) - s.options.readingValidity

	if outdated.Seconds() > 0 {
		return false
	}

	return true
}

func (s *Sensor) Reading() (float64, error) {

	if s.UpdatedTime.IsZero() {
		return 0, fmt.Errorf("no sensor reading available (yet?) for %s sensor", s.options.name)
	}

	outdated := time.Now().Sub(s.UpdatedTime) - s.options.readingValidity

	if outdated.Seconds() > 0 {
		// returns value anyway but with error message
		return s.value, fmt.Errorf("time value outdated %d seconds for %s sensor", outdated.Seconds(), s.options.name)
	}
	return s.value, nil
}

func (s *Sensor) SetValue(value float64) *Sensor {
	s.UpdatedTime = time.Now()
	s.value = value
	return s
}

package config

import "time"

// H ...
type H map[string]interface{}

// TimeIn ...
func (c *Config) TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

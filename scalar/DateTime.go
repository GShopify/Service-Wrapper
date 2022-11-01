package scalar

import (
	"encoding/json"
	"time"
)

type DateTime struct {
	time.Time
}

func NewDateTimeIn(seconds int) DateTime {
	return DateTime{
		time.Now().Add(time.Duration(seconds) * time.Second),
	}
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time)
}

func (d DateTime) String() string {
	return d.Time.String()
}

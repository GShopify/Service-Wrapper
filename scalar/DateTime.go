package scalar

import (
	"encoding/json"
	"io"
	"strconv"
	"time"
)

type DateTime int64

func NewDateTimeIn(seconds int) DateTime {
	return DateTime(
		time.Now().Add(time.Duration(seconds) * time.Second).UnixMilli())
}

func (d DateTime) String() string {
	return time.UnixMilli(int64(d)).String()
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.UnixMilli(int64(d)))
}

func (d DateTime) MarshalGQL(w io.Writer) {
	w.Write([]byte(strconv.Quote(d.String())))
}

func (d *DateTime) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case float64:
		*d = DateTime(v)
	case int:
		*d = DateTime(v)
	case string:
		t, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", v)
		if err != nil {
			return err
		}

		*d = DateTime(t.UnixMilli())
	}

	return nil
}

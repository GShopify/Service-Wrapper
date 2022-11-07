package scalar

import (
	"encoding/json"
	"io"
	"strconv"
	"time"
)

type DateTime int64

var datetimeLayout = "2006-01-02T15:04:05Z0700"

func NewDateTimeIn(seconds int) DateTime {
	return DateTime(
		time.Now().In(UTC).Add(time.Duration(seconds) * time.Second).UnixMilli())
}

func Now() DateTime {
	return DateTime(time.Now().In(UTC).UnixMilli())
}

func (d DateTime) String() string {
	return time.UnixMilli(int64(d)).In(UTC).Format(datetimeLayout)
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.UnixMilli(int64(d)).In(UTC))
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
	case time.Time:
		*d = DateTime(v.In(UTC).UnixMilli())
	case string:
		t, err := time.Parse(datetimeLayout, v)
		if err != nil {
			return err
		}

		*d = DateTime(t.In(UTC).UnixMilli())
	case DateTime:
		*d = v
	}

	return nil
}

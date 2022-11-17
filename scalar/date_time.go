package scalar

import (
	"encoding/json"
	"github.com/opentracing/opentracing-go/log"
	"io"
	"regexp"
	"strconv"
	"time"
)

type DateTime int64

var datetimeLayout = "2006-01-02T15:04:05Z0700"

func NewDateTimeIn(seconds int) DateTime {
	return DateTime(
		time.Now().In(time.UTC).Add(time.Duration(seconds) * time.Second).UnixMilli())
}

func NewDateTimeFrom(t time.Time) DateTime {
	return DateTime(t.In(time.UTC).UnixMilli())
}

func Now() DateTime {
	return DateTime(time.Now().In(time.UTC).UnixMilli())
}

func (d DateTime) String() string {
	return time.UnixMilli(int64(d)).In(time.UTC).Format(datetimeLayout)
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.UnixMilli(int64(d)).In(time.UTC))
}

func (d DateTime) MarshalGQL(w io.Writer) {
	_, err := w.Write([]byte(strconv.Quote(d.String())))
	if err != nil {
		log.Error(err)
	}
}

func (d *DateTime) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case float64:
		*d = DateTime(v)
	case int:
		*d = DateTime(v)
	case time.Time:
		*d = DateTime(v.In(time.UTC).UnixMilli())
	case string:
		if regexp.MustCompile(`^\d{13}$`).MatchString(v) {
			if i, err := strconv.Atoi(v); err == nil {
				*d = DateTime(i)
				return nil
			}
		}

		t, err := time.Parse(datetimeLayout, v)
		if err != nil {
			return err
		}

		*d = DateTime(t.In(time.UTC).UnixMilli())
	case DateTime:
		*d = v
	}

	return nil
}

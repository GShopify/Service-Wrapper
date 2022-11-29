package scalar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
)

type Decimal string

func NewDecimal(f any) Decimal {
	d := Decimal("")
	if err := d.UnmarshalGQL(f); err != nil {
		d = Decimal("0")
	}

	return d
}

func (d Decimal) Valid() bool {
	_, err := strconv.ParseFloat(d.String(), 64)
	return err == nil
}

func (d Decimal) String() string {
	return string(d)
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal([]byte(d))
}

func (d Decimal) MarshalGQL(w io.Writer) {
	_, err := w.Write([]byte(strconv.Quote(d.String())))
	if err != nil {
		log.Default().Println(err)
	}
}

func (d *Decimal) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case float64:
		*d = Decimal(fmt.Sprintf("%f", v))
		return nil
	case int:
		*d = Decimal(fmt.Sprintf("%d", v))
		return nil
	case string:
		*d = Decimal(v)
		if !d.Valid() {
			return fmt.Errorf("illegal Decimal value: `%s`", v)
		}

		return nil
	}

	return fmt.Errorf("illegal Decimal value: `%v`", d)
}

func (d *Decimal) Float() float64 {
	f, _ := strconv.ParseFloat(d.String(), 64)
	return f
}

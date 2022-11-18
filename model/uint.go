package model

import (
	"io"
	"log"
	"strconv"
)

type UInt uint64

func (i *UInt) UnmarshalGQL(v interface{}) error {
	var (
		value = uint64(0)
		err   error
	)

	switch v := v.(type) {
	case string:
		if value, err = strconv.ParseUint(v, 10, 64); err != nil {
			return err
		}
	case int:
		if v < 0 {
			v *= -1
		}

		value = uint64(v)

	case float64:
		if v < 0 {
			v *= -1
		}

		value = uint64(v)
	}

	*i = UInt(value)
	return nil
}

func (i UInt) MarshalGQL(w io.Writer) {
	_, err := w.Write(strconv.AppendUint(nil, uint64(i), 10))
	if err != nil {
		log.Default().Println(err)
	}
}

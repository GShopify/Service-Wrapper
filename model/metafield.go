package model

import (
	"fmt"
	"github.com/gshopify/service-wrapper/scalar"
)

func MetaFieldValue(t MetaFieldType, v any) string {
	switch t {
	case MetaFieldTypeSingleLineTextField:
		return fmt.Sprintf("%s", v)
	case MetaFieldTypeBoolean:
		return fmt.Sprintf("%v", v)
	case MetaFieldTypeDateTime:
		dt := scalar.DateTime(0)
		if err := dt.UnmarshalGQL(v); err != nil {
			return fmt.Sprintf("%v", v)
		}

		return dt.String()
	case MetaFieldTypeNumberInteger:
		return fmt.Sprintf("%d", v)
	}

	return fmt.Sprintf("%v", v)
}

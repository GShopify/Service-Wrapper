package model

import (
	"fmt"
	"github.com/gshopify/service-wrapper/scalar"
)

type MetafieldType string

const (
	MetafieldTypeBoolean             MetafieldType = "boolean"
	MetafieldTypeCollectionReference MetafieldType = "collection_reference"
	MetafieldTypeColor               MetafieldType = "color"
	MetafieldTypeDate                MetafieldType = "date"
	MetafieldTypeDateTime            MetafieldType = "date_time"
	MetafieldTypeDimension           MetafieldType = "dimension"
	MetafieldTypeFileReference       MetafieldType = "file_reference"
	MetafieldTypeJson                MetafieldType = "json"
	MetafieldTypeMoney               MetafieldType = "money"
	MetafieldTypeMultiLineTextField  MetafieldType = "multi_line_text_field"
	MetafieldTypeNumberDecimal       MetafieldType = "number_decimal"
	MetafieldTypeNumberInteger       MetafieldType = "number_integer"
	MetafieldTypePageReference       MetafieldType = "page_reference"
	MetafieldTypeProductReference    MetafieldType = "product_reference"
	MetafieldTypeRating              MetafieldType = "rating"
	MetafieldTypeSingleLineTextField MetafieldType = "single_line_text_field"
	MetafieldTypeUrl                 MetafieldType = "url"
	MetafieldTypeVariantReference    MetafieldType = "variant_reference"
	MetafieldTypeVolume              MetafieldType = "volume"
	MetafieldTypeWeight              MetafieldType = "weight"
)

func (m MetafieldType) String() string {
	return string(m)
}

func MetafieldValue(t MetafieldType, v any) string {
	switch t.String() {
	case "single_line_text_field":
		return fmt.Sprintf("%s", v)
	case "boolean":
		return fmt.Sprintf("%v", v)
	case "date_time":
		dt := scalar.DateTime(0)
		if err := dt.UnmarshalGQL(v); err != nil {
			return fmt.Sprintf("%v", v)
		}

		return fmt.Sprintf("%d", dt)
	}

	return fmt.Sprintf("%v", v)
}

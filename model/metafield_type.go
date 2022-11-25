package model

type MetaFieldType string

const (
	MetaFieldTypeBoolean             MetaFieldType = "boolean"
	MetaFieldTypeCollectionReference MetaFieldType = "collection_reference"
	MetaFieldTypeColor               MetaFieldType = "color"
	MetaFieldTypeDate                MetaFieldType = "date"
	MetaFieldTypeDateTime            MetaFieldType = "date_time"
	MetaFieldTypeDimension           MetaFieldType = "dimension"
	MetaFieldTypeFileReference       MetaFieldType = "file_reference"
	MetaFieldTypeJson                MetaFieldType = "json"
	MetaFieldTypeMoney               MetaFieldType = "money"
	MetaFieldTypeMultiLineTextField  MetaFieldType = "multi_line_text_field"
	MetaFieldTypeNumberDecimal       MetaFieldType = "number_decimal"
	MetaFieldTypeNumberInteger       MetaFieldType = "number_integer"
	MetaFieldTypePageReference       MetaFieldType = "page_reference"
	MetaFieldTypeProductReference    MetaFieldType = "product_reference"
	MetaFieldTypeRating              MetaFieldType = "rating"
	MetaFieldTypeSingleLineTextField MetaFieldType = "single_line_text_field"
	MetaFieldTypeUrl                 MetaFieldType = "url"
	MetaFieldTypeVariantReference    MetaFieldType = "variant_reference"
	MetaFieldTypeVolume              MetaFieldType = "volume"
	MetaFieldTypeWeight              MetaFieldType = "weight"
)

var AllMetaFieldTypes = []MetaFieldType{
	MetaFieldTypeBoolean, MetaFieldTypeCollectionReference, MetaFieldTypeColor, MetaFieldTypeDate, MetaFieldTypeDateTime,
	MetaFieldTypeDimension, MetaFieldTypeFileReference, MetaFieldTypeJson, MetaFieldTypeMoney, MetaFieldTypeMultiLineTextField,
	MetaFieldTypeNumberDecimal, MetaFieldTypeNumberInteger, MetaFieldTypePageReference, MetaFieldTypeProductReference,
	MetaFieldTypeRating, MetaFieldTypeSingleLineTextField, MetaFieldTypeUrl, MetaFieldTypeVariantReference,
	MetaFieldTypeVolume, MetaFieldTypeWeight,
}

func (e MetaFieldType) IsValid() bool {
	for _, t := range AllMetaFieldTypes {
		if t == e {
			return true
		}
	}

	return false
}

func (e MetaFieldType) String() string {
	return string(e)
}

package model

type ProductStatus string

const (
	ProductStatusDraft    ProductStatus = "draft"
	ProductStatusActive   ProductStatus = "active"
	ProductStatusArchived ProductStatus = "archived"
)

var AllProductStatuses = []ProductStatus{
	ProductStatusDraft,
	ProductStatusActive,
	ProductStatusArchived,
}

func (e ProductStatus) IsValid() bool {
	for _, t := range AllProductStatuses {
		if t == e {
			return true
		}
	}

	return false
}

func (e ProductStatus) String() string {
	return string(e)
}

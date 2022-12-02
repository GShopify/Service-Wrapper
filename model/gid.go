package model

type Gid string

const (
	GidProduct       Gid = "Product"
	GidCollection    Gid = "Collection"
	GidOption        Gid = "ProductOption"
	GidVariant       Gid = "ProductVariant"
	GidInventoryItem Gid = "InventoryItem"
	GidLocation      Gid = "Location"
)

var AllGids = []Gid{
	GidProduct, GidCollection, GidOption, GidVariant, GidInventoryItem, GidLocation,
}

func (e Gid) IsValid() bool {
	for _, t := range AllGids {
		if t == e {
			return true
		}
	}

	return false
}

func (e Gid) String() string {
	return string(e)
}

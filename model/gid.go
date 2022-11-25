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

func (e Gid) String() string {
	return string(e)
}

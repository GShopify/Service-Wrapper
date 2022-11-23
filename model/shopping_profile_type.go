package model

type ShoppingProfileType string

const (
	ShoppingProfileTypeDefault    ShoppingProfileType = "default"
	ShoppingProfileTypeGiftCard   ShoppingProfileType = "gift_card"
	ShoppingProfileTypeGiftCustom ShoppingProfileType = "custom"
)

var AllShoppingProfileTypes = []ShoppingProfileType{
	ShoppingProfileTypeDefault,
	ShoppingProfileTypeGiftCard,
	ShoppingProfileTypeGiftCustom,
}

func (e ShoppingProfileType) IsValid() bool {
	for _, t := range AllShoppingProfileTypes {
		if t == e {
			return true
		}
	}

	return false
}

func (e ShoppingProfileType) String() string {
	return string(e)
}

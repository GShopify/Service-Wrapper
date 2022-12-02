package model

type ProductType struct {
	Id         string   `db:"id" json:"id"`
	Type       string   `db:"type" json:"type"`
	ProductIds []string `db:"products" json:"product_ids"`
	Count      int      `db:"count" json:"count"`
}

func (t *ProductType) IsNode() {}

func (t *ProductType) GetID() string {
	return t.Type
}

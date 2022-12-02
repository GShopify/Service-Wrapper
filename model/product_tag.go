package model

type ProductTag struct {
	Id         string   `db:"id" json:"id"`
	Tag        string   `db:"tag" json:"tag"`
	ProductIds []string `db:"products" json:"product_ids"`
	Count      int      `db:"count" json:"count"`
}

func (t ProductTag) IsNode() {}

func (t ProductTag) GetID() string {
	return t.Tag
}

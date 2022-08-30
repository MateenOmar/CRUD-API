package models

type Data struct {
	Category string `json:"category" bson:"category"`
	Type     string `json:"type" bson:"type"`
}

type Product struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Quantity int    `json:"quantity" bson:"quantity"`
}

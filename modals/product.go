package modals

type Product struct {
	ID          string `bson:"_id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Price       int    `bson:"price" json:"price"`
}

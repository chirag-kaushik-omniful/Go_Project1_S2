package modals

type Order struct {
	ID         string `bson:"_id"`
	CustomerId string `bson:"customerid"`
}

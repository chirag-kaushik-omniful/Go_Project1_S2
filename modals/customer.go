package modals

type Customer struct {
	ID       string `bson:"_id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Mobile   string `bson:"mobile" json:"mobile"`
	Password string `bson:"password" json:"password"`
}

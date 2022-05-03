package dao

// Contact is a subset of payment contact information as it is published on NATS
type Contact struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}

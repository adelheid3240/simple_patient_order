package model

type Patient struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

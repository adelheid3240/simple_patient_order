package model

type Patient struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type PatientOrder struct {
	ID          string `bson:"_id,omitempty"` // mongodb auto create id by ignoring this field
	PatientID   string `bson:"patient_id"`
	Message     string `bson:"message"`
	CreatedTime int64  `bson:"created_time"`
}

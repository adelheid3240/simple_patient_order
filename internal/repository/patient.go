package repository

import (
	"context"
	"simplepatientorder/internal/model"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Patient interface {
	List(ctx context.Context, size int64) ([]model.Patient, error)
}

type patient struct {
	collection *mongo.Collection
}

func NewPatient(mongoClient *mongo.Client) Patient {
	return &patient{
		collection: mongoClient.Database("patients").Collection("patients"),
	}
}

func (p *patient) List(ctx context.Context, size int64) ([]model.Patient, error) {
	findOption := new(options.FindOptions)
	findOption.SetLimit(size)

	cur, err := p.collection.Find(ctx, bson.D{}, findOption)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list patients")
	}
	defer cur.Close(ctx)

	var patients []model.Patient
	for cur.Next(ctx) {
		var m model.Patient
		err := cur.Decode(&m)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode patient when list patients")
		}

		patients = append(patients, m)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error when list patients")
	}

	return patients, nil
}

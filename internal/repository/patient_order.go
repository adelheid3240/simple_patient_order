package repository

import (
	"context"
	"simplepatientorder/internal/model"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PatientOrder interface {
	Create(ctx context.Context, order *model.PatientOrder) error
	List(ctx context.Context, patientID string) ([]model.PatientOrder, error)
}

type patienteOrder struct {
	collection *mongo.Collection
}

func NewPatientOrder(mongoClient *mongo.Client) PatientOrder {
	return &patienteOrder{
		collection: mongoClient.Database("patients").Collection("orders"),
	}
}

func (p *patienteOrder) Create(ctx context.Context, order *model.PatientOrder) error {
	_, err := p.collection.InsertOne(ctx, order)
	if err != nil {
		return errors.Wrap(err, "failed to create patient order")
	}

	return nil
}

func (p *patienteOrder) List(ctx context.Context, patientID string) ([]model.PatientOrder, error) {
	cur, err := p.collection.Find(ctx, bson.D{{"patient_id", patientID}})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list patient orders")
	}
	defer cur.Close(ctx)

	var orders []model.PatientOrder
	for cur.Next(ctx) {
		var m model.PatientOrder
		err := cur.Decode(&m)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode patient order when list patient orders")
		}

		orders = append(orders, m)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error when list patient orders")
	}

	return orders, nil
}

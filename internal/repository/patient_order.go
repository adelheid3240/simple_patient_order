package repository

import (
	"context"
	"simplepatientorder/internal/model"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type PatientOrder interface {
	Create(ctx context.Context, order *model.PatientOrder) error
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

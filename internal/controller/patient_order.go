package controller

import (
	"context"
	"simplepatientorder/internal/model"
	"simplepatientorder/internal/repository"
	"time"
)

type PatientOrder interface {
	Create(ctx context.Context, patientID, message string) error
	List(ctx context.Context, patientID string) ([]model.PatientOrder, error)
	Update(ctx context.Context, id string, message string) error
	Delete(ctx context.Context, id string) error
}

type patientOrder struct {
	patientOrderRepo repository.PatientOrder
}

func NewPatientOrder(patientOrderRepo repository.PatientOrder) PatientOrder {
	return &patientOrder{
		patientOrderRepo: patientOrderRepo,
	}
}

func (p *patientOrder) Create(ctx context.Context, patientID, message string) error {
	return p.patientOrderRepo.Create(ctx, &model.PatientOrder{
		PatientID:   patientID,
		Message:     message,
		CreatedTime: time.Now().UnixMilli(),
		UpdatedTime: time.Now().UnixMilli(),
	})
}

func (p *patientOrder) List(ctx context.Context, patientID string) ([]model.PatientOrder, error) {
	return p.patientOrderRepo.List(ctx, patientID)
}

func (p *patientOrder) Update(ctx context.Context, id string, message string) error {
	return p.patientOrderRepo.Update(ctx, &model.PatientOrder{
		ID:          id,
		Message:     message,
		UpdatedTime: time.Now().UnixMilli(),
	})
}

func (p *patientOrder) Delete(ctx context.Context, id string) error {
	return p.patientOrderRepo.Delete(ctx, id)
}

package controller

import (
	"context"
	"simplepatientorder/internal/model"
	"simplepatientorder/internal/repository"
)

type Patient interface {
	List(ctx context.Context, size int64) ([]model.Patient, error)
}

type patient struct {
	patientRepo repository.Patient
}

func NewPatient(patientRepo repository.Patient) Patient {
	return &patient{
		patientRepo: patientRepo,
	}
}

func (p *patient) List(ctx context.Context, size int64) ([]model.Patient, error) {
	return p.patientRepo.List(ctx, size)
}

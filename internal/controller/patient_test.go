package controller

import (
	"context"
	"simplepatientorder/internal/model"
	"simplepatientorder/internal/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_patient_List(t *testing.T) {
	mocker := gomock.NewController(t)
	patientRepo := repository.NewMockPatient(mocker)
	ctrl := NewPatient(patientRepo)
	ctx := context.Background()

	patientRepo.EXPECT().List(ctx, int64(1)).Return([]model.Patient{{ID: "1", Name: "name1"}}, nil)

	patients, err := ctrl.List(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, []model.Patient{
		{ID: "1", Name: "name1"},
	}, patients)
}

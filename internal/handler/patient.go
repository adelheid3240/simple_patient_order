package handler

import (
	"net/http"
	"simplepatientorder/internal/apierr"
	"simplepatientorder/internal/controller"
	"simplepatientorder/internal/dto"

	"github.com/gin-gonic/gin"
)

type Patient interface {
	List(c *gin.Context)
}

type patient struct {
	patientCtrl controller.Patient
}

func NewPatient(patientCtrl controller.Patient) Patient {
	return &patient{
		patientCtrl: patientCtrl,
	}
}

func (p *patient) List(c *gin.Context) {
	patients, err := p.patientCtrl.List(c, 5) // if need paging can add page params
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, apierr.ErrInternal.SetErr(err))
		return
	}

	respPatients := make([]dto.Patient, len(patients))
	for i, p := range patients {
		respPatients[i] = dto.Patient{
			ID:   p.ID,
			Name: p.Name,
		}
	}

	c.JSON(http.StatusOK, dto.ListPatientsResp{Patients: respPatients})
}

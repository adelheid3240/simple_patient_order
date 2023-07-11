package handler

import (
	"errors"
	"net/http"
	"simplepatientorder/internal/controller"
	"simplepatientorder/internal/dto"

	"github.com/gin-gonic/gin"
)

type PatientOrder interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type patientOrder struct {
	patientOrderCtrl controller.PatientOrder
}

func NewPatientOrder(patientOrderCtrl controller.PatientOrder) PatientOrder {
	return &patientOrder{
		patientOrderCtrl: patientOrderCtrl,
	}
}

func (p *patientOrder) Create(c *gin.Context) {
	patientID := c.Param("id")
	if patientID == "" {
		c.AbortWithError(http.StatusForbidden, errors.New("invalid params"))
		return
	}

	order := dto.CreateOrUpdatePatientOrder{}
	if err := c.BindJSON(&order); err != nil {
		c.AbortWithError(http.StatusForbidden, err)
		return
	}

	if err := p.patientOrderCtrl.Create(c, patientID, order.Message); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "")
}

func (p *patientOrder) List(c *gin.Context) {
	patientID := c.Param("id")
	if patientID == "" {
		c.AbortWithError(http.StatusForbidden, errors.New("invalid params"))
		return
	}

	orders, err := p.patientOrderCtrl.List(c, patientID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	respOrders := make([]dto.PatientOrder, len(orders))
	for i, po := range orders {
		respOrders[i] = dto.PatientOrder{
			ID:      po.ID,
			Message: po.Message,
		}
	}

	c.JSON(http.StatusOK, dto.ListPatientOrderResp{Orders: respOrders})
}

func (p *patientOrder) Update(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.AbortWithError(http.StatusForbidden, errors.New("invalid params"))
		return
	}

	order := dto.CreateOrUpdatePatientOrder{}
	if err := c.BindJSON(&order); err != nil {
		c.AbortWithError(http.StatusForbidden, err)
		return
	}

	if err := p.patientOrderCtrl.Update(c, orderID, order.Message); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "")
}

func (p *patientOrder) Delete(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.AbortWithError(http.StatusForbidden, errors.New("invalid params"))
		return
	}

	if err := p.patientOrderCtrl.Delete(c, orderID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "")
}

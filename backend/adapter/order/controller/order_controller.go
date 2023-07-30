package controller

import (
	"errors"
	"net/http"

	entityOrder "github.com/CHTTCH/little_project/backend/entity/order"
	entityPatient "github.com/CHTTCH/little_project/backend/entity/patient"
	usecaseOrder "github.com/CHTTCH/little_project/backend/usecase/order"
	"github.com/CHTTCH/little_project/backend/usecase/repository"
	"github.com/gin-gonic/gin"
)

func CreateOrderController(patientRepo repository.Repository[entityPatient.Patient, string], orderRepo repository.Repository[entityOrder.Order, int]) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := usecaseOrder.CreateOrderInput{}

		if patientId, isExist := c.GetPostForm("PatientId"); isExist && patientId != "" {
			input.SetPatientId(patientId)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("patient id sholud exist").Error()})
			return
		}

		if message, isExist := c.GetPostForm("Message"); isExist && message != "" {
			input.SetMessage(message)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("message sholud exist").Error()})
			return
		}

		output := usecaseOrder.CreateOrder(patientRepo, orderRepo, input)

		if output.GetResult() {
			c.JSON(http.StatusOK, gin.H{"Id": output.GetId()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": output.GetMessage()})
		}
	}
}

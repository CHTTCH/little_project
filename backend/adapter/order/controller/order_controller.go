package controller

import (
	"errors"
	"net/http"
	"strconv"

	entityOrder "github.com/CHTTCH/little_project/backend/entity/order"
	entityPatient "github.com/CHTTCH/little_project/backend/entity/patient"
	usecaseOrder "github.com/CHTTCH/little_project/backend/usecase/order"
	"github.com/CHTTCH/little_project/backend/usecase/repository"
	"github.com/gin-gonic/gin"
)

func CreateOrderController(patientRepo repository.Repository[entityPatient.Patient, string], orderRepo repository.Repository[entityOrder.Order, int]) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := new(usecaseOrder.CreateOrderInput)

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

func EditOrderController(orderRepo repository.Repository[entityOrder.Order, int]) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := new(usecaseOrder.EditOrderInput)

		if orderId, isExist := c.GetPostForm("OrderId"); isExist && orderId != "" {
			orderId, err := strconv.Atoi(orderId)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("order id sholud be int").Error()})
				return
			}

			input.SetOrderId(orderId)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("order id sholud exist").Error()})
			return
		}

		if modifiedMessage, isExist := c.GetPostForm("ModifiedMessage"); isExist && modifiedMessage != "" {
			input.SetModifiedMessage(modifiedMessage)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("modified message sholud exist").Error()})
			return
		}

		output := usecaseOrder.EditOrder(orderRepo, input)

		if output.GetResult() {
			c.JSON(http.StatusOK, gin.H{"Id": output.GetId()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": output.GetMessage()})
		}
	}
}

func FindAllOrdersController(repo repository.Repository[entityOrder.Order, int]) gin.HandlerFunc {
	return func(c *gin.Context) {
		output := usecaseOrder.FindAllOrders(repo)

		if output.GetResult() {
			c.JSON(http.StatusOK, output.GetData())
		} else if output.Message == "no orders" {
			c.JSON(http.StatusOK, []any{})
		}else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("find all orders failed").Error()})
		}
	}
}

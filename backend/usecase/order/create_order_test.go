package order

import (
	"fmt"
	"testing"

	"github.com/CHTTCH/little_project/backend/entity/patient"
	mockOrderRepo "github.com/CHTTCH/little_project/backend/test/mock_repository/order"
	mockPatientRepo "github.com/CHTTCH/little_project/backend/test/mock_repository/patient"
	"github.com/stretchr/testify/assert"
)

func TestCreateAnOrderThenSuccess(t *testing.T) {
	assert := assert.New(t)

	patientId := "1"
	patientName := "小明"
	patientRepo := mockPatientRepo.NewMockPatientRepository()
	patientRepo.Save(patient.NewPatient(patientId, patientName))

	message := "超過120請施打8u"
	orderRepo := mockOrderRepo.NewMockOrderRepository()
	orderInput := NewCreateOrderInput(patientId, message)

	createOrderOutput := CreateOrder(patientRepo, orderRepo, orderInput)

	expectedOrderId := 1
	assert.Equal(true, createOrderOutput.GetResult())
	assert.Equal(fmt.Sprintf("%d", expectedOrderId), createOrderOutput.GetId())
	assert.Equal("", createOrderOutput.GetMessage())

	order, _ := orderRepo.FindById(expectedOrderId)
	assert.Equal(message, order.GetMessage())

	patient, _ := patientRepo.FindById(patientId)
	assert.Equal(expectedOrderId, patient.GetOrderId())
}

func TestCreateAnOrderWithNotExistPatientWillFailed(t *testing.T) {
	assert := assert.New(t)

	notExistPatientId := "1"
	message := "超過120請施打8u"
	patientRepo := mockPatientRepo.NewMockPatientRepository()
	orderRepo := mockOrderRepo.NewMockOrderRepository()
	orderInput := NewCreateOrderInput(notExistPatientId, message)

	createOrderOutput := CreateOrder(patientRepo, orderRepo, orderInput)

	assert.Equal(false, createOrderOutput.GetResult())
	assert.Equal("patient not found", createOrderOutput.GetMessage())
	assert.Equal("", createOrderOutput.GetId())

	orders, _ := orderRepo.FindAll()
	assert.Equal(0, len(orders))
}

package order

import (
	"fmt"
	"testing"

	mockOrderRepo "github.com/CHTTCH/little_project/backend/adapter/order/mock_repository"
	mockPatientRepo "github.com/CHTTCH/little_project/backend/adapter/patient/mock_repository"
	usecasePatient "github.com/CHTTCH/little_project/backend/usecase/patient"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	patientId := "1"
	patientName := "小明"
	patientRepo := mockPatientRepo.NewMockPatientRepository()
	patientInput := usecasePatient.NewCreatePatientInput(patientId, patientName)
	_ = usecasePatient.CreatePatient(patientRepo, patientInput)

	orderId := 1
	message := "超過120請施打8u"
	orderRepo := mockOrderRepo.NewMockOrderRepository()
	orderInput := NewCreateOrderInput(patientId, message)

	createOrderOutput := CreateOrder(patientRepo, orderRepo, orderInput)

	patients, _ := patientRepo.FindAll()
	orders, _ := orderRepo.FindAll()

	assert.Equal(t, orderId, orders[0].GetId())
	assert.Equal(t, message, orders[0].GetMessage())
	assert.Equal(t, orderId, patients[0].GetOrderId())
	assert.Equal(t, fmt.Sprintf("%d", orderId), createOrderOutput.GetId())
	assert.Equal(t, true, createOrderOutput.GetResult())
}

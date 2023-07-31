package order

import (
	"fmt"
	"testing"

	mockOrderRepo "github.com/CHTTCH/little_project/backend/adapter/order/mock_repository"
	mockPatientRepo "github.com/CHTTCH/little_project/backend/adapter/patient/mock_repository"
	usecasePatient "github.com/CHTTCH/little_project/backend/usecase/patient"
	"github.com/stretchr/testify/assert"
)

func TestEditOrder(t *testing.T) {
	patientId := "1"
	patientName := "小明"
	patientRepo := mockPatientRepo.NewMockPatientRepository()
	createPatientInput := usecasePatient.NewCreatePatientInput(patientId, patientName)
	_ = usecasePatient.CreatePatient(patientRepo, createPatientInput)

	orderId := 1
	message := "超過120請施打8u"
	orderRepo := mockOrderRepo.NewMockOrderRepository()
	createOrderInput := NewCreateOrderInput(patientId, message)
	_ = CreateOrder(patientRepo, orderRepo, createOrderInput)

	modifiedMessage := "超過144請施打10u"
	editOrderInput := NewEditOrderInput(orderId, modifiedMessage)
	editOrderOutput := EditOrder(orderRepo, editOrderInput)

	order, err := orderRepo.FindById(orderId)

	assert.Equal(t, nil, err)
	assert.Equal(t, modifiedMessage, order.GetMessage())
	assert.Equal(t, fmt.Sprintf("%d", orderId), editOrderOutput.GetId())
	assert.Equal(t, true, editOrderOutput.GetResult())
}

package order

import (
	"testing"

	mockOrderRepo "github.com/CHTTCH/little_project/backend/test/mock_repository/order"
	mockPatientRepo "github.com/CHTTCH/little_project/backend/test/mock_repository/patient"
	usecasePatient "github.com/CHTTCH/little_project/backend/usecase/patient"
	"github.com/stretchr/testify/assert"
)

func TestFindAllOrder(t *testing.T) {
	patientId := "1"
	patientName := "小明"
	patientRepo := mockPatientRepo.NewMockPatientRepository()
	patientInput := usecasePatient.NewCreatePatientInput(patientId, patientName)
	_ = usecasePatient.CreatePatient(patientRepo, patientInput)

	orderId := 1
	message := "超過120請施打8u"
	orderRepo := mockOrderRepo.NewMockOrderRepository()
	orderInput := NewCreateOrderInput(patientId, message)
	_ = CreateOrder(patientRepo, orderRepo, orderInput)

	output := FindAllOrders(orderRepo)

	assert.Equal(t, orderId, output.GetData()[0].GetId())
	assert.Equal(t, message, output.GetData()[0].GetMessage())
}

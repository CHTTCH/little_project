package order

import (
	"fmt"
	"github.com/CHTTCH/little_project/backend/entity/order"
	"github.com/CHTTCH/little_project/backend/entity/patient"
	commandOutput "github.com/CHTTCH/little_project/backend/usecase/command_output"
	"github.com/CHTTCH/little_project/backend/usecase/repository"
)

func CreateOrder(patientRepo repository.Repository[patient.Patient, string], orderRepo repository.Repository[order.Order, int], input *CreateOrderInput) commandOutput.CommandOutput {

	patient, err := patientRepo.FindById(input.GetPatientId())
	if err != nil {
		return commandOutput.CommandOutput{Result: false, Message: "patient not found"}
	} else if patient.GetOrderId() != 0 {
		return commandOutput.CommandOutput{Result: false, Message: "patient already has order"}
	}

	o := order.NewOrder(input.GetMessage())

	if err := orderRepo.Save(o); err != nil {
		return commandOutput.CommandOutput{Result: false, Message: err.Error()}
	}

	orderId := o.GetId()
	patient.SetOrderId(orderId)

	if err := patientRepo.Save(patient); err != nil {
		return commandOutput.CommandOutput{Result: false, Message: err.Error()}
	}

	return commandOutput.CommandOutput{Id: fmt.Sprintf("%d", orderId), Result: true}
}

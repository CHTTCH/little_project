package order

import (
	"fmt"
	"github.com/CHTTCH/little_project/backend/entity/order"
	commandOutput "github.com/CHTTCH/little_project/backend/usecase/command_output"
	"github.com/CHTTCH/little_project/backend/usecase/repository"
)

func EditOrder(orderRepo repository.Repository[order.Order, int], input *EditOrderInput) commandOutput.CommandOutput {
	order, err := orderRepo.FindById(input.GetOrderId())

	if err != nil {
		return commandOutput.CommandOutput{Result: false, Message: "order not found"}
	} else {
		order.SetMessage(input.modifiedMessage)
	}

	if err := orderRepo.Save(order); err != nil {
		return commandOutput.CommandOutput{Result: false, Message: err.Error()}
	}

	return commandOutput.CommandOutput{Id: fmt.Sprintf("%d", order.GetId()), Result: true}
}

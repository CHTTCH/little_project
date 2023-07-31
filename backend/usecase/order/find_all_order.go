package order

import (
	"strings"
	"github.com/CHTTCH/little_project/backend/entity/order"
	queryOutput "github.com/CHTTCH/little_project/backend/usecase/query_output"
	"github.com/CHTTCH/little_project/backend/usecase/repository"
)

func FindAllOrders(repo repository.Repository[order.Order, int]) queryOutput.QueryOutput[order.Order] {
	data, err := repo.FindAll()
	
	if err != nil && strings.Contains(err.Error(), "ERROR: relation \"orders\" does not exist") {
		return queryOutput.QueryOutput[order.Order]{Result: false, Message: "no orders"}
	} else if err != nil {
		return queryOutput.QueryOutput[order.Order]{Result: false, Message: "find all failed"}
	}

	return queryOutput.QueryOutput[order.Order]{Data: data, Result: true}
}

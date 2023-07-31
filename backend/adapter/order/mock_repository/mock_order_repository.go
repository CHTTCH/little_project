package mock_repository

import (
	customizeError "github.com/CHTTCH/little_project/backend/adapter/error"
	"github.com/CHTTCH/little_project/backend/entity/order"
)

type MockOrderRepository struct {
	orderList []order.Order
}

func NewMockOrderRepository() *MockOrderRepository {
	return &MockOrderRepository{orderList: []order.Order{}}
}

func (r *MockOrderRepository) Save(o *order.Order) error {
	o.SetId(1)
	r.orderList = append(r.orderList, *o)
	return nil
}

func (r *MockOrderRepository) FindAll() ([]order.Order, error) {
	return r.orderList, nil
}

func (r *MockOrderRepository) FindById(id int) (*order.Order, error) {
	for _, order := range r.orderList {
		if order.GetId() == id {
			return &order, nil
		}
	}

	return nil, &customizeError.CustoMizeError{Message: "patient not found"}
}

package order

import (
	"testing"

	"github.com/CHTTCH/little_project/backend/entity/order"
	mockOrderRepo "github.com/CHTTCH/little_project/backend/test/mock_repository/order"
	"github.com/stretchr/testify/assert"
)

func TestFindAllOrdersThenSuccess(t *testing.T) {
	assert := assert.New(t)

	repo := mockOrderRepo.NewMockOrderRepository()

	o1 := order.NewOrder("超過120請施打8u")
	repo.Save(o1)

	output := FindAllOrders(repo)

	assert.Equal(true, output.GetResult())
	assert.Equal(1, len(output.GetData()))
	assert.Equal(*o1, output.GetData()[0])
	assert.Equal("", output.GetMessage())
}

func TestFindAllOrdersWhenNoOrderExistThenSuccess(t *testing.T) {
	assert := assert.New(t)

	repo := mockOrderRepo.NewMockOrderRepository()

	output := FindAllOrders(repo)

	assert.Equal(true, output.GetResult())
	assert.Equal(0, len(output.GetData()))
	assert.Equal("", output.GetMessage())
}

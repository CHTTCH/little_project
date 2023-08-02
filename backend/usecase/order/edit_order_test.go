package order

import (
	"fmt"
	"testing"

	"github.com/CHTTCH/little_project/backend/entity/order"
	mockOrderRepo "github.com/CHTTCH/little_project/backend/test/mock_repository/order"
	"github.com/stretchr/testify/assert"
)

func TestEditAnOrderThenSuccess(t *testing.T) {
	assert := assert.New(t)

	message := "超過120請施打8u"
	orderRepo := mockOrderRepo.NewMockOrderRepository()
	order := order.NewOrder(message)
	orderRepo.Save(order)

	orderId := order.GetId()
	newMessage := "超過150請施打8u"
	input := NewEditOrderInput(orderId, newMessage)

	output := EditOrder(orderRepo, input)

	assert.Equal(true, output.GetResult())
	assert.Equal(fmt.Sprintf("%d", orderId), output.GetId())
	assert.Equal("", output.GetMessage())

	editedOrder, _ := orderRepo.FindById(orderId)
	assert.Equal(newMessage, editedOrder.GetMessage())
}

func TestOrderNotExistWillFailed(t *testing.T) {
	assert := assert.New(t)

	message := "超過120請施打8u"
	notExistOrderId := 1
	orderRepo := mockOrderRepo.NewMockOrderRepository()
	input := NewEditOrderInput(notExistOrderId, message)

	output := EditOrder(orderRepo, input)

	assert.Equal(false, output.GetResult())
	assert.Equal("order not found", output.GetMessage())
	assert.Equal("", output.GetId())
}

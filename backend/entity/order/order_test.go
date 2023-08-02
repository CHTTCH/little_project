package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllGetter(t *testing.T) {
	assert := assert.New(t)

	id := 1
	message := "超過120請施打8u"

	o := Order{Id: id, Message: message}

	assert.Equal(id, o.GetId())
	assert.Equal(message, o.GetMessage())
}

func TestNewOrder(t *testing.T) {
	assert := assert.New(t)

	order := NewOrder("超過120請施打8u")

	assert.Equal(0, order.GetId())
	assert.Equal("超過120請施打8u", order.GetMessage())
}

func TestSetId(t *testing.T) {
	assert := assert.New(t)

	order := Order{Id: 1, Message: "message"}

	order.SetId(2)

	assert.Equal(2, order.GetId())
	assert.Equal("message", order.GetMessage())
}

func TestSetMessage(t *testing.T) {
	assert := assert.New(t)

	order := Order{Id: 3, Message: "oldMessage"}

	order.SetMessage("newMessage")

	assert.Equal("newMessage", order.GetMessage())
	assert.Equal(3, order.GetId())
}

package patient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatientStruct(t *testing.T) {
	id := "1"
	name := "小明"
	orderId := 1

	p := Patient{Id: id, Name: name, OrderId: orderId}

	assert.Equal(t, id, p.GetId())
	assert.Equal(t, name, p.GetName())
	assert.Equal(t, orderId, p.GetOrderId())
}

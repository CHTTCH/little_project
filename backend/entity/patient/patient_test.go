package patient

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatient(t *testing.T) {
	id := "1"
	name := "小明"
	orderId := "1"

	p := Patient{
		Id: id,
		Name: name,
		OrderId: orderId,
	}

	assert.Equal(t, id, p.getId())
	assert.Equal(t, name, p.getName())
	assert.Equal(t, orderId, p.getOrderId())
}
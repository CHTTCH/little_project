package order

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderStruct(t *testing.T) {
	id := 1
	message := "超過120請施打8u"

	o := Order{Id: id, Message: message}

	assert.Equal(t, id, o.GetId())
	assert.Equal(t, message, o.GetMessage())
}

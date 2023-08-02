package patient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllGetter(t *testing.T) {
	assert := assert.New(t)

	patient := Patient{Id: "id", Name: "name", OrderId: 1}

	assert.Equal("id", patient.GetId())
	assert.Equal("name", patient.GetName())
	assert.Equal(1, patient.GetOrderId())
}

func TestNewPatient(t *testing.T) {
	assert := assert.New(t)

	patient := NewPatient("id", "name")

	assert.Equal("id", patient.GetId())
	assert.Equal("name", patient.GetName())
	assert.Equal(0, patient.GetOrderId())
}

func TestSetId(t *testing.T) {
	assert := assert.New(t)

	patient := Patient{Id: "oldId", Name: "name", OrderId: 1}

	patient.SetId("newId")

	assert.Equal("newId", patient.GetId())
	assert.Equal("name", patient.GetName())
	assert.Equal(1, patient.GetOrderId())
}

func TestSetName(t *testing.T) {
	assert := assert.New(t)

	patient := Patient{Id: "id", Name: "oldName", OrderId: 1}

	patient.SetName("newName")

	assert.Equal("newName", patient.GetName())
	assert.Equal("id", patient.GetId())
	assert.Equal(1, patient.GetOrderId())
}

func TestSetOrderId(t *testing.T) {
	assert := assert.New(t)

	patient := Patient{Id: "id", Name: "name", OrderId: 1}

	patient.SetOrderId(2)

	assert.Equal(2, patient.GetOrderId())
	assert.Equal("id", patient.GetId())
	assert.Equal("name", patient.GetName())
}

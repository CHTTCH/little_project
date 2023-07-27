package patient

import (
	entityPatient "github.com/CHTTCH/little_project/backend/entity/patient"
	adapterPatient "github.com/CHTTCH/little_project/backend/adapter/patient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePatient(t *testing.T) {
	id := "1"
	name := "小明"
	orderId := "1"

	repo := &adapterPatient.MockPatientRepository{PatientList: []entityPatient.Patient{}}
	input := &CreatePatientInput{id: id, name: name, orderId: orderId}
	output := &CreatePatientOutput{}

	CreatePatient(repo, input, output)

	assert.Equal(t, id, repo.GetPatientList()[0].GetId())
	assert.Equal(t, name, repo.GetPatientList()[0].GetName())
	assert.Equal(t, orderId, repo.GetPatientList()[0].GetOrderId())
	assert.Equal(t, id, output.GetId())
}

package patient

import (
	"testing"
	mockPatientRepo "github.com/CHTTCH/little_project/backend/adapter/patient/mock_repository"
	entityPatient "github.com/CHTTCH/little_project/backend/entity/patient"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatientSucceed(t *testing.T) {
	id := "1"
	name := "小明"

	repo := &mockPatientRepo.MockPatientRepository{PatientList: []entityPatient.Patient{}}
	input := CreatePatientInput{id: id, name: name}

	output := CreatePatient(repo, input)

	patients, _ := repo.FindAll()

	assert.Equal(t, id, patients[0].GetId())
	assert.Equal(t, name, patients[0].GetName())
	assert.Equal(t, id, output.GetId())
	assert.Equal(t, true, output.GetResult())
}

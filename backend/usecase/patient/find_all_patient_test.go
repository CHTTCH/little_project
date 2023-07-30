package patient

import (
	"testing"
	mockPatientRepo "github.com/CHTTCH/little_project/backend/adapter/patient/mock_repository"
	entityPatient "github.com/CHTTCH/little_project/backend/entity/patient"
	"github.com/stretchr/testify/assert"
)

func TestFindAllPatientSucceed(t *testing.T) {
	id1 := "1"
	name1 := "小明"

	repo := &mockPatientRepo.MockPatientRepository{PatientList: []entityPatient.Patient{}}
	input1 := CreatePatientInput{id: id1, name: name1}
	_ = CreatePatient(repo, input1)

	id2 := "2"
	name2 := "小楊"
	input2 := CreatePatientInput{id: id2, name: name2}
	_ = CreatePatient(repo, input2)

	output := FindAllPatient(repo)

	assert.Equal(t, id1, output.GetData()[0].GetId())
	assert.Equal(t, name1, output.GetData()[0].GetName())
	assert.Equal(t, id2, output.GetData()[1].GetId())
	assert.Equal(t, name2, output.GetData()[1].GetName())

}

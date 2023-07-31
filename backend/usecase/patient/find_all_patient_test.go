package patient

import (
	mockPatientRepo "github.com/CHTTCH/little_project/backend/adapter/patient/mock_repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllPatientSucceed(t *testing.T) {
	id1 := "1"
	name1 := "小明"

	repo := mockPatientRepo.NewMockPatientRepository()
	input1 := NewCreatePatientInput(id1, name1)
	// input1 := CreatePatientInput{id: id1, name: name1}
	_ = CreatePatient(repo, input1)

	id2 := "2"
	name2 := "小楊"
	input2 := NewCreatePatientInput(id2, name2)
	_ = CreatePatient(repo, input2)

	output := FindAllPatient(repo)

	assert.Equal(t, id1, output.GetData()[0].GetId())
	assert.Equal(t, name1, output.GetData()[0].GetName())
	assert.Equal(t, id2, output.GetData()[1].GetId())
	assert.Equal(t, name2, output.GetData()[1].GetName())
}

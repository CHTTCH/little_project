package patient

import (
	mockRepo "github.com/CHTTCH/little_project/backend/adapter/patient/mock_repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePatientSucceed(t *testing.T) {
	id := "1"
	name := "小明"

	repo := mockRepo.NewMockPatientRepository()
	input := NewCreatePatientInput(id, name)

	output := CreatePatient(repo, input)

	patients, _ := repo.FindAll()

	assert.Equal(t, id, patients[0].GetId())
	assert.Equal(t, name, patients[0].GetName())
	assert.Equal(t, id, output.GetId())
	assert.Equal(t, true, output.GetResult())
}

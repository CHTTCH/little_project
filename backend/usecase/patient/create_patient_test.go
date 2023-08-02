package patient

import (
	customizeError "github.com/CHTTCH/little_project/backend/adapter/error"
	"github.com/CHTTCH/little_project/backend/entity/patient"
	mockRepo "github.com/CHTTCH/little_project/backend/test/mock_repository/patient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAPatientThenSuccess(t *testing.T) {
	assert := assert.New(t)

	repo := mockRepo.NewMockPatientRepository()
	input := NewCreatePatientInput("id", "name")

	output := CreatePatient(repo, input)

	assert.Equal(true, output.GetResult())
	assert.Equal("id", output.GetId())
	assert.Equal("", output.GetMessage())

	patient, _ := repo.FindById("id")
	assert.Equal("id", patient.GetId())
	assert.Equal("name", patient.GetName())
}

type mockRepositoryForSavingFailed struct{}

func (r *mockRepositoryForSavingFailed) Save(p *patient.Patient) error {
	return &customizeError.CustomizeError{Message: "repo save failed"}
}

func (r *mockRepositoryForSavingFailed) FindById(id string) (*patient.Patient, error) {
	return nil, nil
}

func (r *mockRepositoryForSavingFailed) FindAll() ([]patient.Patient, error) {
	return nil, nil
}

func TestRepoSavePatientFailedThenCreatePatientFailed(t *testing.T) {
	assert := assert.New(t)

	input := NewCreatePatientInput("id", "name")

	output := CreatePatient(&mockRepositoryForSavingFailed{}, input)

	assert.Equal(false, output.GetResult())
	assert.Equal("save failed", output.GetMessage())
	assert.Equal("id", output.GetId())
}

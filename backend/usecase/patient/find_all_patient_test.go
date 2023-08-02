package patient

import (
	customizeError "github.com/CHTTCH/little_project/backend/adapter/error"
	"github.com/CHTTCH/little_project/backend/entity/patient"
	mockPatientRepo "github.com/CHTTCH/little_project/backend/test/mock_repository/patient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllPatientsThenSuccess(t *testing.T) {
	assert := assert.New(t)

	repo := mockPatientRepo.NewMockPatientRepository()

	p1 := patient.NewPatient("1", "小明")
	p2 := patient.NewPatient("2", "小羊")
	repo.Save(p1)
	repo.Save(p2)

	output := FindAllPatients(repo)

	assert.Equal(true, output.GetResult())
	assert.Equal(2, len(output.GetData()))
	assert.Equal(*p1, output.GetData()[0])
	assert.Equal(*p2, output.GetData()[1])
	assert.Equal("", output.GetMessage())
}

type mockToFailedFindAllRepository struct{}

func (r *mockToFailedFindAllRepository) Save(p *patient.Patient) error {
	return nil
}

func (r *mockToFailedFindAllRepository) FindById(id string) (*patient.Patient, error) {
	return nil, nil
}

func (r *mockToFailedFindAllRepository) FindAll() ([]patient.Patient, error) {
	return nil, &customizeError.CustomizeError{Message: "repo find all failed"}
}

func TestRepoFindAllFailedThenFindAllPatientsFailed(t *testing.T) {
	assert := assert.New(t)

	output := FindAllPatients(&mockToFailedFindAllRepository{})

	assert.Equal(false, output.GetResult())
	assert.Equal("find all failed", output.GetMessage())
	assert.Equal(0, len(output.GetData()))
}

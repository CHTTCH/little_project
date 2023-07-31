package mock_repository

import (
	customizeError "github.com/CHTTCH/little_project/backend/adapter/error"
	"github.com/CHTTCH/little_project/backend/entity/patient"
)

type MockPatientRepository struct {
	patientList []patient.Patient
}

func NewMockPatientRepository() *MockPatientRepository {
	return &MockPatientRepository{patientList: []patient.Patient{}}
}

func (r *MockPatientRepository) Save(p *patient.Patient) error {
	if p.GetOrderId() != 0 {
		for index, patient := range r.patientList {
			if patient.GetId() == p.GetId() {
				r.patientList[index] = *p
				return nil
			}
		}
	} else {
		r.patientList = append(r.patientList, *p)
	}

	return nil
}

func (r *MockPatientRepository) FindAll() ([]patient.Patient, error) {
	return r.patientList, nil
}

func (r *MockPatientRepository) FindById(id string) (*patient.Patient, error) {
	for _, patient := range r.patientList {
		if patient.GetId() == id {
			return &patient, nil
		}
	}

	return nil, &customizeError.CustoMizeError{Message: "patient not found"}
}

package mock_repository

import (
	customizeError "github.com/CHTTCH/little_project/backend/adapter/error"
	"github.com/CHTTCH/little_project/backend/entity/patient"
)

type MockPatientRepository struct {
	PatientList []patient.Patient
}

func (r *MockPatientRepository) Save(p *patient.Patient) error {
	if p.OrderId != 0 {
		for index, patient := range r.PatientList {
			if patient.GetId() == p.GetId() {
				r.PatientList[index] = *p
				return nil
			}
		}
	} else {
		r.PatientList = append(r.PatientList, *p)
	}

	return nil
}

func (r *MockPatientRepository) FindAll() ([]patient.Patient, error) {
	return r.PatientList, nil
}

func (r *MockPatientRepository) FindById(id string) (*patient.Patient, error) {
	for _, patient := range r.PatientList {
		if patient.GetId() == id {
			return &patient, nil
		}
	}

	return nil, &customizeError.CustoMizeError{Message: "patient not found"}
}

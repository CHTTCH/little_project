package mock_repository

import (
	"github.com/CHTTCH/little_project/backend/entity/patient"
)

type MockPatientRepository struct {
	PatientList []patient.Patient
}

func (r *MockPatientRepository) Save(p patient.Patient) error{
	r.PatientList = append(r.PatientList, p)
	return nil
}

func (r *MockPatientRepository) FindAll() ([]patient.Patient, error) {
	return r.PatientList, nil
}
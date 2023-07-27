package patient

import (
	"github.com/CHTTCH/little_project/backend/entity/patient"
)

type MockPatientRepository struct {
	PatientList []patient.Patient
}

func (r *MockPatientRepository) Save(p patient.Patient) {
	r.PatientList = append(r.PatientList, p)
}

func (r *MockPatientRepository) GetPatientList() []patient.Patient {
	return r.PatientList
}
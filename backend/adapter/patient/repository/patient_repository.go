package repository

import (
	"github.com/CHTTCH/little_project/backend/entity/patient"
	"gorm.io/gorm"
)

type PatientRepository struct {
	dB *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{dB: db}
}

func (repo *PatientRepository) Save(p *patient.Patient) error {
	if err := repo.dB.AutoMigrate(new(patient.Patient)); err != nil {
		return err
	}

	if err := repo.dB.Save(p).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PatientRepository) FindAll() ([]patient.Patient, error) {
	patients := new([]patient.Patient)

	if err := repo.dB.Find(patients).Error; err != nil {
		return nil, err
	}

	return *patients, nil
}

func (repo *PatientRepository) FindById(id string) (*patient.Patient, error) {
	patient := &patient.Patient{}

	if err := repo.dB.First(patient, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return patient, nil
}

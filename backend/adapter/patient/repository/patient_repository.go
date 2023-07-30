package repository

import (
	"github.com/CHTTCH/little_project/backend/entity/patient"
	"gorm.io/gorm"
)

type PatientRepository struct {
	DB *gorm.DB
}

func (repo *PatientRepository) Save(p *patient.Patient) error {
	if err := repo.DB.AutoMigrate(new(patient.Patient)); err != nil {
		return err
	}

	if err := repo.DB.Save(p).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PatientRepository) FindAll() ([]patient.Patient, error) {
	patients := new([]patient.Patient)

	if err := repo.DB.Find(patients).Error; err != nil {
		return nil, err
	}

	return *patients, nil
}

func (repo *PatientRepository) FindById(id string) (*patient.Patient, error) {
	patient := &patient.Patient{}

	if err := repo.DB.First(patient, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return patient, nil
}

package repository

import (
	"github.com/CHTTCH/little_project/backend/entity/patient"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo *Repository) Save(p patient.Patient) error {
	if err:= repo.DB.AutoMigrate(new(patient.Patient)); err != nil {
        return err
    }

	if err := repo.DB.Save(p).Error; err != nil {
        return err
    }

	return nil
}

func (repo *Repository) FindAll() ([]patient.Patient, error) {
	patients := new([]patient.Patient)

	if err := repo.DB.Find(patients).Error; err != nil {
		return nil, err
	}

	return *patients, nil
}

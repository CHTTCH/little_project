package scripts

import (
	customizeError "github.com/CHTTCH/little_project/backend/adapter/error"
	entityPatient "github.com/CHTTCH/little_project/backend/entity/patient"
	usecasePatient "github.com/CHTTCH/little_project/backend/usecase/patient"
	"github.com/CHTTCH/little_project/backend/usecase/repository"
	"gorm.io/gorm"
)

var patients = []map[string]string{
	{"Id": "1", "Name": "小明"},
	{"Id": "2", "Name": "小蔡"},
	{"Id": "3", "Name": "小楊"},
	{"Id": "4", "Name": "小李"},
	{"Id": "5", "Name": "小高"},
}

func CreateFivePatients(db *gorm.DB, repo repository.Repository[entityPatient.Patient, string]) error {
	if hasTable := db.Migrator().HasTable("patients"); hasTable {
		return nil
	} else {
		for _, patient := range patients {
			output := usecasePatient.CreatePatient(repo, usecasePatient.NewCreatePatientInput(patient["Id"], patient["Name"]))
			if !output.Result {
				return &customizeError.CustomizeError{Message: "create five patients failed"}
			}
		}
	}
	return nil
}

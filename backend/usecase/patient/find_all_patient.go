package patient

import (
	"github.com/CHTTCH/little_project/backend/entity/patient"
	queryOutput "github.com/CHTTCH/little_project/backend/usecase/query_output"
	"github.com/CHTTCH/little_project/backend/usecase/repository"
)

func FindAllPatient(repo repository.Repository[patient.Patient, string]) queryOutput.QueryOutput[patient.Patient] {
	data, err := repo.FindAll()

	if err != nil {
		return queryOutput.QueryOutput[patient.Patient]{Result: false, Message: "find all failed"}
	}

	return queryOutput.QueryOutput[patient.Patient]{Data: data, Result: true}
}

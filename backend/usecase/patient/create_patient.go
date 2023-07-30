package patient

import (
	"github.com/CHTTCH/little_project/backend/usecase/repository"
	"github.com/CHTTCH/little_project/backend/entity/patient"
	commandOutput "github.com/CHTTCH/little_project/backend/usecase/command_output"
)

func CreatePatient(repo repository.Repository[patient.Patient], input CreatePatientInput) commandOutput.CommandOutput{
	p := patient.Patient{
		Id:      input.GetId(),
		Name:    input.GetName(),
	}

	err := repo.Save(p)

	if err != nil {
		return 	commandOutput.CommandOutput{ Id: input.GetId(), Result: false, Message: "save failed" }
	}
	
	return commandOutput.CommandOutput{ Id: input.GetId(), Result: true }
}

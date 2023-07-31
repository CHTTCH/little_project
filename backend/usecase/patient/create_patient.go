package patient

import (
	"github.com/CHTTCH/little_project/backend/entity/patient"
	commandOutput "github.com/CHTTCH/little_project/backend/usecase/command_output"
	"github.com/CHTTCH/little_project/backend/usecase/repository"
)

func CreatePatient(repo repository.Repository[patient.Patient, string], input *CreatePatientInput) commandOutput.CommandOutput {
	p := patient.NewPatient(input.GetId(), input.GetName())

	err := repo.Save(p)

	if err != nil {
		return commandOutput.CommandOutput{Id: input.GetId(), Result: false, Message: "save failed"}
	}

	return commandOutput.CommandOutput{Id: input.GetId(), Result: true}
}

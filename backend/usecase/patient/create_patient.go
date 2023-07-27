package patient

import (
	"github.com/CHTTCH/little_project/backend/entity/patient"
)

func CreatePatient(repo Repository[patient.Patient], input *CreatePatientInput, output *CreatePatientOutput) {
	p := patient.Patient{
		Id:      input.GetId(),
		Name:    input.GetName(),
		OrderId: input.GetOrderId(),
	}

	repo.Save(p)

	output.SetId(input.GetId())
}

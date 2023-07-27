package patient

type CreatePatientOutput struct {
	id string
}

func (i *CreatePatientOutput) SetId(id string) {
	i.id = id
}

func (i *CreatePatientOutput) GetId() string {
	return i.id
}

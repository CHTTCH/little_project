package patient

type CreatePatientInput struct {
	id      string
	name    string
}

func (i *CreatePatientInput) SetId(id string) {
	i.id = id
}

func (i *CreatePatientInput) SetName(name string) {
	i.name = name
}

func (i *CreatePatientInput) GetId() string {
	return i.id
}

func (i *CreatePatientInput) GetName() string {
	return i.name
}

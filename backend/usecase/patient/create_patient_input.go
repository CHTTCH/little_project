package patient

type CreatePatientInput struct {
	id      string
	name    string
	orderId string
}

func (i *CreatePatientInput) GetId() string {
	return i.id
}

func (i *CreatePatientInput) GetName() string {
	return i.name
}

func (i *CreatePatientInput) GetOrderId() string {
	return i.orderId
}

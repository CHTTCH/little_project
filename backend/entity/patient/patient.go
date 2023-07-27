package patient

type Patient struct {
	Id string
	Name string
	OrderId string
}

func (p Patient) GetId() string {
	return p.Id
}

func (p Patient) GetName() string {
	return p.Name
}

func (p Patient) GetOrderId() string {
	return p.OrderId
}
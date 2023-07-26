package patient

type Patient struct {
	Id string
	Name string
	OrderId string
}

func (p Patient) getId() string {
	return p.Id
}

func (p Patient) getName() string {
	return p.Name
}

func (p Patient) getOrderId() string {
	return p.OrderId
}
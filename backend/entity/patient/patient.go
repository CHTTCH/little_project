package patient

type Patient struct {
	Id      string
	Name    string
	OrderId int
}

func (p *Patient) SetId(id string) {
	p.Id = id
}

func (p Patient) GetId() string {
	return p.Id
}

func (p *Patient) SetName(name string) {
	p.Name = name
}

func (p Patient) GetName() string {
	return p.Name
}

func (p *Patient) SetOrderId(orderId int) {
	p.OrderId = orderId
}

func (p Patient) GetOrderId() int {
	return p.OrderId
}

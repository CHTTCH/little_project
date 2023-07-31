package order

type EditOrderInput struct {
	orderId         int
	modifiedMessage string
}

func NewEditOrderInput(orderId int, modifiedMessage string) *EditOrderInput {
	return &EditOrderInput{orderId: orderId, modifiedMessage: modifiedMessage}
}

func (i *EditOrderInput) SetOrderId(orderId int) {
	i.orderId = orderId
}

func (i *EditOrderInput) GetOrderId() int {
	return i.orderId
}

func (i *EditOrderInput) SetModifiedMessage(modifiedMessage string) {
	i.modifiedMessage = modifiedMessage
}

func (i *EditOrderInput) GetModifiedMessage() string {
	return i.modifiedMessage
}

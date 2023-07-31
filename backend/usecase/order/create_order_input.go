package order

type CreateOrderInput struct {
	patientId string
	message   string
}

func NewCreateOrderInput(patientId, message string) *CreateOrderInput {
	return &CreateOrderInput{patientId: patientId, message: message}
}

func (i *CreateOrderInput) SetPatientId(patientId string) {
	i.patientId = patientId
}

func (i *CreateOrderInput) GetPatientId() string {
	return i.patientId
}

func (i *CreateOrderInput) SetMessage(message string) {
	i.message = message
}

func (i *CreateOrderInput) GetMessage() string {
	return i.message
}

package order

type Order struct {
	Id      int `gorm:"primary_key;auto_increment"`
	Message string
}

func NewOrder(message string) *Order {
	return &Order{Message: message}
}

func (o Order) GetId() int {
	return o.Id
}

func (o *Order) SetId(id int) {
	o.Id = id
}

func (o Order) GetMessage() string {
	return o.Message
}

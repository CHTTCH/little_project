package query_output

import (
	"github.com/CHTTCH/little_project/backend/entity/patient"
	"github.com/CHTTCH/little_project/backend/entity/order"
)

type QueryOutput[T patient.Patient | order.Order] struct {
	Data    []T
	Result  bool
	Message string
}

func (q QueryOutput[T]) GetData() []T {
	return q.Data
}

func (q QueryOutput[T]) GetResult() bool {
	return q.Result
}

func (q QueryOutput[T]) GetMessage() string {
	return q.Message
}

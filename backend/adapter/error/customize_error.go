package error

import "fmt"

type CustoMizeError struct {
	Message string
	Detail  string
}

func (e *CustoMizeError) Error() string {
	return fmt.Sprintf("error: %s, detail: %s", e.Message, e.Detail)
}

package error

import "fmt"

type CustomizeError struct {
	Message string
}

func (e *CustomizeError) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}

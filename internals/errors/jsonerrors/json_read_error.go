package jsonerrors

import "fmt"

type JSONReadError struct {
	Err error
}

func (m *JSONReadError) Error() string {
	return fmt.Sprintf("Encountered JSON read error: %v", m.Err)
}

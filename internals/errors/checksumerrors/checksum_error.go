package checksumerrors

import "fmt"

type ChecksumError struct {
	Err error
}

func (m *ChecksumError) Error() string {
	return fmt.Sprintf("Encountered checksum error: %v", m.Err)
}

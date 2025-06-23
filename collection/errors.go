package collection

import "fmt"

type IndexOutOfBoundsError struct {
	Index int
	Size  int
}

func (e *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("IndexOutOfBoundsException: Index: %d, Size: %d", e.Index, e.Size)
}

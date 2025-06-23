package collection

type List[T any] interface {
	Add(element T)

	AddAt(index int, element T) error

	Clear()

	Get(index int) (T, error)

	Remove(index int) (T, error)

	Set(index int, element T) error

	Size() int

	IsEmpty() bool

	ToSlice() []T
}

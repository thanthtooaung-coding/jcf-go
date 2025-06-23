package collection

type ArrayList[T any] struct {
	elements []T
}

func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{elements: make([]T, 0)}
}

func (list *ArrayList[T]) Add(element T) {
	list.elements = append(list.elements, element)
}

func (list *ArrayList[T]) AddAt(index int, element T) error {
	if index < 0 || index > list.Size() {
		return &IndexOutOfBoundsError{Index: index, Size: list.Size()}
	}

	list.elements = append(list.elements, *new(T))
	copy(list.elements[index+1:], list.elements[index:]) // Shift
	list.elements[index] = element
	return nil
}

func (list *ArrayList[T]) Clear() {
	list.elements = make([]T, 0)
}

func (list *ArrayList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= list.Size() {
		return zero, &IndexOutOfBoundsError{Index: index, Size: list.Size()}
	}
	return list.elements[index], nil
}

func (list *ArrayList[T]) Remove(index int) (T, error) {
	var zero T
	if index < 0 || index > list.Size() {
		return zero, &IndexOutOfBoundsError{Index: index, Size: list.Size()}
	}

	removeElement := list.elements[index]
	list.elements = append(list.elements[:index], list.elements[index+1:]...)
	return removeElement, nil
}

func (list *ArrayList[T]) Set(index int, element T) error {
	if index < 0 || index >= list.Size() {
		return &IndexOutOfBoundsError{Index: index, Size: list.Size()}
	}
	list.elements[index] = element
	return nil
}

func (list *ArrayList[T]) Size() int {
	return len(list.elements)
}

func (list *ArrayList[T]) IsEmpty() bool {
	return list.Size() == 0
}

func (list *ArrayList[T]) ToSlice() []T {
	dataCopy := make([]T, list.Size())
	copy(dataCopy, list.elements)
	return dataCopy
}

type ComparableArrayList[T comparable] struct {
	ArrayList[T]
}

func NewComparableArrayList[T comparable]() *ComparableArrayList[T] {
	return &ComparableArrayList[T]{
		ArrayList[T]{elements: make([]T, 0)},
	}
}

func (list *ComparableArrayList[T]) Contains(element T) bool {
	for _, e := range list.elements {
		if e == element {
			return true
		}
	}
	return false
}

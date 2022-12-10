package common

import "fmt"

type Pair[T any] struct {
	Left  T
	Right T
}

func NewPair[T any](left, right T) *Pair[T] {
	return &Pair[T]{
		Left:  left,
		Right: right,
	}
}

func (p *Pair[T]) String() string {
	return fmt.Sprintf("{%v, %v}", p.Left, p.Right)
}

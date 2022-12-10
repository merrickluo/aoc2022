package common

import (
	"bytes"
	"fmt"
)

type Tree[T any] struct {
	Value    T
	Children []*Tree[T]
}

func (t *Tree[T]) String() string {
	buf := bytes.Buffer{}

	buf.WriteString(fmt.Sprintf("%v", t.Value))
	for _, child := range t.Children {
		buf.WriteString(child.String())
	}

	return buf.String()
}

func (t *Tree[T]) Walk(walkFunc func(t *Tree[T])) {
	walkFunc(t)

	for _, child := range t.Children {
		child.Walk(walkFunc)
	}
}

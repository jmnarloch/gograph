package list

import (
	"container/list"
)

type ListIterator struct {
	list    *list.List
	current *list.Element
}

func New(list *list.List) *ListIterator {
	return &ListIterator{
		list:    list,
		current: list.Front(),
	}
}

func (l *ListIterator) HasNext() bool {

	return l.current != nil
}

func (l *ListIterator) Next() interface{} {

	if l.current == nil {
		return nil
	}

	val := l.current
	l.current = l.current.Next()
	return val.Value
}

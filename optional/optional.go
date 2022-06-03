package optional

import "errors"

type Optional[T any] interface {
	Get() (*T, error)
	IsPresent() bool
	IfPresent(consumer func(val T)) Optional[T]
	OrElse(other T) T
	OrElseGet(otherFunc func() T) T
	Filter(predicate func(val T) bool) Optional[T]
}

type optional[T any] struct {
	val *T
}

func (optional *optional[T]) Get() (*T, error) {
	if optional.IsPresent() {
		return optional.val, nil
	}
	return nil, errors.New("no value present")
}

func (optional *optional[T]) IsPresent() bool {
	return optional.val != nil
}

func (optional *optional[T]) IfPresent(consumer func(val T)) Optional[T] {
	val, _ := optional.Get()
	if val != nil {
		consumer(*val)
	}
	return optional
}

func (optional *optional[T]) OrElse(other T) T {
	val, _ := optional.Get()
	if val != nil {
		return *val
	}
	return other
}

func (optional *optional[T]) OrElseGet(otherFunc func() T) T {
	val, _ := optional.Get()
	if val != nil {
		return *val
	}
	return otherFunc()
}

func (optional *optional[T]) Filter(predicate func(val T) bool) Optional[T] {
	val, _ := optional.Get()
	if val != nil && predicate(*val) {
		return Of[T](*val)
	}
	return Empty[T]()
}

func Map[T any, S any](optional Optional[T], mapper func(val T) S) Optional[S] {
	val, _ := optional.Get()
	if val != nil {
		result := mapper(*val)
		return Of[S](result)
	}
	return Empty[S]()
}

func Of[T any](val T) Optional[T] {
	return &optional[T]{&val}
}

func Empty[T any]() Optional[T] {
	return &optional[T]{}
}

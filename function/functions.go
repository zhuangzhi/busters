package function

import (
	"errors"
	"sync"
)

var (
	// ErrorAlreadyInvoked ... error for already invoked
	ErrorAlreadyInvoked = errors.New("already invoked")

	// ErrorEmptyFunction ... error for empty function
	ErrorEmptyFunction = errors.New("empty function")
)

// Consumer is a function that accepts a single input argument and returns no result.
type Consumer[T any] func(T)

// BiConsumer is a function that accepts two input arguments and returns no result.
type BiConsumer[T any, R any] func(T, R)

type TripleConsumer[T any, R any, U any] func(T, R, U)

// Supplier is a function that returns a result.
type Supplier[T any] func() (T, error)

// BiSupplier is a function that returns two results.
type BiSupplier[T any, R any] func() (T, R, error)

// Function is a function that accepts one argument and produces a result.
type Function[T any, R any] func(T) (R, error)

// BiFunction is a function that accepts two arguments and produces a result.
type BiFunction[T any, R any, U any] func(T, R) (U, error)

type Callable[T any] func() (T, error)

type Predicate[T any] func(T) (bool, error)

func (c Consumer[T]) Apply(t T) {
	if c != nil {
		c(t)
	}
}

func (c Consumer[T]) AndThen(next func(T)) Consumer[T] {
	return func(t T) {
		c.Apply(t)
		Consumer[T](next).Apply(t)
	}
}

func (c Consumer[T]) Once() Consumer[T] {
	once := sync.Once{}
	return func(t T) {
		once.Do(func() {
			c.Apply(t)
		})
	}
}

func (c BiConsumer[T, R]) Apply(t T, r R) {
	if c != nil {
		c(t, r)
	}
}

func (c BiConsumer[T, R]) AndThen(next func(T, R)) BiConsumer[T, R] {
	return func(t T, r R) {
		c.Apply(t, r)
		BiConsumer[T, R](next).Apply(t, r)
	}
}

func (c BiConsumer[T, R]) Once() BiConsumer[T, R] {
	once := sync.Once{}
	return func(t T, r R) {
		once.Do(func() {
			c.Apply(t, r)
		})
	}
}

func (c TripleConsumer[T, R, U]) Apply(t T, r R, u U) {
	if c != nil {
		c(t, r, u)
	}
}

func (c TripleConsumer[T, R, U]) AndThen(next func(T, R, U)) TripleConsumer[T, R, U] {
	return func(t T, r R, u U) {
		c.Apply(t, r, u)
		TripleConsumer[T, R, U](next).Apply(t, r, u)
	}
}

func (c TripleConsumer[T, R, U]) Once() TripleConsumer[T, R, U] {
	once := sync.Once{}
	return func(t T, r R, u U) {
		once.Do(func() {
			c.Apply(t, r, u)
		})
	}
}

func (p Supplier[T]) Supply() (T, error) {
	if p != nil {
		return p()
	}
	return *new(T), ErrorEmptyFunction
}

func (p BiSupplier[T, R]) Supply() (T, R, error) {
	if p != nil {
		return p()
	}
	return *new(T), *new(R), ErrorEmptyFunction
}

func (f Function[T, R]) Apply(t T) (R, error) {
	if f != nil {
		return f(t)
	}
	return *new(R), ErrorEmptyFunction
}

func (f BiFunction[T, R, U]) Apply(t T, r R) (U, error) {
	if f != nil {
		return f(t, r)
	}
	return *new(U), ErrorEmptyFunction
}

func (c Callable[T]) Call() (T, error) {
	if c != nil {
		return c()
	}
	return *new(T), ErrorEmptyFunction
}

func (c Callable[T]) AndThen(next func() (T, error)) Callable[T] {
	return func() (T, error) {
		t, err := c.Call()
		if err != nil {
			return t, err
		}
		return next()
	}
}

func (c Callable[T]) Once() Callable[T] {
	once := sync.Once{}
	var t T
	var err error = ErrorAlreadyInvoked
	return func() (T, error) {
		once.Do(func() {
			t, err = c.Call()
		})
		return t, err
	}
}

func (p Predicate[T]) Test(t T) (bool, error) {
	if p != nil {
		return p(t)
	}
	return false, ErrorEmptyFunction
}

func (p Predicate[T]) And(next func(T) (bool, error)) Predicate[T] {
	return func(t T) (bool, error) {
		b, err := p.Test(t)
		if err != nil {
			return b, err
		}
		if !b {
			return b, nil
		}
		return next(t)
	}
}

func (p Predicate[T]) Or(next func(T) (bool, error)) Predicate[T] {
	return func(t T) (bool, error) {
		b, err := p.Test(t)
		if err != nil {
			return b, err
		}
		if b {
			return b, nil
		}
		return next(t)
	}
}

type Result[T any] struct {
	Value T
	Error error
}

type OnResult[T any] Consumer[Result[T]]

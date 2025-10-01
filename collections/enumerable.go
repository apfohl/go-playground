package collections

import "iter"

type Enumerable[T any] iter.Seq[T]

func (e Enumerable[T]) Take(n int) Enumerable[T] {
	return func(yield func(T) bool) {
		counter := 0
		for value := range e {
			if counter >= n {
				break
			}

			if !yield(value) {
				break
			}

			counter++
		}
	}
}

func (e Enumerable[T]) Filter(predicate func(T) bool) Enumerable[T] {
	return func(yield func(T) bool) {
		for value := range e {
			if predicate(value) {
				if !yield(value) {
					break
				}
			}
		}
	}
}

func (e Enumerable[T]) ForEach(f func(T)) {
	for value := range e {
		f(value)
	}
}

func Flatten[T any](sources Enumerable[Enumerable[T]]) Enumerable[T] {
	return func(yield func(T) bool) {
		for source := range sources {
			for value := range source {
				if !yield(value) {
					break
				}
			}
		}
	}
}

func FlatMap[T, R any](sources Enumerable[Enumerable[T]], mapper func(T) R) Enumerable[R] {
	return func(yield func(R) bool) {
		for source := range sources {
			for value := range source {
				if !yield(mapper(value)) {
					break
				}
			}
		}
	}
}

func Reduce[T, R any](source Enumerable[T], accumulator R, aggregator func(R, T) R) R {
	for value := range source {
		accumulator = aggregator(accumulator, value)
	}

	return accumulator
}

func Map[T, U any](e Enumerable[T], mapping func(T) U) Enumerable[U] {
	return func(yield func(U) bool) {
		for value := range e {
			if !yield(mapping(value)) {
				break
			}
		}
	}
}

package gofn

// Reduce reduces a slice of values using given reducer into a single value.
//
//  Reduce(0, []int{1, 2, 3}, func(acc int, v int) int { return acc + v } == (0 + 1 + 2 + 3)
func Reduce[T any, R any](initial R, slices []T, reducer func(R, T) R) R {
	for _, v := range slices {
		initial = reducer(initial, v)
	}

	return initial
}

// Map maps a slice of values using given mapper into a new slice.
func Map[T any, R any](slices []T, mapper func(T) R) []R {
	return Reduce(make([]R, 0, len(slices)), slices, func(acc []R, v T) []R {
		return append(acc, mapper(v))
	})
}

// Filter filters a slice of values using given predicate into a new slice.
func Filter[T any](slices []T, predicate func(T) bool) []T {
	return Reduce(make([]T, 0), slices, func(acc []T, v T) []T {
		if predicate(v) {
			return append(acc, v)
		}

		return acc
	})
}

// Option is a function that applies an option to a given value T.
type Option[T any] func(T)

// ApplyOptions applies all given options to a given value T.
func ApplyOptions[T any](initial T, opts ...Option[T]) {
	Reduce(initial, opts, func(acc T, opt Option[T]) T {
		opt(acc)
		return acc
	})
}

// Decorate wraps function f with given decorators.
//
//  Decorate(f, g, h, i) == i(h(g(f)))
func Decorate[F any](f F, decorators ...func(F) F) F {
	return Reduce(f, decorators, func(acc F, decorator func(F) F) F {
		return decorator(acc)
	})
}

// ReversedDecorate wraps function f with given decorators in reverse order.
//
// 	ReversedDecorate(f, g, h, i) == g(h(i(f)))
func ReversedDecorate[F any](f F, decorators ...func(F) F) F {
	return Decorate(f, Reverse(decorators)...)
}

// Reverse reverses a slice of values.
func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

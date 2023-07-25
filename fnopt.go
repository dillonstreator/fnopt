package fnopt

// OptFn[T] is a function type used to enable generic `T` creation with options.
type OptFn[T any] func(cfg *T)

// New[T] creates a new struct of type `T` and modifies it by applying the provided
// option functions to it.
//
// This is useful for functional options in methods.
func New[T any](optFns ...OptFn[T]) *T {
	t, _ := NewE(optFnsToOptFnEs(optFns...)...)
	return t
}

// NewFrom[T] modifies an existing struct of type `T` by applying the provided option functions to it.
//
// This is useful for constructing structs with optional functions.
func NewFrom[T any](t *T, optFns ...OptFn[T]) {
	NewFromE(t, optFnsToOptFnEs(optFns...)...)
}

// OptFnE[T] is a function type used to enable generic `T` creation with error handling.
type OptFnE[T any] func(cfg *T) error

// NewE[T] creates a new struct of type `T` and modifies it by applying the provided
// option functions to it with error handling.
//
// This is useful for functional options in methods.
func NewE[T any](optFns ...OptFnE[T]) (*T, error) {
	cfg := new(T)

	err := NewFromE(cfg, optFns...)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// NewFromE[T] modifies an existing struct of type `T` by applying the provided option functions to it with error handling.
//
// This is useful for constructing structs with optional functions.
func NewFromE[T any](t *T, optFns ...OptFnE[T]) error {
	for _, optFn := range optFns {
		err := optFn(t)
		if err != nil {
			return err
		}
	}

	return nil
}

func optFnsToOptFnEs[T any](fns ...OptFn[T]) []OptFnE[T] {
	out := make([]OptFnE[T], len(fns))
	for i, fn := range fns {
		_fn := fn
		out[i] = func(cfg *T) error {
			_fn(cfg)
			return nil
		}
	}
	return out
}

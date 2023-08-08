package fnopt

// OptFn is a function type used to define functional options for mutating the generic type `T`
type OptFn[T any] func(cfg *T)

// New creates a new struct pointer of type `T` and modifies it by applying
// the provided option functions to it.
func New[T any](optFns ...OptFn[T]) *T {
	t, _ := NewE(optFnsToOptFnEs(optFns...)...)
	return t
}

// From modifies an existing struct pointer of type `T` by applying the
// provided option functions to it.
func From[T any](t *T, optFns ...OptFn[T]) {
	FromE(t, optFnsToOptFnEs(optFns...)...)
}

// OptFnE[T] is a function type used to enable generic `T` creation with error handling.
type OptFnE[T any] func(cfg *T) error

// NewE creates a new struct pointer of type `T` and modifies it by applying
// the provided option functions to it and propagating any errors.
func NewE[T any](optFns ...OptFnE[T]) (*T, error) {
	cfg := new(T)

	err := FromE(cfg, optFns...)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// FromE modifies an existing struct pointer of type `T` by applying the
// provided option functions to it and propagating any errors.
func FromE[T any](t *T, optFns ...OptFnE[T]) error {
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

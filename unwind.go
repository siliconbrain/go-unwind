package unwind

// Handler is a stack unwind handler function
type Handler = func(reason interface{})

// DoWithHandler executes fn and calls h only when fn returns abnormally.
//
// The handler receives the panic's reason as a parameter. If the reason is `nil`, the goroutine was terminated via `runtime.Goexit`.
func DoWithHandler(h Handler, fn func()) {
	done := false
	defer func() {
		if !done {
			h(recover())
		}
	}()
	fn()
	done = true
}

// DoWithHandler1 executes fn and returns its return value. It calls h only when fn returns abnormally.
//
// The handler receives the panic's reason as a parameter. If the reason is `nil`, the goroutine was terminated via `runtime.Goexit`.
func DoWithHandler1[T any](h Handler, fn func() T) T {
	done := false
	defer func() {
		if !done {
			h(recover())
		}
	}()
	r := fn()
	done = true
	return r
}

// DoWithHandler2 executes fn and returns its return values. It calls h only when fn returns abnormally.
//
// The handler receives the panic's reason as a parameter. If the reason is `nil`, the goroutine was terminated via `runtime.Goexit`.
func DoWithHandler2[T1, T2 any](h Handler, fn func() (T1, T2)) (T1, T2) {
	done := false
	defer func() {
		if !done {
			h(recover())
		}
	}()
	r1, r2 := fn()
	done = true
	return r1, r2
}

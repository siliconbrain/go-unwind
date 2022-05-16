package unwind

// Handler is a stack unwind handler function
type Handler func(reason any)

// Go executes the specified function in a goroutine and calls the handler iff the function did not return normally.
// If the reason passed to the handler is `nil` the function was interrupted by `runtime.Goexit`; otherwise it panicked.
func (h Handler) Go(fn func()) {
	go func() {
		done := false
		defer func() {
			if !done {
				h(recover())
			}
		}()
		fn()
		done = true
	}()
}

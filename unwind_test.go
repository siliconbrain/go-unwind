package unwind

import (
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_Go(t *testing.T) {
	testCases := map[string]struct {
		fn     func()
		called bool
		reason interface{}
		waiter waiter
	}{
		"ordinary": {
			fn:     func() {},
			called: false,
			reason: nil,
			waiter: newWaiter(1),
		},
		"panic": {
			fn:     func() { panic(42) },
			called: true,
			reason: 42,
			waiter: newWaiter(2),
		},
		"goexit": {
			fn:     func() { runtime.Goexit() },
			called: true,
			reason: nil,
			waiter: newWaiter(2),
		},
	}
	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			var called bool = false
			var reason any = nil
			Handler(func(r any) {
				called = true
				reason = r
				testCase.waiter.Done()
			}).Go(func() {
				defer testCase.waiter.Done()
				testCase.fn()
			})
			testCase.waiter.Wait()
			assert.Equal(t, testCase.called, called)
			assert.Equal(t, testCase.reason, reason)
		})
	}
}

type waiter interface {
	Done()
	Wait()
}

func newWaiter(c int) waiter {
	w := new(sync.WaitGroup)
	w.Add(c)
	return w
}

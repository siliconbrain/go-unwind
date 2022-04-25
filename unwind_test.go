package unwind

import (
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoWithHandler_WithPanic(t *testing.T) {
	handled := false

	h := func(reason interface{}) {
		assert.Equal(t, 42, reason)
		handled = true
	}
	DoWithHandler(h, func() {
		panic(42)
	})
	assert.True(t, handled)
}

func TestDoWithHandler_WithGoexit(t *testing.T) {
	handled := false

	var wg sync.WaitGroup
	wg.Add(1)

	h := func(reason interface{}) {
		assert.Nil(t, reason)
		handled = true
		wg.Done()
	}
	go DoWithHandler(h, func() {
		runtime.Goexit()
	})
	wg.Wait()
	assert.True(t, handled)
}

func TestDoWithHandler1(t *testing.T) {
	handled := false

	h := func(reason interface{}) {
		handled = true
	}
	result := DoWithHandler1(h, func() int {
		return 42
	})

	assert.Equal(t, 42, result)
	assert.False(t, handled)
}

func TestDoWithHandler1_WithPanic(t *testing.T) {
	handled := false

	h := func(reason interface{}) {
		assert.Equal(t, 42, reason)
		handled = true
	}
	result := DoWithHandler1(h, func() int {
		panic(42)
	})

	assert.Equal(t, 0, result)
	assert.True(t, handled)
}

func TestDoWithHandler2(t *testing.T) {
	handled := false

	h := func(reason interface{}) {
		handled = true
	}
	result1, result2 := DoWithHandler2(h, func() (int, string) {
		return 42, "towel"
	})

	assert.Equal(t, 42, result1)
	assert.Equal(t, "towel", result2)
	assert.False(t, handled)
}

func TestDoWithHandler2_WithPanic(t *testing.T) {
	handled := false

	h := func(reason interface{}) {
		assert.Equal(t, 42, reason)
		handled = true
	}
	result1, result2 := DoWithHandler2(h, func() (int, string) {
		panic(42)
	})

	assert.Equal(t, 0, result1)
	assert.Equal(t, "", result2)
	assert.True(t, handled)
}

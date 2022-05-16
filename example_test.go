package unwind_test

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/siliconbrain/go-unwind"
)

func TestReadmeExample(t *testing.T) {
	out := &strings.Builder{}

	var wg sync.WaitGroup
	wg.Add(1)
	h := unwind.Handler(func(reason interface{}) {
		if reason != nil {
			fmt.Fprintln(out, "panic with reason:", reason)
		} else {
			fmt.Fprintln(out, "unrecoverable unwind")
		}
		wg.Done()
	})
	h.Go(func() {
		panickyFunc(42)
		wg.Done()
	})
	wg.Wait()
	assert.Equal(t, "panic with reason: you forgot your towel\n", out.String())
}

func panickyFunc(v int) {
	if v == 42 {
		panic("you forgot your towel")
	}
}

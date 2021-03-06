# go-unwind

A tool for Golang to handle (irregular) stack unwinding.

## Usage

Add `go-unwind` to your project with
```sh
go get github.com/siliconbrain/go-unwind
```

Use the unwind handler in your code like
```go
import (
	"fmt"
	"sync"

	"github.com/siliconbrain/go-unwind"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	h := unwind.Handler(func(reason interface{}) {
		if reason != nil {
			fmt.Println("panic with reason:", reason)
		} else {
			fmt.Println("unrecoverable unwind")
		}
		wg.Done()
	})
	h.Go(func() {
		panickyFunc(42)
		wg.Done()
	})
	wg.Wait()
}

func panickyFunc(v int) {
	if v == 42 {
		panic("you forgot your towel")
	}
}
```

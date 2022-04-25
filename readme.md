# go-unwind

A tool for Golang to handle (irregular) stack unwinding.

## Usage

Add `go-unwind` to your project with
```sh
go get github.com/siliconbrain/go-unwind
```

Use the unwind handler in your code like
```go
import "github.com/siliconbrain/go-unwind"

func main() {
	h := func(reason interface{}) {
		if reason != nil {
			fmt.Println("panic with reason:", reason)
		} else {
			fmt.Println("unrecoverable unwind")
		}
	}
	unwind.DoWithHandler(h, func() {
		panickyFunc(42)
	})
}

func panickyFunc(v int) {
	if v == 42 {
		panic("you forgot your towel")
	}
}
```

There are also 2 convenience functions `DoWithHandler1` and `DoWithHandler2` that accept functions with 1 or 2 return values and return those return values (respectively).

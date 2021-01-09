# go-errors

[![GoDoc](https://pkg.go.dev/badge/github.com/VKoptev/go-errors)](https://pkg.go.dev/github.com/VKoptev/go-errors)
![golangci-lint](https://github.com/VKoptev/go-errors/workflows/golangci-lint/badge.svg)

Errors wrapper for more convenience :)


## Usage

It's just a wrapper to faster and convenient write go-code. You can use it according to your imagination :)

To create error with reason:
```go
ErrSomethingWentWrong := errors.WithReason("something went wrong")

return ErrSomethingWentWrong
```

To consider that this error is caused by another error:
```go
return ErrSomethingWentWrong.WithErr(err)
```

To include data that leads to error:
```go
return ErrSomethingWentWrong.WithX(data)
```

Bad data without reason and not being caused error (why not?):
```go
errors.WithX(data)
```

Error builder:

```go
package usefulnesses

import (
	"encoding/json"
	"net"

	"github.com/VKoptev/go-errors"
)

var (
	ErrSomethingWentWrong = errors.WithReason("something went wrong")
)

func Send(conn net.Conn, x interface{}) error {
	b, err := json.Marshal(x)
	if err != nil {
		return ErrSomethingWentWrong.WithErr(err).WithX(x)
	}

	if _, err := conn.Write(b); err != nil {
		return ErrSomethingWentWrong.WithErr(err).WithX(b)
	}

	return nil
}
```

### Helpers

There are 2 useful functions: OneOf and EachOf

OneOf returns true if err is matched with at least one of specified errors:
```go
var (
    errWentWrong = errors.WithReason("went wrong")
    errWentCompletelyWrong = errors.WithReason("went wrong")
)
...
if errors.OneOf(err, errWentWrong, errWentCompletelyWrong) {
    return
}
```

EachOf returns true if err is matched with each of specified errors:
```go
var errWentWrong = errors.WithReason("went wrong")

func a() error {
    if condition {
        return errWentWrong
    }
    return errWentWrong.WithErr(io.EOF)
}
...
if errors.EachOf(a(), errWentWrong, io.EOF) {
    return
}
```

### Aliases

To avoid importing builtin package two aliases is added:
* `Is(error, ...errors)` - alias for builtin function, but it has such an interface as OneOf and EachOf. **(!)** Only first item of variadic argument is used.
* `As(error, interface{})`

### Links
[Code example](https://github.com/VKoptev/go-errors).

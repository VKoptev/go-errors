# go-errors

[![GoDoc](https://pkg.go.dev/badge/github.com/VKoptev/go-errors)](https://pkg.go.dev/github.com/VKoptev/go-errors)
![golangci-lint](https://github.com/VKoptev/go-errors/workflows/golangci-lint/badge.svg)

Errors wrapper for more convenience :)


## Usage

```go
package something_useful

import (
	"errors"
	"log"

	errs "github.com/VKoptev/go-errors"
)

// SomethingUseful public errors.
var (
	ErrWentCompletelyWrong = errs.WithReason("went completely wrong")
)

// SomethingUseful private errors.
var (
	errWentWrong = errs.WithReason("went wrong")
)

func SomethingUseful() error {
	for _, item := range list {
		if err := partOfSomethingUseful(item); err != nil {
			if errors.Is(err, errWentWrong) {
				log.Printf("there is error: %v", err)
				continue
			}
			return ErrWentCompletelyWrong.WithErr(err)
		}
	}
	return nil
}

func partOfSomethingUseful(data) error {
	if err := doSmthUnimaginableWithData(data); err != nil {
		return ErrWentCompletelyWrong.WithX(data)
	}
	if err := doSmthAcceptable(); err != nil {
		return errWentWrong
	}
	return nil
}

```

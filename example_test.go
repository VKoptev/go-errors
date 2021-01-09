package errors_test

import (
	"fmt"
	"os"

	"github.com/VKoptev/go-errors"
)

// Imagine that it's public errors of your useful_package.
var (
	ErrSomethingWentWrong = errors.WithReason("something went completely wrong")
)

// Also there are private error of your useful_package.
var (
	errSomethingGoingProbablyWrong = errors.WithReason("something is going probably wrong")
)

func DoSomethingUseful() ([]int, error) {
	var result []int

	for i := 0; i < 10; i++ {
		err := doVeryUsefulPart(i)
		if err != nil {
			// Checking errors is same that builtin package.
			// If got errSomethingGoingProbablyWrong (or error wrapping it) we should skip number.
			if errors.Is(err, errSomethingGoingProbablyWrong) {
				continue
			}

			// Otherwise we must return public error to user and... ohgosh! data that produced error. Nice.
			return nil, ErrSomethingWentWrong.WithX(i)
		}

		result = append(result, i)
	}

	return result, nil
}

func doVeryUsefulPart(i int) error {
	if i%2 == 1 { // is it odd number? suspiciously...
		return errSomethingGoingProbablyWrong
	}

	if i == 7 { // amount of Voldemort's horcruxes!
		//nolint:golint,lll,goerr113
		return fmt.Errorf("I'm Voldemort! I can ignore linter and write very long in-placed errors capitalized and ended with punctuation!")
	}

	return nil
}

func Example() {
	r, err := DoSomethingUseful()
	if err != nil {
		var terr *errors.Error

		if errors.As(err, &terr) {
			// If returned error may be matched as *errors.Error
			// let's just use it pretty print included wrapped errors and data.
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Otherwise let's form pretty print.
		fmt.Printf("Got error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Got data: %v", r)
	os.Exit(0)
}

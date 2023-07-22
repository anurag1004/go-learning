package main

import (
	"errors"
	"fmt"
)

/*
However, it's generally recommended to use %w instead of %v when wrapping errors with additional context,
as %w creates a new error that includes the original error and additional context,
while %v only formats the error into a string. This makes it easier to unwrap the error and get the original error message.
*/
func foo() error {
	
	return fmt.Errorf("foo: %w", bar())
}
func bar() error {
	return errors.New("bar")
}
func main() {
	err := foo()
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
	}
}
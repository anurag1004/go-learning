package main

import (
	"errors"
	"fmt"
	"time"
)

type DivideByZero string

func (ces DivideByZero) Error() string { // DevideByZero now implements Error interface
	return fmt.Sprintf("Error: %s,%v\n", string(ces), time.Now())
}
func divide(i, j int) (float32, error) {
	if j == 0 {
		return 0, DivideByZero("Cannot divide by zero!")
	}
	result := (float32(i) / float32(j))
	return result, nil
}
func foo() error {
	return errors.New("Error From foo()")
}
func main() {
	if ans, err := divide(10, 0); err != nil {
		if errordes, ok := err.(DivideByZero); ok {
			fmt.Println("DividebyZeroType::" + errordes.Error())
		} else {
			fmt.Println(errordes)
		}
	} else {
		fmt.Println(ans)
	}
	if err := foo(); err != nil {
		fmt.Println(err)
	}
}

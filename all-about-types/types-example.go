package main

import "fmt"

type Time int // defaults to mili seconds
type Player struct {
	name string
	age  int
}
type Players []Player

func (t Time) getSeconds() float32 {
	return float32(t) / 1000
}
func (t *Time) incrementByOne() {
	*t++
}
func printTimes(ts ...Time) {
	fmt.Println(ts) // [] Time
}
func printNums(num1, num2 int) {
	fmt.Println(num1, num2)
}
func main() {
	t := Time(2000)
	fmt.Println(t)
	fmt.Println(t.getSeconds())
	i := 0
	for i < 1000 {
		//increment 1000ms
		t.incrementByOne()
		i++
	}
	fmt.Println(t.getSeconds())
	printTimes(t, t, t, t, t)
}

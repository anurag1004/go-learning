package main

import "fmt"

type Speaker interface {
	GetSoundQuality() int
}
type Boat struct{}
type Sony struct{}

func (b Boat) GetSoundQuality() int {
	return 72
}
func (s Sony) GetSoundQuality() int {
	return 89
}

func compareSoundQuality(s1, s2 Speaker) int {
	return s1.GetSoundQuality() - s2.GetSoundQuality()
}
func describe(i interface{}) {
	fmt.Printf("%T:%v\n", i, i)
}
func drawLine() {
	fmt.Println("----------------")
}
func showArgs(args ...interface{}) {
	fmt.Println(args...)
}
func main() {
	drawLine()
	/////////////////////

	b := Boat{}
	s := Sony{}
	fmt.Println(compareSoundQuality(b, s))

	//////////////////////

	drawLine()
	var x interface{} // holds a type nil and pointer that points to value nil
	describe(x)
	y := 10 // y holds a type int and a pointer that points to value 10
	describe(y)
	x = y
	describe(x)

	/////////////////////

	drawLine()
	// type assertion
	val, ok := x.(int)
	fmt.Println("INT Type Exists: ", ok)
	fmt.Println(val)

	val2, ok2 := x.(float32) // if a specified type is not present then not declaring ok2 will generate a panic
	fmt.Println("FLAOT32 Type Exists: ", ok2)
	fmt.Println(val2)

	//////////////////////
	drawLine()

	//////////////////////
	// Type switch

	switch v := x.(type) {
	case int:
		fmt.Println("Type is INT", v)
	case float32:
		fmt.Println("Type is FLOAT32", v)
	default:
		fmt.Println("Undefined type!", v)
	}

	drawLine()

	/////////////////////

	arr := []int{1, 2, 3, 4, 5}
	foo := 34.12
	str := "hello"
	showArgs(arr, foo, str)

	drawLine()

	////////////////////
}

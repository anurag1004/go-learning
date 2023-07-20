package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
	"github.com/wayneashleyberry/truecolor/pkg/color"
)

func main() {
	fmt.Println("hi")
	color.Color(186, 218, 85).Println("Hello, World!")
	color.Black().Background(186, 218, 85).Println("Hello, World!")
	color.White().Underline().Print("Hello, World!\n")
	color.White().Dim().Println("Hello, World!")
	color.White().Italic().Println("Hello, World!")
	color.White().Bold().Println("Hello, World!")
	color.Color(255, 165, 00).Printf("Hello, %s!\n", "World")
	fmt.Printf("Hello, %s\n", color.Color(255, 0, 0).Sprint("World", "!"))

	nums := []int{1, 3, 5, 2, 4, 8}
	odds := nums[:3]
	evens := nums[3:]

	nums[1], nums[3] = 9, 6
	s.Show("nums", nums)
	s.Show("odds : nums[:3]", odds)
	s.Show("evens: nums[3:]", evens)

	fmt.Println(util.stringLen("anurag"))
}

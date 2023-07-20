package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

/*
Here Player and count now both belongs to stringer interface.
fmt package while printing relies on this internal String() interface
So basically we r overriding the String() behaviour for Player and count
*/
type Player struct {
	name string
	age  int
}
type count int

func (p Player) String() string {
	return fmt.Sprintf("-----------\nName:%v\nAge:%v\n-----------\n", p.name, p.age)
}

func (c count) String() string {
	return "Count is : " + strconv.Itoa(int(c))
}
func WriteOut(f io.Writer, s fmt.Stringer) {
	_, err := f.Write(bytes.NewBufferString(s.String()).Bytes())
	if err != nil {
		fmt.Println("Error while writing the data!")
		os.Exit(-1)
	}
}
func main() {
	buff := bytes.NewBuffer([]byte{})
	buff.WriteString("Hello World\nHello Universe")
	str := "hii"
	newBuff := bytes.NewBufferString(str)
	newBuff.WriteString("\nblah blah 😁\n")
	fmt.Println(newBuff.String())

	p := Player{
		name: "Anurag",
		age:  23,
	}
	c := count(29)
	fmt.Println(p)
	fmt.Println(c)

	f, _ := os.Create("player.txt")
	WriteOut(f, p)

	cByte := bytes.NewBuffer([]byte{})
	WriteOut(cByte, c)
	fmt.Println(cByte.String())
}

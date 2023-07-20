package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

type deck []string

func newCardDeck() deck {
	// d := make(deck, 0, 30) // creats a slice of type deck with initial length as 0 and capacity as 30
	d := deck{}
	symbols := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	values := []string{"Six", "Ace", "Two", "Jack", "Queen", "Three", "Seven"}
	for _, symbol := range symbols {
		for _, value := range values {
			d = append(d, value+" Of "+symbol)
		}
	}
	return d
}
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
func (d deck) shuffle() {
	for i := range d {
		randIndx := int(rand.Float32() * float32(len(d)))
		// fmt.Println(randIndx, i)
		d[i], d[randIndx] = d[randIndx], d[i]
	}
}

func (d deck) writeToFile(fileName string) (string, error) {
	data := []byte(strings.Join(d, ","))
	err := ioutil.WriteFile(fileName, data, 0666)
	if err != nil {
		return "FAILED", err
	}
	fmt.Println("Data wrote to file: ", fileName)
	return "SUCCESS", nil
}

func getDeckFromFile(fileName string) (deck, error) {
	data_bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return deck(strings.Split(string(data_bs), ",")), nil
}

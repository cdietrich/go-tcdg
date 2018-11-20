package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSuites := []string{"Clubs", "Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	result := make(deck, 0, len(cardSuites)*len(cardValues))
	for _, s := range cardSuites {
		for _, v := range cardValues {
			result = append(result, fmt.Sprint(v, " of ", s))
		}
	}
	return result
}

func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func deal(d deck, n int) (deck, deck) {
	return d[:n], d[n:]
	// result := deck{}
	// for i := 0; i < n; i++ {
	// 	k := rand.Intn(len(d))
	// 	result = append(result, d[k])
	// 	d = append(d[:k], d[k+1:]...)
	// }
	// return d, result
}

func (d deck) shuffe() {
	rand.Seed(time.Now().UnixNano())
	//source := rand.NewSource(time.Now().UnixNano())
	//r := rand.New(source)
	for i := range d {
		newPosition := rand.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) saveToFile2(n string) {
	err := ioutil.WriteFile(n, []byte(d.toString()), 600)
	if err != nil {
		log.Fatal(err)
	}
}

func newDeckFromFile2(n string) deck {
	b, err := ioutil.ReadFile(n)
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(string(b), ",")
	return deck(s)
}

func (d deck) saveToFile(n string) {
	file, err := os.Create(n)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(d)
	if err != nil {
		log.Fatal(err)
	}
}

func newDeckFromFile(n string) deck {
	d := deck{}
	file, err := os.Open(n)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&d)
	if err != nil {
		panic(err)
	}
	return d
}

func (d deck) toString() string {
	s := strings.Join([]string(d), ",")
	return s
}

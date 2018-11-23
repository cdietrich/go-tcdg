package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Printf("%q\n", m)
	fmt.Printf("%+v\n", m)
	fmt.Println(m["red"])
	m["black"] = "#000000"
	fmt.Println(m)
	_, ok := m["white"]
	if !ok {
		fmt.Println("No White")
	}
	delete(m, "red")
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("The hex code for %v is %v\n", k, v)
	}

	m2 := make(map[string]string)
	m2["x"] = "X"
	m2["y"] = "Y"
	fmt.Println(m2)
}

package main

import (
	"fmt"

	"../../pkg/fileio"
	"../../pkg/parsing"
)

func main() {
	x := `{"name": "Domagoj", "age" :19, "shoe_size": 44.5, "male": true}`
	fmt.Println(parsing.Parsel(x))

	// example for how to test quickly
	// NOTE: data does not have to be string, change it how you like
	fileio.SaveFile("user", "file", "data")
}

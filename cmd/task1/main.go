package main

import (
	"text/template/parse"

	"../../pkg/fileio"
)

func main() {

	parse.Parse(`{"name": "Domagoj", "age" :19, "shoe_size": 44.5, "male": true}`)

	// example for how to test quickly
	// NOTE: data does not have to be string, change it how you like
	fileio.SaveFile("user", "file", "data")
}

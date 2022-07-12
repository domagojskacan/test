package main

import "../../pkg/fileio"

func main() {

	// example for how to test quickly
	// NOTE: data does not have to be string, change it how you like
	fileio.SaveFile("user", "file", "data")
}

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `uri:"username"`
	Filename string `uri:"filename"`
	Data     string `json:"data"`
}

func getValue(c *gin.Context) {
	var input user
	if err := c.ShouldBindUri(&input); err == nil {
		fmt.Printf("%+v", input)
	} else {
		fmt.Println("error")
	}

	if err1 := c.ShouldBindJSON(&input); err1 == nil {
		fmt.Printf("%+v", input)
	} else {
		fmt.Println("error")
	}
	fmt.Println(input.Username, input.Filename, input.Data)
}

func main() {
	router := gin.Default()
	router.Run("localhost:9091")
	//x := `{"name": "Domagoj", "age" :19, "shoe_size": 44.5, "male": true}`
	//fmt.Println(parsing.Parsel(x))

	router.PUT("/user/:username/:filename", getValue)

	// example for how to test quickly
	// NOTE: data does not have to be string, change it how you like
	//fileio.SaveFile("user", "file", "data")

}

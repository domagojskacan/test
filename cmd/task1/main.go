package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../pkg/fileio"
	"../../pkg/parsing"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//x := `{"name": "Domagoj", "age" :19, "shoe_size": 44.5, "male": true}`
	//fmt.Println(parsing.Parsel(x))
	router.GET("/user/:username/:filename", func(c *gin.Context) {
		name := c.Param("username")
		action := c.Param("filename")
		contentType := c.Request.Header[http.CanonicalHeaderKey(("Content-Type"))]
		message := name + " is " + action
		fmt.Println(contentType)
		c.String(http.StatusOK, message)
		switch contentType[0] {
		case "application/json":
			fmt.Println("case je appjson")
			data, _ := ioutil.ReadAll(c.Request.Body)
			s := string([]byte(data))
			v := parsing.ParseJson(s)
			fileio.SaveCsv(name, action, v)
			fileio.SaveJson(name, action, v)
			fileio.SaveYaml(name, action, v)

		case "application/csv":
			fmt.Println("case je appcsv")
			csvReader := csv.NewReader(c.Request.Body)
			data, err := csvReader.ReadAll()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%T %v", data, data)
			v := parsing.ParseCsv(data)
			fileio.SaveCsv(name, action, v)
			fileio.SaveJson(name, action, v)
			fileio.SaveYaml(name, action, v)

		case "application/yaml":
			fmt.Println("case je appyaml")
			data, _ := ioutil.ReadAll(c.Request.Body)
			v := parsing.ParseYaml(data)

			fileio.SaveJson(name, action, v)
			fileio.SaveYaml(name, action, v)
			fileio.SaveCsv(name, action, v)
		}
	})

	//router.PUT("/user/:username/:filename", getValue)

	// example for how to test quickly
	// NOTE: data does not have to be string, change it how you like
	//fileio.SaveFile("user", "file", "data")
	router.Run("localhost:9091")
}

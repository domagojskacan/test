package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../pkg/delete"

	"../../pkg/fileio"
	"../../pkg/parsing"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//x := `{"name": "Domagoj", "age" :19, "shoe_size": 44.5, "male": true}`
	//fmt.Println(parsing.Parsel(x))
	router.POST("/user/:username/:filename", func(c *gin.Context) {
		name := c.Param("username")
		action := c.Param("filename")
		contentType := c.Request.Header[http.CanonicalHeaderKey(("Content-Type"))]

		switch contentType[0] {
		case "application/json":
			fmt.Println("case je appjson")
			data, _ := ioutil.ReadAll(c.Request.Body)
			s := string([]byte(data))
			v := parsing.ParseJson(s)
			err := fileio.SaveCsv(name, action, v)
			if err != nil {
				c.String(http.StatusNotAcceptable, "error")
			}

			fileio.SaveJson(name, action, v)
			fileio.SaveYaml(name, action, v)

		case "application/csv":
			fmt.Println("case je appcsv")
			csvReader := csv.NewReader(c.Request.Body)
			data, err := csvReader.ReadAll()
			if err != nil {
				c.String(http.StatusNotAcceptable, "error")
			}
			fmt.Printf("%T %v", data, data)
			v := parsing.ParseCsv(data)

			err1 := fileio.SaveCsv(name, action, v)
			if err1 != nil {
				c.String(http.StatusNotAcceptable, "error")
			}
			fileio.SaveJson(name, action, v)
			fileio.SaveYaml(name, action, v)

		case "application/yaml":
			fmt.Println("case je appyaml")
			data, _ := ioutil.ReadAll(c.Request.Body)
			v := parsing.ParseYaml(data)

			err := fileio.SaveJson(name, action, v)
			if err != nil {
				c.String(http.StatusNotAcceptable, "error")
			}
			fileio.SaveYaml(name, action, v)
			fileio.SaveCsv(name, action, v)
		}
	})

	router.DELETE("/user/:username/:filename", func(c *gin.Context) {
		name := c.Param("username")
		action := c.Param("filename")
		contentType := c.Request.Header[http.CanonicalHeaderKey(("Content-Type"))]

		switch contentType[0] {
		case "application/json":

			err := delete.DelJson(name, action)
			if err != nil {
				c.String(http.StatusNotFound, "error")
			} else {
				c.String(http.StatusOK, "file obrisan")
			}
		case "application/csv":

			err := delete.DelCsv(name, action)
			if err != nil {
				c.String(http.StatusNotFound, "error")
			} else {
				c.String(http.StatusOK, "file obrisan")
			}
		case "application/yaml":

			err := delete.DelYaml(name, action)
			if err != nil {
				c.String(http.StatusNotFound, "error")
			} else {
				c.String(http.StatusOK, "file obrisan")
			}

		}
	})

	router.GET("/user/:username/:filename", func(c *gin.Context) {
		name := c.Param("username")
		action := c.Param("filename")
		contentType := c.Request.Header[http.CanonicalHeaderKey(("Content-Type"))]

		switch contentType[0] {
		case "application/json":
			err, data := fileio.ReadJson(name, action)
			if err != nil {
				c.String(http.StatusNotFound, "error")
			} else {
				c.String(http.StatusOK, data)
			}
		case "application/csv":
			err, data := fileio.ReadYaml(name, action)
			if err != nil {
				c.String(http.StatusNotFound, "error")
			} else {
				c.String(http.StatusOK, data)
			}
		case "application/yaml":
			err, data := fileio.ReadYaml(name, action)
			if err != nil {
				c.String(http.StatusNotFound, "error")
			} else {
				c.String(http.StatusOK, data)
			}
		}

	})
	router.Run("localhost:9091")
}

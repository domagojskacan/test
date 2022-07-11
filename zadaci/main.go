package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//type i interface
type data struct {
	Ime     string `json:"Ime"`
	Prezime string `json:"Prezime"`
	Godina  int    `json:"Godina"`
}

var d data

var pod []data

//var dataj struct {
//	Ime     string  `json:"Ime"`
//	Prezime string `json:"Prezime"`
//	Godina  int `json:"Godina"`
//}

var x data

//case json
func getValue(context *gin.Context) {
	context.JSON(http.StatusOK, d)
	//x := d
	krv, _ := json.Marshal(d)
	_ = ioutil.WriteFile("file.json", krv, 0644)

	//if err := convertJSONToCSV("data.json", "data.csv"); err != nil {
	//	log.Fatal(err)}
}

//case csv

//case yaml

//convert json to csv
func JSONtoCSV(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	var ranking []data
	if err := json.NewDecoder(sourceFile).Decode(&ranking); err != nil {
		return err
	}

	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"Ime", "Prezime", "Godina"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range ranking {
		var csvRow []string
		csvRow = append(csvRow, r.Ime, r.Prezime, fmt.Sprint(r.Godina))
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

//convert

//convert

//convert

//convert

//convert

//convert

//
//

func main() {
	router := gin.Default()
	router.GET("/")
	router.POST("/d", getValue)
	router.Run("localhost:9091")

}

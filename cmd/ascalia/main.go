package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Metric_id  int     `json:"Metric_id"`
	Vrijednost float64 `json:"Vrijednost"`
	Ts         int64   `json:"Ts"`
}

var pod = data{}
var podaci = []data{}
var x float64

func diff(a, b float64) float64 {
	if a < b {
		return b - a
	}
	return a - b
}

func getValue(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, podaci)
}

func addpodaci(context *gin.Context) {
	var newValue data
	if err := context.BindJSON(&newValue); err != nil {
		return
	}
	for _, element := range podaci {
		if newValue.Metric_id == element.Metric_id {
			x = diff(newValue.Vrijednost, pod.Vrijednost)
		}
		podaci = append(podaci, newValue)
		context.IndentedJSON(http.StatusCreated, newValue)

	}
}

func main() {
	podaci = append(podaci, pod)
	router := gin.Default()
	//router.GET("/podaci", getValue)
	router.POST("/podaci", addpodaci)
	router.Run("localhost:9091")
	fmt.Println(x)
}

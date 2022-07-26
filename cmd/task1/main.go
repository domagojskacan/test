package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"os"
	"os/signal"
	"syscall"
	"time"

	"../../pkg/delete"
	"../../pkg/fileio"
	"../../pkg/parsing"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.Run("localhost:9091")

	srv := &http.Server{
		Addr:    ":9091",
		Handler: router,
	}

	router.POST("/user/:username/:filename", func(c *gin.Context) {
		name := c.Param("username")
		action := c.Param("filename")

		contentType := c.Request.Header[http.CanonicalHeaderKey(("Content-Type"))]

		switch contentType[0] {
		case "application/json":
			fmt.Println("case je appjson")
			data, _ := ioutil.ReadAll(c.Request.Body)
			str := string([]byte(data))
			v := parsing.ParseJson(str)
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
				return
			}
			c.String(http.StatusOK, "file obrisan")

		case "application/csv":

			err := delete.DelCsv(name, action)
			if err != nil {
				c.String(http.StatusNotFound, "error")
				return
			}
			c.String(http.StatusOK, "file obrisan")

		case "application/yaml":

			err := delete.DelYaml(name, action)
			if err != nil {
				c.String(http.StatusNotFound, "error")
				return
			}
			c.String(http.StatusOK, "file obrisan")

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
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 3 seconds.")
	}
	log.Println("Server exiting")

}

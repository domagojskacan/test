package fileio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func SaveJson(username string, filename string, data map[string]interface{}) {
	path := filepath.Join("../../test", username)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Println(err)
		}
	}
	pod, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error:", err)
	}

	path = filepath.Join("../../test", username, filename+".json")
	err1 := ioutil.WriteFile(path, pod, 0644)
	if err != nil {
		log.Fatal(err1)
	}
}

func SaveCsv(username string, filename string, data map[string]interface{}) {

	path := filepath.Join("../../test", username)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Println(err)

		}
	}

	var keys string
	var values string

	for key, value := range data {

		fmt.Println(key)
		fmt.Println(value)
		if len(keys) == 0 {
			keys = key
			values = fmt.Sprintf("%v", value)
			continue
		}
		keys = fmt.Sprintf("%s, %s", keys, key)
		values = fmt.Sprintf("%s, %v", values, value)
	}

	finalString := fmt.Sprintf("%s \n %s", keys, values)
	fmt.Println(finalString)
	path = filepath.Join("../../test", username, filename+".csv")

	bs := []byte(finalString)

	err := ioutil.WriteFile(path, bs, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveYaml(username string, filename string, data map[string]interface{}) {
	path := filepath.Join("../../test", username)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Println(err)

		}
	}

	pod, err := yaml.Marshal(data)
	if err != nil {
		fmt.Println("error:", err)
	}
	path = filepath.Join("../../test", username, filename+".yaml")
	err1 := ioutil.WriteFile(path, pod, 0644)
	if err != nil {
		log.Fatal(err1)
	}
}

func ReadYaml(username string, filename string) (string, error) {
	return "", nil
}

func ReadJson(username string, filename string) (string, error) {
	return "", nil
}

func ReadCsv(username string, filename string) (string, error) {
	return "", nil
}

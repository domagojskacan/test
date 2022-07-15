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

func SaveJson(username string, filename string, data map[string]interface{}) error {
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

	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("file alredy exist")
	} else {
		err := ioutil.WriteFile(path, pod, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func SaveCsv(username string, filename string, data map[string]interface{}) error {

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

		if len(keys) == 0 {
			keys = key
			values = fmt.Sprintf("%v", value)
			continue
		}
		keys = fmt.Sprintf("%s, %s", keys, key)
		values = fmt.Sprintf("%s, %v", values, value)
	}

	finalString := fmt.Sprintf("%s \n %s", keys, values)
	path = filepath.Join("../../test", username, filename+".csv")

	bs := []byte(finalString)

	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("file alredy exist")
	} else {
		err := ioutil.WriteFile(path, bs, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func SaveYaml(username string, filename string, data map[string]interface{}) error {
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
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("file alredy exist")
	} else {
		err := ioutil.WriteFile(path, pod, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func ReadYaml(username string, filename string) (error, string) {
	path := filepath.Join("../../test", username, filename+".csv")
	if _, err := os.Stat(path); err == nil {
		data, _ := ioutil.ReadFile(path)
		s := string([]byte(data))
		return nil, s
	} else {
		ret := "file does not exist"
		return err, ret
	}
}

func ReadJson(username string, filename string) (error, string) {
	path := filepath.Join("../../test", username, filename+".json")
	if _, err := os.Stat(path); err == nil {
		data, _ := ioutil.ReadFile(path)
		s := string([]byte(data))
		return nil, s
	} else {
		ret := "file does not exist"
		return err, ret
	}
}

func ReadCsv(username string, filename string) (error, string) {
	path := filepath.Join("../../test", username, filename+".csv")
	if _, err := os.Stat(path); err == nil {
		data, _ := ioutil.ReadFile(path)
		s := string([]byte(data))
		return nil, s
	} else {
		ret := "file does not exist"
		return err, ret
	}
}

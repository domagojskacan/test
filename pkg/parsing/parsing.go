package parsing

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

func ParseJson(jsonStr string) map[string]interface{} {

	m := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func ParseCsv(csv [][]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for index, key := range csv[0] {
		ret[key] = csv[1][index]

	}
	return ret
}

func ParseYaml(data []byte) map[string]interface{} {
	ret := make(map[string]interface{})

	err := yaml.Unmarshal(data, &ret)
	if err != nil {
		fmt.Println(err)

	}
	return ret
}

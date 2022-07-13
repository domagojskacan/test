package parse

import (
	"encoding/json"
	"fmt"
)

func parse(x string) map[string]interface{} {

	type mapa map[string]interface{}

	var m mapa

	err := json.Unmarshal([]byte(x), &m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

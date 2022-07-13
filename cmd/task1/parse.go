package parse

import (
	"encoding/json"
	"fmt"
)

func parse(x string) (m map[string]interface{}) {

	type mapa map[string]interface{}

	err := json.Unmarshal([]byte(x), &m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

package parsing

import (
	"encoding/json"
	"fmt"
)

func Parsel(jsonStr string) map[string]interface{} {

	type mapa map[string]interface{}

	var m mapa

	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

package parsing

import (
	"fmt"
	"testing"
)

func TestParseJson(t *testing.T) {
	mapa := map[string]interface{}{"name": "Domagoj", "age": 19.0, "shoe_size": 44.5, "male": true}
	r := ParseJson(`{"name": "Domagoj", "age" :19, "shoe_size": 44.5, "male": true}`)
	res := equal(mapa, r)
	if res == true {
		fmt.Println("Parse json radi kako treba")
	} else {
		t.Error("expected", mapa, "got", r)
	}
}

func TestParseYaml(t *testing.T) {
	mapa := map[string]interface{}{"name": "Domagoj", "age": 19, "shoe_size": 44.5, "male": true}
	b := []byte("name: Domagoj\nage: 19\nshoe_size: 44.5\nmale: true")
	r := ParseYaml(b)
	res := equal(mapa, r)
	if res == true {
		fmt.Println("Parse yaml radi kako treba")
	} else {
		t.Error("expected", mapa, "got", r)
	}
}

func TestParseCsv(t *testing.T) {
	mapa := map[string]interface{}{"name": "Domagoj", "age": "19", "shoe_size": "44.5", "male": "true"}

	data := [][]string{[]string{"name", "age", "shoe_size", "male"}, []string{"Domagoj", "19", "44.5", "true"}}
	r := ParseCsv(data)
	res := equal(mapa, r)
	if res == true {
		fmt.Println("Parse csv radi kako treba")
	} else {
		t.Error("expected", mapa, "got", r)
	}
}

func equal(mapa1, mapa2 map[string]interface{}) bool {
	if len(mapa1) != len(mapa2) {
		return false
	}
	for k, mapa1v := range mapa1 {
		if mapa2v, ok := mapa2[k]; !ok || mapa2v != mapa1v {
			return false
		}
	}
	return true
}

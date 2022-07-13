package testing

import "testing"

func TestParse (T *Testing.T) {
	r := parse (`{"name": "Domagoj", "age" :19, "shoe_size": 44.5, "male": true}`)
	if r != map[name:Domagoj age:19 shoe_size:44.5 male:true] {t.Error ("expected", [name:Domagoj age:19 shoe_size:44.5 male:true], "got" r)
}
}
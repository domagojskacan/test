package testing

import "testing"

func TestParse (T *Testing.T) {
	r := parse(`{"name": "Domagoj", "age" :19, "shoe_size": 44.5, "male": true}`)
	if r != map[age:19 male:true name:Domagoj shoe_size:44.5] 
	{t.Error ("expected", [age:19 male:true name:Domagoj shoe_size:44.5], "got" r)
}
}
package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func main() {
	jsonParsed, err := gabs.ParseJSON([]byte(`{
	"outter":{
		"inner":{
			"value1":10,
			"value2":22
		},
		"alsoInner":{
			"value1":20,
			"array1":[
				30, 40
			]
		}
	}
}`))
	if err != nil {
		panic(err)
	}

	var value float64
	var ok bool

	value, ok = jsonParsed.Path("outter.inner.value1").Data().(float64)
	// value == 10.0, ok == true

	value, ok = jsonParsed.Search("outter", "inner", "value1").Data().(float64)
	// value == 10.0, ok == true

	value, ok = jsonParsed.Search("outter", "alsoInner", "array1", "1").Data().(float64)
	// value == 40.0, ok == true

	gObj, err := jsonParsed.JSONPointer("/outter/alsoInner/array1/1")
	if err != nil {
		panic(err)
	}
	value, ok = gObj.Data().(float64)
	// value == 40.0, ok == true

	value, ok = jsonParsed.Path("does.not.exist").Data().(float64)
	// value == 0.0, ok == false

	exists := jsonParsed.Exists("outter", "inner", "value1")
	// exists == true

	exists = jsonParsed.ExistsP("does.not.exist")
	// exists == false

	fmt.Println(value, ok, exists)
}

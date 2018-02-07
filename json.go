package main

import (
	"encoding/json"
	_ "reflect"
)

func encode(input interface{}) string {
	output, err := json.Marshal(input)
	handleError(err)
	return string(output)
}

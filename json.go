package main

import (
	"encoding/json"
)

func encode(input interface{}) string {
	output, err := json.Marshal(input)
	return output
}
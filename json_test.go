package main

import "testing"

func TestEncode(t *testing.T) {
	// a struct
	input := &Restaurant{Id: 1, Name: "foo", LocationId: 1}
	output := encode(input)
	expectedOutput := "{\"Id\":1,\"name\":\"foo\",\"locationId\":\"1\"}"
	if output != expectedOutput {
		t.Errorf("Mismatch: wanted %s got %s", expectedOutput, output)
	}
}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	validJSON := `{}`
	invalidJSON := `"this is valid wtf"`

	fmt.Println(json.Valid([]byte(validJSON)))   // true
	fmt.Println(json.Valid([]byte(invalidJSON))) // false
}

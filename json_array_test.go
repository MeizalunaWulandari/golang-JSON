package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Meizaluna",
		LastName:  "Wulandari",
		Hobbies: []string{
			"Gaming",
			"Reading",
			"Coding",
			"Sleeping",
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

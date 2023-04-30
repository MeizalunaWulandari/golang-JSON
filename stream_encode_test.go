package golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encode := json.NewEncoder(writer)

	customer := &Customer{
		FirstName: "Meizaluna",
		LastName:  "Wulandari",
	}

	encode.Encode(customer)
}

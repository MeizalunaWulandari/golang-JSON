package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PortalCode string
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Married   bool
	Hobbies   []string
	Address   []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Meizaluna",
		LastName:  "Wulandari",
		Age:       16,
		Married:   false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

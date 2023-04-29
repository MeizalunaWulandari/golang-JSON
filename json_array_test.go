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

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `
		{
			"FirstName":"Meizaluna",
			"LastName":"Wulandari",
			"Age":16,
			"Married":false,
			"Hobbies":[
				"Gaming",
				"Reading",
				"Coding",
				"Sleeping"
			]
		}
	`
	jsonBytes := []byte(jsonString)
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Meizaluna",
		Address: []Address{
			{
				Street:     "Jl. Pramuka 17",
				Country:    "Indonesia",
				PortalCode: "76148",
			},
			{
				Street:     "Jl. Sultan Adam",
				Country:    "Indonesia",
				PortalCode: "88102",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `
		{
			"FirstName":"Meizaluna",
			"LastName":"Wulandari",
			"Age":16,
			"Married":false,
			"Hobbies":[
				"Sleeping",
				"Reading",
				"Coding"
			],
			"Address":[
				{
					"Street":"Jl. Pramuka 17",
					"Country":"Indonesia",
					"PortalCode":"76148"},
				{
					"Street":"Jl. Sultan Adam",
					"Country":"Indonesia",
					"PortalCode":"88102"
				}
			]
		}

	`
	jsonBytes := []byte(jsonString)
	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Address)

}

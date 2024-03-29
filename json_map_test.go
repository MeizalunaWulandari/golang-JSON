package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMap(t *testing.T) {
	jsonString := `{"id":"PDT003","name":"Titan Gamen Pink","price":230000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestJSONMapDecode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "PDT003",
		"name":  "Titan Gamen Pink",
		"price": 230000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

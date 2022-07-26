package main

import (
	"encoding/json"

	"fmt"
)

func main() {
	obj := make(map[string]map[string]string)
	id := []string{"1234", "5678"}
	customer := []string{"1234", "5678"}

	for i := 0; i < len(id); i++ {
		for j := 0; j < len(customer); j++ {

			obj[id[i]] = map[string]string{}
			obj[id[i]]["id"] = id[i]
			obj[id[i]]["customer"] = customer[i]
		}

	}

	for k, v := range obj {

		fmt.Println(k, v)
	}

	jsonStr, _ := json.MarshalIndent(obj, "", " ")
	fmt.Println(string(jsonStr))
}

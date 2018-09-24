package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	var jsonData map[string]interface{}
	any := url.Values{"method": {"get"}, "id": {"1"}}
	jsons, _ := json.Marshal(any)
	json.Unmarshal(jsons, &jsonData)
	fmt.Println(jsons)
	fmt.Println(jsonData)
}

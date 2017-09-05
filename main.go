package main

import (
	"fmt"
	"encoding/json"

	"github.com/dmitryk-dk/deploybotApi/structs"
	"github.com/dmitryk-dk/deploybotApi/client"
)

func main() {
	client := client.Client{}
	var users structs.Users
	token, endpoint := client.ReadConfig()
	params := map[string]string{
		"limit":"20",
	}
	data, _ := client.GetData("/users", *endpoint, *token, params)
	json.Unmarshal(data, &users)


	fmt.Printf("%v \n", users)
}

package main

import (
	"encoding/json"
	"fmt"

	"github.com/dmitryk-dk/deploybotApi/client"
	"github.com/dmitryk-dk/deploybotApi/structs"
	"github.com/dmitryk-dk/deploybotApi/trigger"
)

func main() {
	client := client.Client{}
	var users structs.Environments
	trigger := trigger.Trigger{
		EnvironmentId:       61302,
		UserId:              38924,
		DeployedVersion:     "47cdbd47887549d19d5ffab24c91fae7b6ffcb56",
		DeployFromScratch:   false,
		TriggerNotification: false,
		Comment:             "Merge master",
	}
	token, endpoint := client.ReadConfig()
	params := map[string]string{
		"limit":         "20",
		"repository_id": "1810",
	}
	resp, err := trigger.TriggerDeployment("/deployments/61302", *endpoint, *token)
	fmt.Printf("resp %v \n", resp)
	fmt.Printf("error %v \n", err)
	data, _ := client.GetData("/environments", *endpoint, *token, params)
	json.Unmarshal(data, &users)

	fmt.Printf("%v \n", users)
}

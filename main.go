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
	var env structs.Environments
	trigger := trigger.Trigger{
		EnvironmentId:       2,
		UserId:              4,
		DeployedVersion:     "",
		DeployFromScratch:   false,
		TriggerNotification: false,
		Comment:             "Merge master",
	}
	token, endpoint := client.ReadConfig()
	params := map[string]string{
		"limit": "20",
	}
	resp, _ := trigger.TriggerDeployment("/deployments", *endpoint, *token)
	data, _ := client.GetData("/environments", *endpoint, *token, params)
	json.Unmarshal(data, &env)
	fmt.Printf("resp %v \n", resp)
	fmt.Printf("%v \n", env)
}

package trigger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Trigger struct {
	EnvironmentId       int    `json:"environment_id"`
	UserId              int    `json:"user_id"`
	DeployedVersion     string `json:"deployed_version"`
	DeployFromScratch   bool   `json:"deploy_from_scratch"`
	TriggerNotification bool   `json:"trigger_notification"`
	Comment             string `json:"comment"`
}

type TriggerResponse struct {
	Id                  int    `json:"id"`
	RepositoryId        int    `json:"repository_id"`
	EnvironmentId       int    `json:"environment_id"`
	UserId              int    `json:"user_id"`
	DeployedVersion     string `json:"deployed_version"`
	DeployFromScratch   bool   `json:"deploy_from_scratch"`
	TriggerNotification bool   `json:"trigger_notification"`
	IsAutomatic         bool   `json:"is_automatic"`
	Comment             string `json:"comment"`
	AuthorName          string `json:"author_name"`
	State               string `json:"state"`
	Retries             int    `json:"retries"`
	CreatedAt           string `json:"created_at"`
	DeployedAt          string `json:"deployed_at"`
}

func (t *Trigger) TriggerDeployment(action, endpoint, token string) (*TriggerResponse, error) {
	var triggerResponse TriggerResponse
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	route, err := url.Parse(endpoint + action)
	if err != nil {
		return nil, err
	}

	query := url.Values{}

	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	rawData := bytes.NewReader(data)
	route.RawQuery = query.Encode()
	req, err := http.NewRequest("POST", route.String(), rawData)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Api-Token", token)
	if err != nil {
		return nil, err
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &triggerResponse)

	return &triggerResponse, nil
}

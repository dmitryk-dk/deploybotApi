package client

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/url"
	"net/http"
)

type Client struct {
	Endpoint string `json:"endpoint"`
	Token    string `json:"token"`
}

func (c Client) ReadConfig() (*string, *string) {
	flagName := c.ReadFlags()
	file, err := ioutil.ReadFile(*flagName)
	err = json.Unmarshal(file, &c)
	if err != nil {
		return nil, nil
	}
	return &c.Token, &c.Endpoint
}

func (c Client) ReadFlags() (configFile *string) {
	configFile = flag.String("config", "config.json", "Path to config file")
	flag.Parse()
	return
}

func (c Client) GetData(action, endpoint, token string, params map[string]string) ([]byte, error) {
	route, err := url.Parse(endpoint + action)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Add("token", token)

	for key, value := range params {
		query.Add(key, value)
	}

	route.RawQuery = query.Encode()
	response, err := http.Get(route.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body, nil
}

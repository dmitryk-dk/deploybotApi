package client

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	Endpoint string `json:"endpoint"`
	Token    string `json:"token"`
}

func (c *Client) ReadConfig() (*string, *string) {
	flagName := c.ReadFlags()
	file, err := ioutil.ReadFile(*flagName)
	err = json.Unmarshal(file, &c)
	if err != nil {
		return nil, nil
	}
	return &c.Token, &c.Endpoint
}

func (c *Client) ReadFlags() (configFile *string) {
	configFile = flag.String("config", "config.json", "Path to config file")
	flag.Parse()
	return
}

func (c *Client) GetData(action, endpoint, token string, params map[string]string) ([]byte, error) {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

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
	req, err := http.NewRequest("GET", route.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

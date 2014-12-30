package etrade

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Common functions

func unmarshalResponse(resp *http.Response, out interface{}) error {
	data, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(data, out)

	return err
}

func (c *OauthClient) Get(url string, opts map[string]string) (resp *http.Response, err error) {
	log.Println("Getting ", url)
	return c.Consumer.Get(url, opts, &c.Config.AccessToken)
}

func (c *OauthClient) GetUnmarshal(url string, opts map[string]string, v interface{}) error {
	resp, err := c.Get(url, opts)
	if err != nil {
		resp.Body.Close()
		return err
	}
	defer resp.Body.Close()

	err = unmarshalResponse(resp, v)
	if err != nil {
		return err
	}

	return nil
}

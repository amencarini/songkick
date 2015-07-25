package songkick

import (
	"net/http"
	"net/url"
)

var apikey string

const (
	baseURL = "http://api.songkick.com/api/3.0/"
)

type Client struct {
	apikey string
	http   *http.Client
}

func NewClient(key string) Client {
	return Client{
		apikey: key,
		http:   &http.Client{},
	}
}

func (c *Client) clientParams(values map[string]string) string {
	values["apikey"] = c.apikey
	q := url.Values{}
	for k, v := range values {
		q.Set(k, v)
	}
	return q.Encode()
}

func (c *Client) get(resource string, params map[string]string) (*http.Response, error) {
	url := baseURL + resource + "?" + c.clientParams(params)
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

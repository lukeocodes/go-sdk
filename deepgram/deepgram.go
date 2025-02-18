// Deprecated: This is a deprecated version of the SDK and
// you should use https://pkg.go.dev/github.com/deepgram-devs/deepgram-go-sdk
// instead.

package deepgram

import (
	"encoding/json"
	"net/http"
)

var sdkVersion string = "0.7.1"
var dgAgent string = "deepgram-go-sdk/v" + sdkVersion

type Client struct {
	ApiKey string
	Host   string
	Path   string
}

func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
		Host:   "api.deepgram.com",
		Path:   "/v1/projects",
	}
}

func (c *Client) WithHost(host string) *Client {
	c.Host = host
	return c
}

func (c *Client) WithPath(path string) *Client {
	c.Path = path
	return c
}

func GetJson(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

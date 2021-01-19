package clockify

import (
	"net/http"
	"net/url"
	"time"
)

type Clockify struct {
	apiKey     string
	baseUrl    url.URL
	httpClient http.Client
}

func (c Clockify) getUrl(path string) *url.URL {
	u := c.baseUrl
	u.Path += path
	return &u
}

func (c *Clockify) addAuthentication(req *http.Request) {
	req.Header.Add("X-Api-Key", c.apiKey)
}

func (c *Clockify) Name() string {
	return "Clockify"
}

func New(apiKey string) *Clockify {
	return &Clockify{
		apiKey: apiKey,
		baseUrl: url.URL{
			Scheme: "https",
			Host:   "api.clockify.me",
			Path:   "/api/v1",
		},
		httpClient: http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

package toggl

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
)

type Toggl struct {
	apiKey     string
	baseUrl    url.URL
	httpClient http.Client

	workspaceId string
}

func (t Toggl) getUrl(path string) *url.URL {
	u := t.baseUrl
	u.Path += path
	return &u
}

func (t *Toggl) addAuthentication(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:api_token", t.apiKey)))))
}

func (t *Toggl) Name() string {
	return "Toggl"
}

func (t *Toggl) SetWorkspaceId(wsId string) {
	t.workspaceId = wsId
}

func New(apiKey string) *Toggl {
	return &Toggl{
		apiKey: apiKey,
		baseUrl: url.URL{
			Scheme: "https",
			Host:   "api.track.toggl.com",
			Path:   "/api/v8",
		},
		httpClient: http.Client{},
	}
}

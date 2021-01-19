package toggl

import (
	"errors"
	"net/http"
)

func (t *Toggl) Test() error {
	u := t.getUrl("/me")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	t.addAuthentication(req)
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return errors.New("invalid return code from Clockify API")
	}
	return nil
}

package toggl

import (
	"fmt"
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
		return fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return nil
}

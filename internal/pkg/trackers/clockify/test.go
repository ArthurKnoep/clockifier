package clockify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Clockify) Test() error {
	u := c.getUrl("/user")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	c.addAuthentication(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return nil
}

func (c *Clockify) GetUserId() (string, error) {
	u := c.getUrl("/user")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return "", err
	}
	c.addAuthentication(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	var usr user
	if err := json.NewDecoder(resp.Body).Decode(&usr); err != nil {
		return "", err
	}
	return usr.Id, nil
}

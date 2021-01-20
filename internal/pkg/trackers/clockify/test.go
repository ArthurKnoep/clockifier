package clockify

import (
	"encoding/json"
	"errors"
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
		return errors.New("invalid return code from Clockify API")
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
		return "", errors.New("invalid return code from Clockify API")
	}
	var usr user
	if err := json.NewDecoder(resp.Body).Decode(&usr); err != nil {
		return "", err
	}
	return usr.Id, nil
}

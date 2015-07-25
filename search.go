package songkick

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Search struct {
	Resultspage struct {
		Status  string `json:"status"`
		Results struct {
			Artists []Artist `json:"artist"`
			Events  []Event  `json:"event"`
		} `json:"results"`
		Perpage      int `json:"perPage"`
		Page         int `json:"page"`
		Totalentries int `json:"totalEntries"`
	} `json:"resultsPage"`
}

func (c *Client) doSearch(url string, params map[string]string) (Search, error) {
	var result Search

	resp, err := c.get(url, params)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, errors.New("Failed connection, error")
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

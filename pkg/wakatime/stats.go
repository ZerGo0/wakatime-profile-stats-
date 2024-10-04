package wakatime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetStats(period string) (*WakaStats, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", fmt.Sprintf("%s/users/current/stats/"+period, c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.doFunc(c, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var stats WakaStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, err
	}

	return &stats, nil
}

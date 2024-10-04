package wakatime

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetStats(period string) (*WakaStats, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/current/stats/"+period, c.baseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.doFunc(c, req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var stats WakaStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &stats, nil
}

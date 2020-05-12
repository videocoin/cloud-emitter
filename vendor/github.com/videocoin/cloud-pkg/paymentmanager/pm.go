package paymentmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	host string
	cli  *http.Client
}

func NewClient(host string, opts ...Option) *Client {
	c := &Client{
		host: host,
		cli:  &http.Client{},
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

func (c *Client) GetReward(address string) (*Reward, error) {
	url := fmt.Sprintf("https://%s/api/v1/receivers/%s/rewards", c.host, address)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error %d", resp.StatusCode)
	}

	reward := new(Reward)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, reward)
	if err != nil {
		return nil, err
	}

	return reward, nil
}

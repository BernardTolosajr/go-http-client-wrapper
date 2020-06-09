package client

import (
	"context"
	"fmt"
	"net/url"
)

type DeleteMethod method

func (c *DeleteMethod) Call(ctx context.Context, path string, parameters map[string]string) (interface{}, error) {
	baseUrl, err := url.Parse(fmt.Sprintf("%s%s", c.client.BaseURL, path))

	if err != nil {
		return nil, err
	}

	if parameters != nil {
		params := url.Values{}
		for k, v := range parameters {
			params.Add(k, v)
		}
		baseUrl.RawQuery = params.Encode()
	}

	req, err := c.client.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

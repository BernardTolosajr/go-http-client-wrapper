package client

import (
	"context"
	"fmt"
)

type DeleteMethod method

func (c *DeleteMethod) Call(ctx context.Context, path string, parameter map[string]string) (interface{}, error) {
	url, err := buildQueryParameter(fmt.Sprintf("%s%s", c.BaseURL, path), parameter)

	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

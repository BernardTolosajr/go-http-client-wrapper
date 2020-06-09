package client

import (
	"context"
	"fmt"
)

// GetMethod wrap a client
type GetMethod method

// Call accept context, path and parameter
func (c *GetMethod) Call(ctx context.Context, path string, parameter map[string]string) (interface{}, error) {
	url, err := buildQueryParameter(fmt.Sprintf("%s%s", c.BaseURL, path), parameter)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

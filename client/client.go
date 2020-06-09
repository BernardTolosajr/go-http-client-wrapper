package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client is
type Client struct {
	client  *http.Client
	BaseURL *url.URL
	Get     *GetMethod
	Post    *PostMethod
	Put     *PutMethod
	Delete  *DeleteMethod
}

type method struct {
	*Client
}

// NewClient returns a new wraper http client. If a nil httpClient is
// provided, a new http.Client will be used.
func NewClient(base string, httpClient *http.Client) *Client {
	baseURL, _ := url.Parse(base)

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}

	c.Get = &GetMethod{c}
	c.Post = &PostMethod{c}
	c.Delete = &DeleteMethod{c}
	c.Put = &PutMethod{c}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do sends an API request and returns response. The API response is
// JSON decoded or returned as an error
// The provided ctx must be non-nil, if it is nil an error is returned. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request) (interface{}, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}

	req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	defer resp.Body.Close()

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func buildQueryParameter(uri string, parameter map[string]string) (string, error) {
	baseURL, err := url.Parse(uri)

	if err != nil {
		return "", err
	}

	if parameter == nil {
		return baseURL.String(), nil
	}

	if parameter != nil {
		params := url.Values{}
		for k, v := range parameter {
			params.Add(k, v)
		}
		baseURL.RawQuery = params.Encode()
	}

	return baseURL.String(), nil
}

package client

import "context"

type PutMethod method

func (p *PutMethod) Call(ctx context.Context, path string, body interface{}) (interface{}, error) {
	req, err := p.NewRequest("PUT", path, body)
	if err != nil {
		return nil, err
	}

	resp, err := p.Do(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

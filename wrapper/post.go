package wrapper

import "context"

type PostMethod method

func (p *PostMethod) Call(ctx context.Context, path string, body interface{}) (interface{}, error) {
	req, err := p.client.NewRequest("POST", path, body)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sample struct {
	Name string `json:"name"`
}

func TestPostMethod(t *testing.T) {
	ts := server(func(w http.ResponseWriter, r *http.Request) {
		var payload sample
		json.NewDecoder(r.Body).Decode(&payload)

		assert.Equal(t, "bt", payload.Name)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"foo": "bar"}`))
	})

	defer ts.Close()

	c := NewClient(HOST, nil)

	payload := &sample{
		Name: "bt",
	}

	resp, err := c.Post.Call(context.TODO(), "/foo", payload)
	if err != nil {
		panic(err)
	}

	result := resp.(map[string]interface{})

	assert.Equal(t, result["foo"], "bar")
}

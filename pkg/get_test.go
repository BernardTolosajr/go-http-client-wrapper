package pkg

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ts := server(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"foo": "bar"}`))
	})

	defer ts.Close()

	c := NewClient(HOST, nil)

	resp, err := c.Get.Call(context.TODO(), "/path", nil)
	if err != nil {
		panic(err)
	}

	result := resp.(map[string]interface{})

	assert.Equal(t, result["foo"], "bar")
}

func TestGetMethodWithParams(t *testing.T) {
	ts := server(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "foo=baz", r.URL.RawQuery)
		w.WriteHeader(http.StatusOK)
	})

	defer ts.Close()

	c := NewClient(HOST, nil)

	var params = make(map[string]string)
	params["foo"] = "baz"

	c.Get.Call(context.TODO(), "/path", params)
}

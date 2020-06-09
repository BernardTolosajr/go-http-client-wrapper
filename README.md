# go-http-client-wrapper

[![Test Status](https://github.com/BernardTolosajr/go-http-client-wrapper/workflows/Go/badge.svg)](https://github.com/BernardTolosajr/go-http-client-wrapper/actions?query=workflow%3AGo)
[![codecov](https://codecov.io/gh/BernardTolosajr/go-http-client-wrapper/branch/master/graph/badge.svg)](https://codecov.io/gh/BernardTolosajr/go-http-client-wrapper)

go-http-client-wrapper is a tiny go http client wrapper for calling API.

## Usage ##
```go
import "github.com/BernardTolosajr/go-http-client-wrapper/client" // with go modules enabled (GO111MODULE=on or 
```

Construct a new default client.

example:
```go
client := client.NewClient("http://yourbaseurl", nil)
// GET command with empty parameter
res, _ := client.Get.Call(context.Background(), "somepath", nil)

// GET command with parameter
var params = make(map[string]string)
params["foo"] = "baz"
res, _ := client.Get.Call(context.Background(), "somepath", params)

// POST command
sample := &sample{}
res, _ := client.Post.Call(context.Background(), "somepath", sample)

// PUT command
sample := &sample{}
res, _ := client.Put.Call(context.Background(), "somepath/id", sample)


// DELETE command
sample := &sample{}
res, _ := client.Delete.Call(context.Background(), "somepath/id", sample)
```

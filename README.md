# go-http-client-wrapper

go-http-client-wrapper is a Go http client wrapper for calling API.

## Usage ##
```go
import "github.com/BernardTolosajr/go-http-client-wrapper/client" // with go modules enabled (GO111MODULE=on or 
```

Construct a new default client.

example:
```go
client := client.NewClient("http://yourbaseurl", nil)
// GET command with empty parameter
res, _ := client.Get.Call(context.Background(), "/somepath", nil)

// GET command with parameter
var params = make(map[string]string)
params["foo"] = "baz"
res, _ := client.Get.Call(context.Background(), "/somepath", params)

// POST command
sample := &sample{}
res, _ := client.Post.Call(context.Background(), "/somepath", sample)
```

## TODO ##
 - DELETE method
 - PATCH method

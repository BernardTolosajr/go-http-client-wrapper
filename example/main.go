package main

import (
	"context"
	"fmt"

	"github.com/BernardTolosajr/go-http-client-wrapper/client"
)

func main() {
	client := client.NewClient("https://api.github.com/", nil)

	result, err := client.Get.Call(context.TODO(), "users/BernardTolosajr", nil)

	if err != nil {
		panic(err)
	}

	fmt.Printf("result %v", result)
}

package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()
	creds, err := transport.Creds(ctx, option.WithScopes(compute.CloudPlatformScope))
	CheckErr(err)

	token, err := creds.TokenSource.Token()
	CheckErr(err)

	fmt.Println(token)
}

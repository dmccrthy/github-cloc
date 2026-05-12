package main

import (
	"context"
	"fmt"
	"net/http"
)

//go:generate go run github.com/Khan/genqlient

func main() {
	ctx := context.Background()
	client := graphql.NewClient("https://api.github.com/graphql", http.DefaultClient)
	resp, err := getUser(ctx, client, "benjaminjkraft")
}

func Stuff() {
	fmt.Println("hi")
}

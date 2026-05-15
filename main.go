package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/dmccrthy/github-cloc/api"
)

func main() {
	ctx := context.Background()
	client := graphql.NewClient("https://api.github.com/graphql", http.DefaultClient)
	res, err := api.FindIssueID(ctx, client)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

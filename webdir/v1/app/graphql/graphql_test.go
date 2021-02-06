package graphql

import (
	"context"
	"log"
	"testing"

	"github.com/machinebox/graphql"
)

type ResponseStruct struct {
}

func TestGetQueryresp(t *testing.T) {

	// create a client (safe to share across requests)
	client := graphql.NewClient("https://machinebox.io/graphql")

	// make a request
	req := graphql.NewRequest(`
    query ($key: String!) {
        items (id:$key) {
            field1
            field2
            field3
        }
    }
`)

	// set any variables
	req.Var("key", "value")

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
}

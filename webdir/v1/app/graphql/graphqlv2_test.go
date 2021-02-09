package graphql

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"testing"
)

type graphQLRequest struct {
	Query     string `json:"query"`
	Variables string `json:"variables"`
}

func TestGraphqlV2(t *testing.T) {

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	client := http.Client{Transport: tr}

	query := `query {
					people {
						firstname
						lastname
						website
					}
				}`

	gqlMarshalled, err := json.Marshal(graphQLRequest{Query: query})
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "test.anosa.ga/graphql", bytes.NewBuffer(gqlMarshalled))
	if err != nil {
		panic(err)
	}
	t.Error(string(gqlMarshalled))

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	b, _ := httputil.DumpResponse(resp, true)
	t.Error(string(b))

	query = `query($number_of_repos:Int!) {
		viewer {
			name
			 repositories(last: $number_of_repos) {
				 nodes {
					 name
					 }
				}
			}
		}`

	variables := `variables {
						"number_of_repos": 3
						}`

	gqlMarshalled, err = json.Marshal(graphQLRequest{Query: query, Variables: variables})
	if err != nil {
		panic(err)
	}

	b, _ = httputil.DumpResponse(resp, true)
	t.Error(string(b))
}

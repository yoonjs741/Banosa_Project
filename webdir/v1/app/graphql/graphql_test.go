package graphql

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type ResponseStruct struct {
}

func TestGetQueryresp(t *testing.T) {
	jsonData := map[string]string{
		"query": `
			{
				people {
					firstname
					lastname
					website
					}
				}
			}
		`,
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", "https://test.anosa.ga/graphql", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	t.Error(string(data))
}

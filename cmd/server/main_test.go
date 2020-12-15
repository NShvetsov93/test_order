package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCaseNews struct {
	name      string
	body      []byte
	URL       string
	wantErr   bool
	errString string
	// err     error
}

const localURL = "http://localhost:8080"

func TestHttp(t *testing.T) {
	for _, item := range getTestCases() {
		t.Run(item.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPost, localURL+item.URL, bytes.NewBuffer(item.body))
			res, err := http.DefaultClient.Do(req)
			if item.wantErr {
				body, _ := ioutil.ReadAll(res.Body)
				assert.Equal(t, item.errString, string(body))
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func getTestCases() []testCaseNews {
	return []testCaseNews{
		{
			name:    "first",
			body:    []byte(`{"product_id": 1,"quantity": 1}`),
			URL:     "/store/add",
			wantErr: false,
		},
		{
			name:      "second",
			body:      []byte(`{"product_id": 1}`),
			URL:       "/store/add",
			wantErr:   true,
			errString: "quantity required\n",
		},
		{
			name:    "third",
			body:    []byte(`{"product_id": 1,"quantity": 1}`),
			URL:     "/store/order",
			wantErr: false,
		},
		{
			body:      []byte(`{"product_id": 1}`),
			URL:       "/store/order",
			wantErr:   true,
			errString: "quantity required\n",
		},
		{
			body:      []byte(`{"product_id": 1,"quantity": 100}`),
			URL:       "/store/order",
			wantErr:   true,
			errString: "Too many requests\n",
		},
	}
}

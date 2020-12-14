package main

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCaseNews struct {
	body    []byte
	URL     string
	wantErr bool
	// err     error
}

const localURL = "http://localhost:8080"

func TestAdd(t *testing.T) {
	for _, item := range getAddTestCases() {
		req, _ := http.NewRequest(http.MethodPost, localURL+item.URL, bytes.NewBuffer(item.body))
		_, err := http.DefaultClient.Do(req)
		if item.wantErr {
			require.Error(t, err)
			assert.Equal(t, true, true)
		} else {
			require.NoError(t, err)
			assert.Equal(t, true, true)
		}
	}
}

func getAddTestCases() []testCaseNews {
	return []testCaseNews{
		{
			body:    []byte(`{"product_id": 1,"quantity": 1}`),
			URL:     "/store/add",
			wantErr: false,
		},
		{
			body:    []byte(`{"product_id": 1}`),
			URL:     "/store/add",
			wantErr: true,
		},
	}
}

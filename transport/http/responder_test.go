package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ResponderCase struct {
	Name          string
	MockResponse  DummyObject
	MockErrorCode *string
	WantResponse  HTTPResponse[DummyObject]
}

type DummyObject struct {
	Content string
}

func TestResponder(t *testing.T) {
	cases := []ResponderCase{
		{
			Name: "single entity response",
			MockResponse: DummyObject{
				Content: "CASE 1",
			},
			MockErrorCode: nil,
			WantResponse: HTTPResponse[DummyObject]{
				Data: DummyObject{
					Content: "CASE 1",
				},
				Error: nil,
			},
		},
	}

	for _, tcase := range cases {
		resp := ResponseWithHTTP[DummyObject](tcase.MockResponse, tcase.MockErrorCode)
		assert.Equal(t, tcase.WantResponse, resp, "response are equals")
	}
}

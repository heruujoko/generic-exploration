package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ResponderCase struct {
	Name              string
	MockResponse      *DummyObject
	MockResponseSlice *[]DummyObject
	MockErrorCode     *string
	WantResponse      *HTTPResponse[DummyObject]
	WantResponseSlice *HTTPResponse[[]DummyObject]
}

type DummyObject struct {
	Content string
}

func TestResponder(t *testing.T) {
	mockErrCode := "ERR"
	cases := []ResponderCase{
		{
			Name: "single entity response",
			MockResponse: &DummyObject{
				Content: "CASE 1",
			},
			MockResponseSlice: nil,
			MockErrorCode:     nil,
			WantResponse: &HTTPResponse[DummyObject]{
				Data: DummyObject{
					Content: "CASE 1",
				},
				Error: nil,
			},
		},
		{
			Name: "slice entity response",
			MockResponseSlice: &[]DummyObject{
				{
					Content: "CASE 1",
				},
				{
					Content: "CASE 2",
				},
			},
			MockResponse:  nil,
			MockErrorCode: nil,
			WantResponseSlice: &HTTPResponse[[]DummyObject]{
				Data: []DummyObject{
					{
						Content: "CASE 1",
					},
					{
						Content: "CASE 2",
					},
				},
				Error: nil,
			},
		},
		{
			Name: "slice error response",
			MockResponseSlice: &[]DummyObject{
				{
					Content: "CASE 1",
				},
				{
					Content: "CASE 2",
				},
			},
			MockResponse:  nil,
			MockErrorCode: &mockErrCode,
			WantResponseSlice: &HTTPResponse[[]DummyObject]{
				Data: []DummyObject{
					{
						Content: "CASE 1",
					},
					{
						Content: "CASE 2",
					},
				},
				Error: &ErrorInResponse{
					Code: mockErrCode,
				},
			},
		},
	}

	for _, tcase := range cases {
		if tcase.WantResponse != nil {
			resp := ResponseWithHTTP[DummyObject](*tcase.MockResponse, tcase.MockErrorCode)
			assert.Equal(t, *tcase.WantResponse, resp, "response are equals")
		}

		if tcase.WantResponseSlice != nil {
			resp := ResponseWithHTTP[[]DummyObject](*tcase.MockResponseSlice, tcase.MockErrorCode)
			assert.Equal(t, *tcase.WantResponseSlice, resp, "slice response are equals")
		}
	}
}

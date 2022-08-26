package rest_client

import (
	"github.com/stretchr/testify/assert"
	"go-foundation/pkg/http_client_mock"
	"net/http"
	"testing"
)

var (
	endpont = "http://localhost/endpoint"
)

func TestPostPessoaPass(t *testing.T) {
	var expectedResponse = `[{"test": 0.9,"test2": 0.028079710144927536}]`
	httpClientMock := &httpclientmock.Mock{
		ResponseBody:   expectedResponse,
		ResponseStatus: 200,
	}
	c := NewRestClient(httpClientMock, endpont)
	c.Post("Bearer X", "12345678912")
	assert.Equal(t, http.MethodPost, httpClientMock.RequestMethod)
	assert.Equal(t, http.StatusOK, httpClientMock.ResponseStatus)
	assert.Equal(t, endpont, httpClientMock.RequestURL)
	assert.Equal(t, expectedResponse, httpClientMock.ResponseBody)
}

func TestPostPessoaFail(t *testing.T) {
	var expectedResponse = `{"error": "Unauthorized: \"credential not found\""}`
	httpClientMock := &httpclientmock.Mock{
		ResponseBody:   expectedResponse,
		ResponseStatus: 401,
	}
	c := NewRestClient(httpClientMock, endpont)
	c.Post("Bearer X", "12345678912")
	assert.Equal(t, http.MethodPost, httpClientMock.RequestMethod)
	assert.Equal(t, http.StatusUnauthorized, httpClientMock.ResponseStatus)
	assert.Equal(t, endpont, httpClientMock.RequestURL)
	assert.Equal(t, expectedResponse, httpClientMock.ResponseBody)
}

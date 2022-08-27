package domain

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mock "go-foundation/pkg/mocks"
	"go-foundation/pkg/rest_client"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	endpont = "mock"
)

func TestPostPass(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var expectedResponse, _ = json.Marshal(`{}`)
	var expectedResponseApi = `[{"cpf":"test"}]`
	input := io.NopCloser(strings.NewReader(`[{"cpf":"test"}]`))
	httpClientMock := &mock.Mock{
		ResponseBody:   expectedResponseApi,
		ResponseStatus: 200,
	}

	c := rest_client.NewRestClient(httpClientMock)
	s := NewEndpoint(c)
	w := httptest.NewRecorder()
	cxt, _ := gin.CreateTestContext(w)
	cxt.Request = &http.Request{
		RequestURI: endpont,
		Header:     make(http.Header),
		Body:       input,
		Method:     http.MethodPost,
	}
	s.Post(cxt)
	assert.Equal(t, http.StatusOK, w.Code)
	require.JSONEq(t, string(expectedResponse), w.Body.String())
}

func TestPostFail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var expectedResponse, _ = json.Marshal(`{}`)
	var expectedResponseApi = `{"error": "Unauthorized: \"credential not found\""}`
	input := io.NopCloser(strings.NewReader(`[{"cpf":"test"}]`))
	httpClientMock := &mock.Mock{
		ResponseBody:   expectedResponseApi,
		ResponseStatus: 401,
	}

	c := rest_client.NewRestClient(httpClientMock)
	s := NewEndpoint(c)
	w := httptest.NewRecorder()
	cxt, _ := gin.CreateTestContext(w)
	cxt.Request = &http.Request{
		RequestURI: endpont,
		Header:     make(http.Header),
		Body:       input,
		Method:     http.MethodPost,
	}
	s.Post(cxt)
	assert.Equal(t, http.StatusOK, w.Code)
	require.JSONEq(t, string(expectedResponse), w.Body.String())
}

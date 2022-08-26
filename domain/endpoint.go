package domain

import (
	"github.com/gin-gonic/gin"
	"go-foundation/pkg/rest_client"
	"net/http"
)

type Endpoint struct {
	httpClient *rest_client.RestClient
}

func NewEndpoint(httpClient *rest_client.RestClient) *Endpoint {
	return &Endpoint{
		httpClient: httpClient,
	}
}

func (s *Endpoint) Post(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "{}")
}

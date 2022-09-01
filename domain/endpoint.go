package domain

import (
	"github.com/edson-dev/go-foundation/pkg/logger"
	"github.com/edson-dev/go-foundation/pkg/rest_client"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Endpoint struct {
	restClient *rest_client.RestClient
}

func NewEndpoint(restClient *rest_client.RestClient) *Endpoint {
	return &Endpoint{
		restClient: restClient,
	}
}

func (s *Endpoint) Post(c *gin.Context) {
	v, e := s.restClient.Post("", "")
	logger.Throw(e)
	c.IndentedJSON(http.StatusOK, string(v))
}

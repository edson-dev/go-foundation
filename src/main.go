package main

import (
	"github.com/edson-dev/go-foundation/domain"
	"github.com/edson-dev/go-foundation/pkg/http_client"
	"github.com/edson-dev/go-foundation/pkg/json"
	"github.com/edson-dev/go-foundation/pkg/logger"
	"github.com/edson-dev/go-foundation/pkg/rest_client"
	"github.com/gin-gonic/gin"
)

func main() {
	s := `{"Name": "India", "test": "Test123"`
	_, err := json.ImportString(s)
	logger.Throw(err)
	router := gin.New()
	client := rest_client.NewRestClient(httpclient.NewHTTPClient(10))
	endpoint := domain.NewEndpoint(client)

	router.GET("/endpoint", endpoint.Post)
	router.GET("/endpoint2", endpoint.Post)
	router.GET("/endpoint3", endpoint.Post)
	router.Run("localhost:8080")
}

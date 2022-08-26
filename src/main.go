package main

import (
	"github.com/gin-gonic/gin"
	"go-foundation/domain"
	httpclient "go-foundation/pkg/http_client"
	"go-foundation/pkg/rest_client"
)

func main() {
	router := gin.New()
	client := rest_client.NewRestClient(httpclient.NewHTTPClient(10))
	endpoint := domain.NewEndpoint(client)

	router.GET("/endpoint", endpoint.Post)
	router.GET("/endpoint2", endpoint.Post)
	router.GET("/endpoint3", endpoint.Post)
	router.Run("localhost:8080")
}

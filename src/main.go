package main

import (
	"fmt"
	"github.com/edson-dev/go-foundation/docs"
	"github.com/edson-dev/go-foundation/domain"
	"github.com/edson-dev/go-foundation/pkg/http_client"
	"github.com/edson-dev/go-foundation/pkg/json"
	"github.com/edson-dev/go-foundation/pkg/logger"
	middlewares "github.com/edson-dev/go-foundation/pkg/midlewares"
	"github.com/edson-dev/go-foundation/pkg/rest_client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	logger.SetLevel("")
}
func main() {
	s := `{"Name": "India", "test": "Test123"`
	_, err := json.ImportString(s)
	println(logger.Level.String())
	logger.Throw(err)
	router := gin.New()
	router.LoadHTMLGlob("docs/*")
	client := rest_client.NewRestClient(httpclient.NewHTTPClient(10))
	endpoint := domain.NewEndpoint(client)
	router.Use(middlewares.Auth())
	router.GET("/endpoint", endpoint.Post)
	router.GET("/swagger.json", func(c *gin.Context) {
		c.HTML(http.StatusOK, "swagger.json", gin.H{
			"title": "json doc",
		})
	})
	router.GET("/docs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "swagger.html", gin.H{
			"authorization": c.Request.Header.Get("Authorization"),
		})
	})
	router.GET("/swagger", func(c *gin.Context) {
		c.Data(http.StatusOK, "yaml", []byte(fmt.Sprintf(docs.Page, docs.Def)))
	})
	router.Run("localhost:8080")
}

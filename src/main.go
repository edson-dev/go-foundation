package main

import (
	"fmt"
	"go-foundation/pkg/json"
)

func main() {
	s := `{"Name": "India", "test": "Test123"}`
	j := json.ImportString(s)
	fmt.Println(j.ExportString())
	/*router := gin.New()
	client := rest_client.NewRestClient(httpclient.NewHTTPClient(10))
	endpoint := domain.NewEndpoint(client)

	router.GET("/endpoint", endpoint.Post)
	router.GET("/endpoint2", endpoint.Post)
	router.GET("/endpoint3", endpoint.Post)
	router.Run("localhost:8080")*/
}

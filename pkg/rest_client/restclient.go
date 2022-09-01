package rest_client

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff"
	httpclient "github.com/edson-dev/go-foundation/pkg/http_client"
	"io"
	"net/http"
	"strings"
	"time"
)

type RestClient struct {
	httpClient httpclient.HTTPClient
	Url        string
	Body       string
}

func NewRestClient(client httpclient.HTTPClient) *RestClient {
	return &RestClient{
		httpClient: client,
		Url:        `www.google.com.br`,
		Body:       `{"body": ["test"],"size": 10,"query": {"this": %s}}`,
	}
}
func (rc *RestClient) Post(autorization string, id string) (data []byte, err error) {
	ebo := backoff.NewExponentialBackOff()
	ebo.MaxInterval = 10 * time.Second
	err = backoff.Retry(func() error {
		req, err := http.NewRequest(http.MethodPost, rc.Url, strings.NewReader(fmt.Sprintf(rc.Body, id)))
		req.Header = MakeHashHeader([]Header{{Key: "Autorization", Value: "Bearer X"}})
		if err != nil {
			fmt.Sprintf("Error on request:{%s}", err.Error())
			return err
		}
		res, err := rc.httpClient.Do(req)
		if err != nil {
			fmt.Sprintf("Error on response:{%s}", err.Error())
			return err
		}
		result, err2 := io.ReadAll(res.Body)
		defer func() {
			err = res.Body.Close()
			if err != nil || err2 != nil {
				fmt.Sprintf(err.Error())
			}
		}()
		data = result
		if res.StatusCode != 200 {
			err = fmt.Errorf("api error (%s) occurred", res.Status)
			return err
		}
		return nil
	}, backoff.WithContext(backoff.WithMaxRetries(ebo, uint64(5)), context.Background()))
	return data, err
}

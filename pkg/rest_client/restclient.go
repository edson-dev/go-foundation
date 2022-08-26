package rest_client

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff"
	"go-foundation/pkg/http_client"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type RestClient struct {
	httpClient httpclient.HTTPClient
	url        string
	body       string
}

func NewRestClient(client httpclient.HTTPClient, url string) *RestClient {
	return &RestClient{
		httpClient: client,
		url:        url,
		body:       `{"body": ["test"],"size": 10,"query": {"this": %s}}`,
	}
}
func (rc *RestClient) Post(autorization string, id string) (data []byte, err error) {
	ebo := backoff.NewExponentialBackOff()
	ebo.MaxInterval = 10 * time.Second
	err = backoff.Retry(func() error {
		req, err := http.NewRequest(http.MethodPost, rc.url, strings.NewReader(fmt.Sprintf(rc.body, id)))
		req.Header.Add("Authorization", autorization)
		if err != nil {
			fmt.Sprintf("Error on request:{%s}", err)
			return err
		}
		res, err := rc.httpClient.Do(req)
		if err != nil {
			fmt.Sprintf("Error on response:{%s}", err)
			return err
		}
		result, err2 := ioutil.ReadAll(res.Body)
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

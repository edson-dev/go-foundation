package httpclient_test

import (
	"github.com/edson-dev/go-foundation/pkg/http_client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewHTTPClient(t *testing.T) {
	client := httpclient.NewHTTPClient(time.Minute)
	assert.NotNil(t, client)
}

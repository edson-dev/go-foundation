package httpclient_test

import (
	"github.com/stretchr/testify/assert"
	"go-foundation/pkg/http_client"
	"testing"
	"time"
)

func TestNewHTTPClient(t *testing.T) {
	client := httpclient.NewHTTPClient(time.Minute)
	assert.NotNil(t, client)
}

package rest_client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeHashHeader(t *testing.T) {
	e := []Header{{Key: "Autorization", Value: "Bearer X"}}
	set := MakeHashHeader(e)
	assert.NotEmpty(t, set)
	assert.Equal(t, []string{"Bearer X"}, set["Autorization"])
}

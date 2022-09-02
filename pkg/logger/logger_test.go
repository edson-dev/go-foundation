package logger

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewLoggerLevel(t *testing.T) {
	assert.NotEqual(t, Instance, nil)
	assert.Equal(t, Level.String(), zap.DebugLevel.String())
}

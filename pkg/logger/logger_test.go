package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLoggerLevel(t *testing.T) {
	assert.NotEqual(t, Instance, nil)
	assert.Equal(t, Instance.Level, logrus.DebugLevel)
}

func TestLogger(t *testing.T) {
	logger, hook := test.NewNullLogger()
	logger.Error("Helloerror")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "Helloerror", hook.LastEntry().Message)

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}

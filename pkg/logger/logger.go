package logger

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

var Instance *logrus.Logger

func init() {
	Instance = logrus.New()
	Instance.Formatter = &logrus.JSONFormatter{}
	Instance.SetLevel(logrus.DebugLevel)
}

func Throw(err error) {
	if err != nil {
		pc, _, line, _ := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		Instance.WithFields(logrus.Fields{
			"file": details.Name(),
			"line": line,
			"type": "relevant",
		}).Errorln(err.Error())
	}
}
func DEBUG(message string) {
	Instance.Debug(message)
}

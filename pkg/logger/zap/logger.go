package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strings"
)

var Instance *zap.Logger
var Level zap.AtomicLevel

func init() {
	Level = zap.NewAtomicLevel()
	Instance = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		Level,
	))
	defer Instance.Sync()
}
func SetLevel(level string) {
	switch level {
	case "fatal":
		Level.SetLevel(zap.FatalLevel)
	case "error":
		Level.SetLevel(zap.ErrorLevel)
	case "warn":
		Level.SetLevel(zap.WarnLevel)
	case "info":
		Level.SetLevel(zap.InfoLevel)
	default:
		Level.SetLevel(zap.DebugLevel)
	}
}
func GetLevel() string {
	return Level.String()
}
func Throw(err error) {
	if err != nil {
		pc, _, line, _ := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		fileStructure := strings.Split(details.Name(), "/")
		Instance.Error(err.Error(),
			zap.String("file", fileStructure[len(fileStructure)-1]),
			zap.Int("line", line),
			zap.String("type", "relevant"))
	}
}
func DEBUG(message string) {
	Instance.Debug(message)
}

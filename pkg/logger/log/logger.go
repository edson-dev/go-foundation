package logger

import (
	"log"
)

var Instance *log.Logger
var Level string

func init() {
	Level = "debug"
	Instance = log.Default()
}
func SetLevel(level string) {

}
func GetLevel() string {
	return "DEBUG"
}

func Throw(err error) {
	Instance.Print(err.Error())
}
func DEBUG(message string) {
	Instance.Print(message)
}

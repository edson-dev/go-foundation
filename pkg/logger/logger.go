package interfaces

type Logger interface {
	SetLevel(level string)
	Throw(err error)
	DEBUG(msg string)
	GetLevel() string
}

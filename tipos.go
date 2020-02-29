package monologger

//LogLevel identifica o nivel de log
type LogLevel int

//Emergency sistema está parado/exposto
const (
	Emergency LogLevel = iota + 1
	Alert
	Critical
	Error
	Warning
	Info
	Debug
)

//LoggerRFC5424 indica os metodos basicos de log
type LoggerRFC5424 interface {
	Emergency(message string, params ...interface{})
	Alert(message string, params ...interface{})
	Critical(message string, params ...interface{})
	Error(message string, params ...interface{})
	Warning(message string, params ...interface{})
	Info(message string, params ...interface{})
	Debug(message string, params ...interface{})
	Log(level LogLevel, message string, params ...interface{})
}

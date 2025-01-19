package pocketlog

// Logger is used to log information
type Logger struct {
	threshold Level
}

// New returns you a logger, ready to log at the required threshold
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	//TODO implement
}

// Infof formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	// TODO implement
}

// Warnf formats and prints a message if the log level is warn or higher
func (l *Logger) Warnf(format string, args ...any) {
	// TODO implement
}

// Errorf formats and prints a message if the log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {
	//TODO implement
}

// Fatalf formats and prints a message if the log level is fatal or higher
func (l *Logger) Fatalf(format string, args ...any) {
	//TODO implement
}

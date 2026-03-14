// stub for testing
package zap

type Logger struct{}

func NewProduction() (*Logger, error) { return &Logger{}, nil }

func (l *Logger) Info(msg string, fields ...Field)  {}
func (l *Logger) Error(msg string, fields ...Field) {}
func (l *Logger) Debug(msg string, fields ...Field) {}
func (l *Logger) Warn(msg string, fields ...Field)  {}

func (l *Logger) With(fields ...Field) *Logger { return l }
func (l *Logger) Named(s string) *Logger       { return l }
func (l *Logger) Sugar() *SugaredLogger        { return &SugaredLogger{} }

type SugaredLogger struct{}

type Field struct{}

func String(key, val string) Field  { return Field{} }
func Int(key string, val int) Field { return Field{} }

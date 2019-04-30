/*
Package picolo is a simple logger for Go.
This minimalist package has :
	- levels
	- configurable output (io.Writer)
	- configurable timeformat
	- prefix for messages
	- sub loggers

For most of the cases the defaults will be more than enough.

	l := picolo.New() // Use defaults
	l.Infof("Info message")

The constuctor function picolo.New accepts options to override defaults:

	// constructor with optional prefix
	l = picolo.New(picolo.WithLevel(LevelDebug), picolo.WithPrefix("[some-prefix]"))
	l.Infof("Info message")
*/
package picolo

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Level is the log level
type Level int

const (
	// Log levels
	LevelDebug Level = iota + 1
	LevelInfo
	LevelWarning
	LevelError
)

// DefaultTimeFormat is our default time format
const DefaultTimeFormat = "2006-01-02 15:04:05.000"

// ErrUnknownLevel is returned by LevelFromString
var ErrUnknownLevel = fmt.Errorf("unknown log level")

// Logger is our logger struct
type Logger struct {
	mu   sync.Mutex // used to synchronize output as well
	opts *options
}

type options struct {
	level      Level
	output     io.Writer
	prefix     string
	timeFormat string
	timeUTC    bool
}

// Option is our option type
type Option func(*options)

// New creates a new Logger
func New(opts ...Option) *Logger {
	loggerOpts := options{
		level:      LevelInfo,
		output:     os.Stdout,
		timeFormat: DefaultTimeFormat,
		timeUTC:    true,
	}

	if len(opts) != 0 {
		for _, opt := range opts {
			opt(&loggerOpts)
		}
	}

	return &Logger{
		opts: &loggerOpts,
	}
}

// WithLevel sets the log level
func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

// WithOutput sets log output
func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

// WithPrefix sets a log prefix
func WithPrefix(prefix string) Option {
	return func(o *options) {
		if prefix != "" {
			prefix += " "
		}
		o.prefix = prefix
	}
}

// WithTimeFormat sets the time format. Specify empty time format to disable datetime in logs.
func WithTimeFormat(format string, utc bool) Option {
	return func(o *options) {
		o.timeFormat = format
		o.timeUTC = utc
	}
}

// NewFrom creates a new logger from a logger, appending the prefix
func NewFrom(from *Logger, morePrefix string) *Logger {
	pf := from.opts.prefix
	if morePrefix != "" {
		pf += morePrefix + " "
	}

	loggerOpts := options{
		prefix: pf,

		level:      from.opts.level,
		output:     from.opts.output,
		timeFormat: from.opts.timeFormat,
		timeUTC:    from.opts.timeUTC,
	}

	return &Logger{
		opts: &loggerOpts,
	}
}

func (l *Logger) fmt(level Level, msg string) string {
	var line string

	if l.opts.timeFormat != "" {
		if l.opts.timeUTC {
			line = time.Now().UTC().Format(l.opts.timeFormat) + " "
		} else {
			line = time.Now().Format(l.opts.timeFormat) + " "
		}
	}

	line += level.String() + " " + l.opts.prefix + msg + "\n"

	return line
}

func (l *Logger) write(level Level, format string, a ...interface{}) {
	if l.opts.level > level || l.opts.output == nil {
		return
	}

	formatted := l.fmt(level, fmt.Sprintf(format, a...))

	l.mu.Lock()
	// ignore write errors
	_, _ = l.opts.output.Write([]byte(formatted))
	l.mu.Unlock()
}

// Debugf logs formatted message in debug level
func (l *Logger) Debugf(format string, a ...interface{}) {
	l.write(LevelDebug, format, a...)
}

// Infof logs formatted message in info level
func (l *Logger) Infof(format string, a ...interface{}) {
	l.write(LevelInfo, format, a...)
}

// Warningf logs formatted message in warning level
func (l *Logger) Warningf(format string, a ...interface{}) {
	l.write(LevelWarning, format, a...)
}

// Errorf logs formatted message in error level
func (l *Logger) Errorf(format string, a ...interface{}) {
	l.write(LevelError, format, a...)
}

func (l *Logger) SetLogLevel(s string) error {
	level, err := LevelFromString(s)
	if err != nil {
		return err
	}
	l.opts.level = level
	return nil
}

// String returns the log level in string representation
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// LevelFromString returns the log level from string representation
func LevelFromString(s string) (Level, error) {
	switch s {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warning":
		return LevelWarning, nil
	case "error":
		return LevelError, nil
	default:
		return 0, ErrUnknownLevel
	}
}

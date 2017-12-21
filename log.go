package picolo

import (
	"fmt"
	"io"
	"sync"
	"time"
)

// Level is the log level
type Level int

const (
	// Log levels

	LevelDebug   Level = iota // Debug level
	LevelInfo                 // Info level
	LevelWarning              // Warning level
	LevelError                // Error level
)

// Option is a bitmask of options
type Option int

const (
	// Bits or'ed together to control what's printed.
	OptDateTime Option = 1 << iota // Include datetime in log line
	OptUTC                         // Use UTC datetime

	OptDefault = OptDateTime | OptUTC // Default values if no options given
)

const (
	timeFormat = "2006-01-02 15:04:05.000 "
)

// Logger is our logger struct
type Logger struct {
	level  Level
	opts   Option
	output io.Writer
	prefix string

	// mu is used to synchronize output as well as protect the above fields
	mu sync.Mutex
}

// New creates a new Logger
func New(lev Level, output io.Writer, opts ...Option) *Logger {
	var finalOpts Option

	if len(opts) == 0 {
		// No options given, use OptDefault
		finalOpts = OptDefault
	} else {
		for _, o := range opts {
			finalOpts |= o
		}
	}

	l := Logger{
		level:  lev,
		opts:   finalOpts,
		output: output,
	}
	return &l
}

// Prefix sets the current logger's prefix
func (l *Logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if prefix != "" {
		prefix += " "
	}
	l.prefix = prefix
}

// NewFrom creates a new logger from a logger, appending the prefix
func NewFrom(from *Logger, morePrefix string) *Logger {
	pf := from.prefix
	if morePrefix != "" {
		pf += morePrefix + " "
	}

	l := &Logger{
		level:  from.level,
		opts:   from.opts,
		output: from.output,
		prefix: pf,
	}

	return l
}

func (l *Logger) fmt(level Level, msg string) string {
	var line string

	if l.opts&OptDateTime > 0 {
		if l.opts&OptUTC > 0 {
			line = time.Now().UTC().Format(timeFormat)
		} else {
			line = time.Now().Format(timeFormat)
		}
	}

	line += level.String() + " " + l.prefix + msg + "\n"

	return line
}

func (l *Logger) write(level Level, format string, a ...interface{}) {
	if l.level > level || l.output == nil {
		return
	}

	formatted := l.fmt(level, fmt.Sprintf(format, a...))

	l.mu.Lock()
	l.output.Write([]byte(formatted))
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

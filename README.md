![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)
![Tag](https://img.shields.io/github/tag/peakgames/picolo.svg)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/peakgames/picolo)
[![Go Report](https://goreportcard.com/badge/github.com/peakgames/picolo)](https://goreportcard.com/report/github.com/peakgames/picolo)

# picolo

Minimalistic logging library.

# Levels

	LevelDebug
	LevelInfo
	LevelWarning
	LevelError

# Options

The constructor accepts several options:

    func WithLevel(level Level) Option
    func WithOutput(output io.Writer) Option
    func WithPrefix(prefix string) Option
    func WithTimeFormat(format string, utc bool) Option

## Defaults

If no options are given, the following are assumed.

    WithLevel(LevelInfo)
    WithOutput(os.Stdout)
    WithTimeFormat(DefaultTimeFormat, true)

# Usage

```go
		l := picolo.New(WithPrefix("[some-prefix]")) // constructor with optional prefix
		l.Debugf("Debug message")
		// TIME LEVEL [some-prefix] Debug message
		
		// Create sub-logger, appending prefix
		k := picolo.NewFrom(l, "[more-prefix]")
		k.Debugf("Debug message")
		// TIME LEVEL [some-prefix] [more-prefix] Debug message
```

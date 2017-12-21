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

	OptDateTime  // Include datetime in log line
	OptUTC       // Use UTC datetime

	OptDefault = OptDateTime | OptUTC

# Timestamp

Always in `YYYY-MM-DD hh:mm:ss.msec`. Can be turned off in options (use `0`)

# Usage

```go
		l := picolo.New(LevelDebug, os.Stdout) // level, io.Writer, [option ...]
		l.SetPrefix("[some-prefix]") // optional
		l.Debugf("Debug message")
		// TIME LEVEL [some-prefix] Debug message

		k := picolo.NewFrom(l, "[more-prefix]")
		k.Debugf("Debug message")
		// TIME LEVEL [some-prefix] [more-prefix] Debug message
```

![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)
![Tag](https://img.shields.io/github/tag/peak/picolo.svg)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/peak/picolo)
[![Go Report](https://goreportcard.com/badge/github.com/peak/picolo)](https://goreportcard.com/report/github.com/peak/picolo)
[![Build Status](https://travis-ci.org/peak/picolo.svg?branch=master)](https://travis-ci.org/peak/picolo)

# picolo

picolo is a minimalistic logging library for go.

# Usage
If no options are given, picolo assumes:
* log level is `INFO`
* output is `os.Stdout`
* time format is `2006-01-02 15:04:05.000`
* prefix is empty string

```go
l := picolo.New() // Use defaults
l.Infof("Info message")
// 2019-04-29 15:34:32.166 INFO Info message
```

`prefix` can be set with `picolo.WithPrefix` option.

```go
l = picolo.New(picolo.WithPrefix("[some-prefix]")) // constructor with optional prefix
l.Infof("Info message")
// 2019-04-29 22:23:24.256 INFO [some-prefix] Info message
```

Sub loggers can be created from an existing logger with `picolo.NewFrom` constructor function.
```go
// Create sub-logger, appending prefix
k := picolo.NewFrom(l, "[more-prefix]")
k.Errorf("Error message: %v", err)
//  2019-04-29 23:24:25.267 ERROR [some-prefix] [more-prefix] Error message: No such file or directory
```

# Log Levels
picolo supports 4 types of log levels:
* DEBUG
* INFO
* WARNING
* ERROR

# Options

The constructor accepts several options:

    WithLevel(level Level)                     // Set level
    WithOutput(output io.Writer)               // Set output
    WithPrefix(prefix string)                  // Set prefix
    WithTimeFormat(format string, utc bool)    // Set time format and UTC flag


# Helpers

Use `picolo.LevelFromString` to parse a string into a log level. This can be used like:

```go
// ...
l := flag.String("logLevel", "debug", "Log level")
flag.Parse()

lvl, err := picolo.LevelFromString(*l)
if err != nil {
	// Unknown log level
}

logger := picolo.New(picolo.WithLevel(lvl))
logger.Infof("Logger is ready.")
// ...
```

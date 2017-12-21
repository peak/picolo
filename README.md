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

    WithLevel(level Level)                     // Set level
    WithOutput(output io.Writer)               // Set output
    WithPrefix(prefix string)                  // Set prefix
    WithTimeFormat(format string, utc bool)    // Set time format and UTC flag

## Defaults

If no options are given, the following are assumed.

    WithLevel(LevelInfo)
    WithOutput(os.Stdout)
    WithTimeFormat(DefaultTimeFormat, true)

The default time format is `2006-01-02 15:04:05.000`.

# Usage

```go
l := picolo.New() // Use defaults
l.Infof("Info message")
// 2017-12-21 22:23:24.256 INFO Info message

l = picolo.New(picolo.WithPrefix("[some-prefix]")) // constructor with optional prefix
l.Infof("Info message")
// 2017-12-21 22:23:24.256 INFO [some-prefix] Info message

// Create sub-logger, appending prefix
k := picolo.NewFrom(l, "[more-prefix]")
k.Errorf("Error message: %v", err)
//  2017-12-21 23:24:25.267 ERROR [some-prefix] [more-prefix] Error message: No such file or directory
```

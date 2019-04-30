package picolo

import (
	"bytes"
	"testing"
)

func TestBasic(t *testing.T) {
	var b bytes.Buffer

	l := New(WithLevel(LevelDebug), WithTimeFormat("", false), WithPrefix("[test-debug]"), WithOutput(&b))
	l.Debugf("Debug message")
	l.Errorf("Error message")

	result := b.Bytes()
	golden := []byte(`DEBUG [test-debug] Debug message
ERROR [test-debug] Error message
`)
	if !bytes.Equal(golden, result) {
		t.Errorf("Basic test: Got: %q Want: %q", string(result), string(golden))
	}
}

func TestLevel(t *testing.T) {
	var b bytes.Buffer

	l := New(WithLevel(LevelWarning), WithTimeFormat("", false), WithOutput(&b))
	l.Debugf("Debug message")
	l.Infof("Info message")
	l.Warningf("Warning message")
	l.Errorf("Error message")

	result := b.Bytes()
	golden := []byte("WARNING Warning message\nERROR Error message\n")
	if !bytes.Equal(golden, result) {
		t.Errorf("Got: %q Want: %q", string(result), string(golden))
	}

	l.SetLogLevel("info")
	l.Debugf("this debug line will be skipped.")
	l.Infof("info message should be written")

	result = b.Bytes()
	golden = append(golden, "INFO info message should be written\n"...)
	if !bytes.Equal(golden, result) {
		t.Errorf("Got: %q Want: %q", string(result), string(golden))
	}
}

func TestCreateSubLogger(t *testing.T) {
	var b bytes.Buffer

	l := New(WithLevel(LevelDebug), WithTimeFormat("", false), WithPrefix("[test-debug]"), WithOutput(&b))

	golden := []byte("DEBUG [test-debug] [omg sublogger] Debug message\n")
	NewFrom(l, "[omg sublogger]").Debugf("Debug message")
	result := b.Bytes()
	if !bytes.Equal(golden, result) {
		t.Errorf("Basic test: Got: %q Want: %q", string(result), string(golden))
	}
}

func TestLevelFromStringTableTest(t *testing.T) {
	tests := []struct {
		name  string
		input string
		level Level
		err   error
	}{
		{"level info", "info", LevelInfo, nil},
		{"level error", "error", LevelError, nil},
		{"unknown level", "invalid level", 0, ErrUnknownLevel},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			level, err := LevelFromString(test.input)
			if err != test.err {
				t.Errorf("expected error %v, got %v", test.err, err)
			}
			if level != test.level {
				t.Errorf("expected level %v, got %v", test.level, level)
			}
		})
	}
}

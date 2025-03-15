package logger

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	var buf bytes.Buffer

	testLogger := &DefaultLogger{
		level:  DEBUG,
		logger: log.New(&buf, "", 0),
		colorFuncs: map[LogLevel]func(a ...interface{}) string{
			DEBUG: func(a ...interface{}) string {
				return "DEBUG"
			},
			INFO: func(a ...interface{}) string {
				return "INFO"
			},
			SUCCESS: func(a ...interface{}) string {
				return "SUCCESS"
			},
			WARN: func(a ...interface{}) string {
				return "WARN"
			},
			ERROR: func(a ...interface{}) string {
				return "ERROR"
			},
		},
	}

	tests := []struct {
		logFunc  func(msg string, args ...interface{})
		message  string
		args     []interface{}
		expected string
	}{
		{testLogger.Debug, "Debugging... %d", []interface{}{123}, "[DEBUG] Debugging... 123\n"},
		{testLogger.Info, "Information... %s", []interface{}{"Info"}, "[INFO] Information... Info\n"},
		{testLogger.Success, "Success! %s", []interface{}{"Ok"}, "[SUCCESS] Success! Ok\n"},
		{testLogger.Warn, "Warning! %s", []interface{}{"Be careful"}, "[WARN] Warning! Be careful\n"},
		{testLogger.Error, "Error! %s", []interface{}{"Critical"}, "[ERROR] Error! Critical\n"},
	}

	for _, test := range tests {
		buf.Reset()
		test.logFunc(test.message, test.args...)
		assert.Equal(t, test.expected, buf.String(), "Log output shoud match")
	}
}

package httputil

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

var validStackFuncs = []func(string) bool{
	func(file string) bool {
		return strings.Contains(file, "/mongodb/api/")
	},
}

// RuntimeCallerStack returns the app's `file:line` stacktrace
// to give more information about an error cause.
func RuntimeCallerStack() (s string) {
	var pcs [10]uintptr
	n := runtime.Callers(1, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		for _, fn := range validStackFuncs {
			if fn(frame.File) {
				s += fmt.Sprintf("\n\t\t\t%s:%d", frame.File, frame.Line)
			}
		}

		if !more {
			break
		}
	}

	return s
}

// HTTPError describes an HTTP error.
type HTTPError struct {
	error
	Stack       string    `json:"-"` // the whole stacktrace.
	CallerStack string    `json:"-"` // the caller, file:lineNumber
	When        time.Time `json:"-"` // the time that the error occurred.
	// ErrorCode int: maybe a collection of known error codes.
	StatusCode int `json:"statusCode"`
	// could be named as "reason" as well
	//  it's the message of the error.
	Description string `json:"description"`
}

func newError(statusCode int, err error, format string, args ...interface{}) HTTPError {
	if format == "" {
		format = http.StatusText(statusCode)
	}

	desc := fmt.Sprintf(format, args...)
	if err == nil {
		err = errors.New(desc)
	}

	return HTTPError{
		err,
		string(debug.Stack()),
		RuntimeCallerStack(),
		time.Now(),
		statusCode,
		desc,
	}
}

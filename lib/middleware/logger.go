package middleware

import (
	"github.com/yarf-framework/yarf"
	"log"
	"net/http"
)

// LoggerWriter will replace (wrap) the http.ResponseWriter to log all content written to the response.
type LoggerWriter struct {
	StatusCode int
	Writer     http.ResponseWriter
}

// Header is a wrapper for http.ResponseWriter.Header()
func (lw *LoggerWriter) Header() http.Header {
	return lw.Writer.Header()
}

// WriteHeader is a wrapper for http.ResponseWriter.WriteHeader()
// It saves the status code to be returned so we can log it.
func (lw *LoggerWriter) WriteHeader(code int) {
	lw.StatusCode = code

	lw.Writer.WriteHeader(code)
}

// Write is a wrapper for
func (lw *LoggerWriter) Write(content []byte) (int, error) {
	return lw.Writer.Write(content)
}

// Logger middleware it's a simple log module that uses the default golang's log package.
// The log output writer can be defined by default with the log.SetOutput(w io.Writer) function.
// For more complex environments where a default logger can't be used across the system,
// a custom solution to replace this should be implemented.
type Logger struct {
	yarf.Middleware
}

// PreDispatch wraps the http.ResponseWriter with a new LoggerWritter
// so we can log information about the response.
func (l *Logger) PreDispatch() error {
	l.Context.Response = &LoggerWriter{
		Writer: l.Context.Response,
	}

	return nil
}

func (l *Logger) PostDispatch() error {
	// If nobody sets the status code, it's a 200
	code := l.Context.Response.(*LoggerWriter).StatusCode
	if code == 0 {
		code = 200
	}

	log.Printf(
		"| %s | %s | %d | %s | %s ",
		l.GetClientIP(),
		l.Context.Request.Method,
		code,
		l.Context.Request.URL.String(),
		l.Context.Params.Encode(),
	)

	return nil
}

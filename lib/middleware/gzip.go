package middleware

import (
	"compress/gzip"
	"github.com/yarf-framework/yarf"
	"strings"
	"net/http"
)

// GzipWriter will replace (wrap) the http.ResponseWriter to gzip all content written to the response.
// It implements the http.ResponseWriter interface.
type GzipWriter struct {
    Writer http.ResponseWriter
}

// Header is a wrapper for http.ResponseWriter.Header()
func (gzw *GzipWriter) Header() http.Header {
    return gzw.Writer.Header()
}

// WriteHeader is a wrapper for http.ResponseWriter.WriteHeader()
func (gzw *GzipWriter) WriteHeader(code int) {
    gzw.Writer.WriteHeader(code)
}

// Write compress the content received and writes it to the client through the http.ResponseWriter
func (gzw *GzipWriter) Write(content []byte) (int, error) {
	// Create writer
	gz := gzip.NewWriter(gzw.Writer)
	defer gz.Close()
	
	// Set response headers
	gzw.Writer.Header().Set("Content-Type", http.DetectContentType(content))
	gzw.Writer.Header().Set("Content-Encoding", "gzip")
	
	// Write gzip bytes
	return gz.Write(content)
}

// Gzip middleware automatically handles gzip compressed responses to clients that accepts the encoding.
// It should be inserted at the beggining of the middleware stack so it can catch every write to the response and encode it right.
type Gzip struct {
	yarf.Middleware
}

// PreDispatch 
func (m *Gzip) PreDispatch() error {
    // Check request header
	if !strings.Contains(m.Context.Request.Header.Get("Accept-Encoding"), "gzip") {
		return nil
	}
	
	// Wrap response writer
	m.Context.Response = &GzipWriter {
	    Writer: m.Context.Response,
	}
	
	return nil
}

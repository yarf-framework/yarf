package middleware

import (
	"bytes"
	"compress/gzip"
	"github.com/yarf-framework/yarf"
	"strings"
)

// Gzip middleware automatically handles gzip compressed responses to clients that accepts the encoding.
type Gzip struct {
	yarf.Middleware
}

// Compress response content to Gzip and set the corresponding HTTP headers
// Will perform the compression only if the client has sent the right Accept-Encoding header.
func (m *Gzip) PostDispatch() error {
	// Check request header
	if !strings.Contains(m.Context.Request.Header.Get("Accept-Encoding"), "gzip") {
		return nil
	}

	// Encode response
	b := new(bytes.Buffer)
	gz := gzip.NewWriter(b)

	// Write gzip bytes and close
	_, err := gz.Write([]byte(m.Context.ResponseContent))
	gz.Close()
	if err != nil {
		return err
	}

	// Save result
	m.Context.ResponseContent = b.String()

	// Set response header
	m.Context.Response.Header().Set("Content-Encoding", "gzip")

	return nil
}

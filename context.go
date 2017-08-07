package yarf

import (
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
)

// ContextData interface represents a common get/set/del set of methods to handle data storage.
// Is designed to be used as the Data property of the Context obejct.
// The Data property is a free storage unit that apps using the framework can implement to their convenience
// to share context data during a request life.
// All methods returns an error status that different implementations can design to fulfill their needs.
type ContextData interface {
	// Get retrieves a data item by it's key name.
	Get(key string) (interface{}, error)

	// Set saves a data item under a key name.
	Set(key string, data interface{}) error

	// Del removes the data item and key name for a given key.
	Del(key string) error
}

// Params wraps a map[string]string and adds Get/Set/Del methods to work with it.
// Inspired on url.Values but simpler as it doesn't handles a map[string][]string
type Params map[string]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string.
func (p Params) Get(key string) string {
	if p == nil {
		return ""
	}

	param, _ := p[key]
	return param
}

// Set sets the key to value. It replaces any existing values.
func (p Params) Set(key, value string) {
	p[key] = value
}

// Del deletes the values associated with key.
func (p Params) Del(key string) {
	delete(p, key)
}

// Context is the data/status storage of every YARF request.
// Every request will instantiate a new Context object and fill in with all the request data.
// Each request Context will be shared along the entire request life to ensure accesibility of its data at all levels.
type Context struct {
	// The *http.Request object as received by the HandleFunc.
	Request *http.Request

	// The http.ResponseWriter object as received by the HandleFunc.
	Response http.ResponseWriter

	// Parameters received through URL route
	Params Params

	// Free storage to be used freely by apps to their convenience.
	Data ContextData

	// Group route storage for dispatch
	groupDispatch []Router
}

// NewContext creates a new *Context object with default values and returns it.
func NewContext(r *http.Request, rw http.ResponseWriter) *Context {
	return &Context{
		Request:  r,
		Response: rw,
		Params:   Params{},
	}
}

// Status sets the HTTP status code to be returned on the response.
func (c *Context) Status(code int) {
	c.Response.WriteHeader(code)
}

// Param is a wrapper for c.Params.Get()
func (c *Context) Param(name string) string {
	return c.Params.Get(name)
}

// FormValue is a wrapper for c.Request.Form.Get() and it calls c.Request.ParseForm().
func (c *Context) FormValue(name string) string {
	c.Request.ParseForm()

	return c.Request.Form.Get(name)
}

// QueryValue is a wrapper for c.Request.URL.Query().Get().
func (c *Context) QueryValue(name string) string {
	return c.Request.URL.Query().Get(name)
}

// GetClientIP retrieves the client IP address from the request information.
// It detects common proxy headers to return the actual client's IP and not the proxy's.
func (c *Context) GetClientIP() (ip string) {
	var pIPs string
	var pIPList []string

	if pIPs = c.Request.Header.Get("X-Real-Ip"); pIPs != "" {
		pIPList = strings.Split(pIPs, ",")
		ip = strings.TrimSpace(pIPList[0])

	} else if pIPs = c.Request.Header.Get("Real-Ip"); pIPs != "" {
		pIPList = strings.Split(pIPs, ",")
		ip = strings.TrimSpace(pIPList[0])

	} else if pIPs = c.Request.Header.Get("X-Forwarded-For"); pIPs != "" {
		pIPList = strings.Split(pIPs, ",")
		ip = strings.TrimSpace(pIPList[0])

	} else if pIPs = c.Request.Header.Get("X-Forwarded"); pIPs != "" {
		pIPList = strings.Split(pIPs, ",")
		ip = strings.TrimSpace(pIPList[0])

	} else if pIPs = c.Request.Header.Get("Forwarded-For"); pIPs != "" {
		pIPList = strings.Split(pIPs, ",")
		ip = strings.TrimSpace(pIPList[0])

	} else if pIPs = c.Request.Header.Get("Forwarded"); pIPs != "" {
		pIPList = strings.Split(pIPs, ",")
		ip = strings.TrimSpace(pIPList[0])

	} else {
		ip = c.Request.RemoteAddr
	}

	return strings.Split(ip, ":")[0]
}

// Redirect sends the corresponding HTTP redirect response with the provided URL and status code.
// It's just a wrapper for net/http.Redirect()
func (c *Context) Redirect(url string, code int) {
	http.Redirect(c.Response, c.Request, url, code)
}

// Render writes a string to the http.ResponseWriter.
// This is the default renderer that just sends the string to the client.
// Check other Render[Type] functions for different types.
func (c *Context) Render(content string) {
	// Write response
	c.Response.Write([]byte(content))
}

// RenderGzip takes a []byte content and if the client accepts compressed responses,
// writes the compressed version of the content to the response.
// Otherwise it just writes the plain []byte to it.
func (c *Context) RenderGzip(content []byte) error {
	// Check if client accepts compression
	if !strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") {
		c.Response.Write(content)
		return nil
	}

	// Detect content type
	c.Response.Header().Set("Content-Type", http.DetectContentType(content))

	// Write compressed content
	gz := gzip.NewWriter(c.Response)
	defer gz.Close()

	c.Response.Header().Set("Content-Encoding", "gzip")
	c.Response.Header().Set("Vary", "Accept-Encoding")
	c.Response.Header().Del("Content-Length")

	gz.Write(content)

	return nil
}

// RenderJSON takes a interface{} object and writes the JSON encoded string of it.
func (c *Context) RenderJSON(data interface{}) {
	// Set content
	encoded, err := json.Marshal(data)
	if err != nil {
		c.Response.Write([]byte(err.Error()))
	} else {
		c.Response.Write(encoded)
	}
}

// RenderJSONIndent is the indented (beauty) of RenderJSON
func (c *Context) RenderJSONIndent(data interface{}) {
	// Set content
	encoded, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		c.Response.Write([]byte(err.Error()))
	} else {
		c.Response.Write(encoded)
	}
}

// RenderGzipJSON takes a interface{} object and writes the JSON verion through RenderGzip.
func (c *Context) RenderGzipJSON(data interface{}) {
	// Create JSON content
	encoded, err := json.Marshal(data)
	if err != nil {
		c.Response.Write([]byte(err.Error()))
	}

	c.RenderGzip(encoded)
}

// RenderXML takes a interface{} object and writes the XML encoded string of it.
func (c *Context) RenderXML(data interface{}) {
	// Set content
	encoded, err := xml.Marshal(data)
	if err != nil {
		c.Response.Write([]byte(err.Error()))
	} else {
		c.Response.Write(encoded)
	}
}

// RenderXMLIndent is the indented (beauty) of RenderXML
func (c *Context) RenderXMLIndent(data interface{}) {
	// Set content
	encoded, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		c.Response.Write([]byte(err.Error()))
	} else {
		c.Response.Write(encoded)
	}
}

// RenderGzipXML takes a interface{} object and writes the XML verion through RenderGzip.
func (c *Context) RenderGzipXML(data interface{}) {
	// Set XML content
	encoded, err := xml.Marshal(data)
	if err != nil {
		c.Response.Write([]byte(err.Error()))
	}

	c.RenderGzip(encoded)
}

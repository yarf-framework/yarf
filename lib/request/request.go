// request package implements several utility functions to work with http.Request objects and related.
package request

import (
	"net/http"
	"strings"
)

// GetClientIP takes a HTTP Request and figures out the client IP address
// It detects common proxy headers to return the actual client's IP and not the proxy's.
//
// Params:
//  - r *http.Request
//
func GetClientIP(r *http.Request) (ip string) {
	var pIps string
	var pIpList []string

	if pIps = r.Header.Get("X-FORWARDED-FOR"); pIps != "" {
		pIpList = strings.Split(pIps, ",")
		ip = strings.TrimSpace(pIpList[0])

	} else if pIps = r.Header.Get("X-FORWARDED"); pIps != "" {
		pIpList = strings.Split(pIps, ",")
		ip = strings.TrimSpace(pIpList[0])

	} else if pIps = r.Header.Get("FORWARDED-FOR"); pIps != "" {
		pIpList = strings.Split(pIps, ",")
		ip = strings.TrimSpace(pIpList[0])

	} else if pIps = r.Header.Get("FORWARDED"); pIps != "" {
		pIpList = strings.Split(pIps, ",")
		ip = strings.TrimSpace(pIpList[0])

	} else {
		ip = r.RemoteAddr
	}

	return strings.Split(ip, ":")[0]
}

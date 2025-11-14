package main

import (
	"net"
	"net/http"
	"strings"
)

func GetIP(r *http.Request) string {
	// Check X-Forwarded-For
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return strings.Split(xff, ",")[0]
	}

	// Check X-Real-IP
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Fallback to RemoteAddr
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

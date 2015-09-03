package main

import (
	"github.com/yarf-framework/yarf"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkSimpleRequest(b *testing.B) {
	// Replicate main.go setup
	y := yarf.New()
	y.Add("/", new(Hello))

	// Create request/response
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		y.ServeHTTP(res, req)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// middleware

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := newResponseWriter(w)
		next.ServeHTTP(rw, r)
		duration := time.Since(start)
		log.Printf("[%s] %s %s | Status: %d | Duration: %v",
			r.Method, r.RemoteAddr, r.URL.Path, rw.statusCode, duration)
	})
}

// health check handler

var serverStartTime = time.Now()

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	uptime := time.Since(serverStartTime).Round(time.Second)

	payload := map[string]string{
		"status":    "OK",
		"uptime":    uptime.String(),
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"version":   "1.0.0",
	}

	json.NewEncoder(w).Encode(payload)
}

// hello handler

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 - Page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "Hello! The Go Web Server is running successfully.")

}

// form handler

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Only POST is accepted on this endpoint."}`, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"ParseForm failed: %v"}`, err), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")
	email := r.FormValue("email")

	if name == "" || email == "" {
		http.Error(w, `{"error":"Name and email are required fields."}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	payload := map[string]string{
		"status":  "success",
		"name":    name,
		"address": address,
		"email":   email,
	}
	json.NewEncoder(w).Encode(payload)
}

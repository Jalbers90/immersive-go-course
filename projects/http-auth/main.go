package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"golang.org/x/time/rate"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		params := r.URL.Query()
		data := map[string]any{
			"Params": params,
		}

		if r.Method == "GET" {
			tmpl, err := template.ParseFiles("index.html")
			if err != nil {
				http.Error(w, "Error parsing template", http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, "Error executing template", http.StatusInternalServerError)
				return
			}

		} else if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			response := []byte("<!DOCTYPE html><html>")
			response = append(response, body...)
			w.Write(response)
		}

	})

	http.HandleFunc("/200", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.Handle("/404", http.NotFoundHandler())

	http.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	http.HandleFunc("/authenticated", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Www-Authenticate", "Basic realm=\"localhost\", charset=\"UTF-8\"")

		name := os.Getenv("AUTH_USERNAME")
		pass := os.Getenv("AUTH_PASSWORD")
		comb := fmt.Sprintf("%s:%s", name, pass)
		encode := base64.StdEncoding.EncodeToString([]byte(comb))
		check := fmt.Sprintf("Basic %s", encode)

		auth := r.Header.Get("Authorization")
		if auth == check {
			w.Write([]byte("authorized\n"))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("UNauthorized\n"))
		}
	})

	limiter := rate.NewLimiter(100, 30)
	http.HandleFunc("/limited", func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			w.Write([]byte("Success"))
		} else {
			w.WriteHeader(http.StatusBadGateway)
		}
	})

	http.ListenAndServe(":8080", nil)
}

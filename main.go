package main

import (
	"h8-movies/handler"
)

type MyString string

var USERNAME = "test@mail.com"

var PASSWORD = "123456"

func main() {

	handler.StartApp()
	// http.HandleFunc("/", middleware1(middleware2(http.HandlerFunc(mainEndpoint))))

	// http.ListenAndServe(":8080", nil)
}

// func middleware2(next http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodGet {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte("invalid method"))
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	}
// }

// func middleware1(next http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		username, password, ok := r.BasicAuth()

// 		if !ok {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("missing required credential"))
// 			return
// 		}

// 		isValidCredential := (username == USERNAME) && (password == PASSWORD)

// 		if !isValidCredential {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("invalid credential"))
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	}
// }

// //type HandlerFunc func(ResponseWriter, *Request)

// func mainEndpoint(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Main Endpoint"))
// }

package main

import "net/http"

func main() {
	// 
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})
		mux.Handle("/blog", &blog{Title: "My Blog"})
		http.ListenAndServe(":8080", mux)
	}()
	
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, person!"))
	})
	mux2.Handle("/blog", &blog{Title: "My Blog"})
	http.ListenAndServe(":8081", mux2)
}

type blog struct {
	Title string
}

func (b *blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.Title))
}
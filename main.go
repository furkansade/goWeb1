package main

import (
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Main Page!"))
	x := r.URL.Path[1:]
	message := ""

	if len(x) > 0 {
		message = "Merhaba, " + x + "!"
	} else {
		message = "Main Page!"
	}

	w.Write([]byte(message))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page!"))
	w.WriteHeader(http.StatusOK)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page!"))
}

func main() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)
}

/*
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Merhaba!"))
	})
*/

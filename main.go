package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	x1 := &messageHandler{"First message!"}
	x2 := &messageHandler{"Second message!"}

	mux.Handle("/first", x1)
	mux.Handle("/second", x2)

	log.Println("Listening...")
	http.ListenAndServe(":8081", mux)

}

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, x.message)
}
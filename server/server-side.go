package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	r := http.NewServeMux()
	r.HandleFunc("/test", test)
	buildHandler := http.FileServer(http.Dir("../client/build"))
	r.Handle("/", buildHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server started on http://127.0.0.1")
	log.Fatal(srv.ListenAndServe())

}

func test(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "../client/build/index.html")
	fmt.Fprintf(w, "test\n")
}

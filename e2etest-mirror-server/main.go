package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n%s\n%s\n", os.Getenv("NODE_NAME"), os.Getenv("POD_UID"), r.RemoteAddr)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

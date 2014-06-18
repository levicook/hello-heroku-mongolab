package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/levicook/go-detect"
)

var (
	port string
)

func init() {
	port = detect.String(os.Getenv("PORT"), "5002")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, env := range os.Environ() {
			fmt.Fprintln(w, env)
		}
	})

	panicIf(http.ListenAndServe(":"+port, nil))
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

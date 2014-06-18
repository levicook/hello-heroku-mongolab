package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/levicook/go-detect"
	"labix.org/v2/mgo"
)

var (
	port    string
	session *mgo.Session
)

func init() {
	port = detect.String(os.Getenv("PORT"), "5002")

	s, err := mgo.Dial(detect.String(os.Getenv("MONGOLAB_URI"), ""))
	panicIf(err)
	session = s
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, env := range os.Environ() {
			fmt.Fprintln(w, env)
		}

		for _, serverAddr := range session.LiveServers() {
			fmt.Fprintln(w, serverAddr)
		}
	})

	panicIf(http.ListenAndServe(":"+port, nil))
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

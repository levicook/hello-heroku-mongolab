package main

import (
	"fmt"
	"net/http"
	"os"
	"labix.org/v2/mgo"
)

var (
	port    string
	session *mgo.Session
)

func init() {
	port = os.Getenv("PORT")

	s, err := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	panicIf(err)
	session = s
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

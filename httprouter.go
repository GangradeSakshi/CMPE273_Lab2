package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Message struct {
	Name string `json:"name"`
}

type MessageReturn struct {
	Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}
func greeting(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var m Message
	var r MessageReturn
	decoder := json.NewDecoder(req.Body)
	fmt.Println(req.Body)

	err := decoder.Decode(&m)
	if err != nil {
		fmt.Println("unable to debug..error")
	}

	r.Greeting = "Hello," + m.Name

	mapA, _ := json.Marshal(r)
	fmt.Fprintf(rw, string(mapA))

}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/greeting", greeting)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

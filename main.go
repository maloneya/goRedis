package main

import (
	"fmt"
	"net/http"
)

type Request struct {
	Key string
}

func (r Request) String() string {
	return fmt.Sprintln(r.Key)
}

func get(c chan Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi!")
		//todo error handling
		keys, _ := r.URL.Query()["key"]
		if len(keys) > 0 {
			c <- Request{keys[0]}
		}
	}
}

func RequestConsuemr(c chan Request) {
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	c := make(chan Request)

	go RequestConsuemr(c)
	http.HandleFunc("/get", get(c))
	http.ListenAndServe(":8090", nil)
}

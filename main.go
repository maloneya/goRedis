package main

import (
	"fmt"
	"net/http"
)

func get(c chan int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "sending")
		c <- 1
	}
}

func RequestConsuemr(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	c := make(chan int)

	go RequestConsuemr(c)
	http.HandleFunc("/get", get(c))
	http.ListenAndServe(":8090", nil)
}

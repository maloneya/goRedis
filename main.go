package main

import (
	"fmt"
	"net/http"
)

type Request struct {
	Key    string
	Result chan string
}

func (r Request) String() string {
	return fmt.Sprintln(r.Key)
}

func get(work_queue chan Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi!")
		//todo error handling
		keys, _ := r.URL.Query()["key"]
		if len(keys) > 0 {
			req := Request{keys[0], make(chan string)}
			work_queue <- req
			res := <-req.Result
			fmt.Printf("%v\n", res)
		}
	}
}

func RequestConsuemr(c <-chan Request, p Provider) {
	for request := range c {
		result := p.fetchKey(request.Key)
		request.Result <- result
	}
}

func main() {
	p := RedisProvider{RedisClientWrapper{}, LRUCache{make(map[string]string)}}
	c := make(chan Request)

	go RequestConsuemr(c, p)
	http.HandleFunc("/get", get(c))
	http.ListenAndServe(":8090", nil)
}

package main

import (
	"fmt"
	"net/http"
)

func (provider RedisProvider) get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi!")
	//todo error handling
	keys, _ := r.URL.Query()["key"]
	if len(keys) > 0 {
		res := provider.fetchKey(keys[0])
		fmt.Printf("%v\n", res)
	}
}

func main() {
	p := RedisProvider{RedisClientWrapper{}, LRUCache{make(map[string]string)}}

	http.HandleFunc("/get", p.get)
	http.ListenAndServe(":8090", nil)
}

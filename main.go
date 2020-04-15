package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fmt.Sprintf("hello : %d", time.Now().Unix())))
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

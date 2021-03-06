package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)

		if err!=nil{
			http.Error(rw,"Oops",http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops"))
			return
		}
		fmt.Fprintf(rw,"Hello %s",d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Good Bye World")
	})
	http.ListenAndServe(":8080", nil)
}

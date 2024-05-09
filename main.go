package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func (rw http.ResponseWriter, r *http.Request)  {
		d, err := io.ReadAll(r.Body)
		if err!=nil{
			http.Error(rw,"Oops!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw,"Hello %s", d)
	})

	http.HandleFunc("/greet", func(http.ResponseWriter, *http.Request) {
		log.Println("Good evening!")
	})
	
	http.ListenAndServe(":9090",nil)
	
}
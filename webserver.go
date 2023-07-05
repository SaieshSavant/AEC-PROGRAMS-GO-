package main

import(
	
	"fmt"
	"log"
	"net/http"
)

func main(){
	handler:=func(w http.ResponseWriter,r*http.Request){
		fmt.Fprintf(w,"Hello world!")

	}

	mux:=http.NewServeMux()
	mux.HandleFunc("/",handler)

	server:=&http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	log.Println("Server listening on port 8080...")
	log.Fatal(server.ListenAndServe())
}

//2023/07/05 10:44:50 Server listening on port 8080...
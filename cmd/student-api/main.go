package main 

import (
	"fmt"
	"github.com/CodeMaverick-143/Golang_learning/internal/config"
	"net/http"
	"log"
)

func main(){
	// load config
	cfg := config.MustLoad()




	//database setup 



	// setup router

	router:= http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("welcome to student api"))
	})



	// setup server

	server:= http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
    fmt.Println("Server started on", cfg.HTTPServer.Address)
	err:= server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}


}


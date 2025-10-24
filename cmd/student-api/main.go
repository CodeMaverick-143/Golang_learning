package main 

import (
	"fmt"
	"github.com/CodeMaverick-143/Golang_learning/internal/config"
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"
	"log/slog"
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

	done := make(chan os.signal,1)

	signal.Notify(done , os.Interrupt,syscall.SIGTERM,syscall.SIGINT)

	go func(){
		err:= server.ListenAndServe()
		if err != nil{
			log.Fatal(err)
		}
	
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx); err != nil{
		slog.Error("Failed to shutdown server",slog.String("error",err.Error()))
	}

	slog.Info("Server shutdown successfully")

	


}


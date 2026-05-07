package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//create a new server with handler
	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello World")
		}),
	}

	//run server in goroutine
	go func() {
		fmt.Println("Server is running on PORT :8080")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	//create channel for os signals, os signal is like sigint and sigterm
	quit := make(chan os.Signal, 1)

	//notify quit channel on receiving os signals
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	//block until a signal is received
	<-quit
	fmt.Println("Shutting Down Server...")

	//create a new context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	//shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}

	fmt.Println("Server Exited Properly")

}

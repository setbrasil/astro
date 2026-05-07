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

const usage = `Astro - Webserver CLI

Usage:
	astro startserver   Start a local webserver
	astro help          Show Help
`

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(usage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help", "--help", "-h":
		fmt.Printf(usage)

	case "startserver":
		startserver()
	}
}

func startserver() {
	fmt.Println("🚀 Server started in http://localhost:8080")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Astro — by SET!")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// inicia o servidor
	go func() {
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			fmt.Println("erro:", err)
		}
	}()

	
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	fmt.Println("\n🛑 Stopping server...")

	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("shutdown error:", err)
	}

	fmt.Println("✅ Server stopped")
}

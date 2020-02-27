package main

import (
	"context"
	"flag"
	"os/signal"
	"time"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	api "{{cookiecutter.app_name}}/api"
	utl "{{cookiecutter.app_name}}/utils"
	customLogger "{{cookiecutter.app_name}}/log"
)



func main() {
	// remove settings from standard logger
	log.SetFlags(0)
	log.SetOutput(new(customLogger.LogWriter))

	var hostPort string
	// Add all CLI arguments after this comment and before flag.Parse()

	hostPortCLI := flag.String("hostport", "127.0.0.1:8800", "Host:port for start the server!")
	flag.Parse()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	// routers here
	router := mux.NewRouter()
	router.HandleFunc("/", api.PingPongHandler).Methods("GET").Name("pingpong")

	router.Use(customLogger.LogMiddleware)

	hpGlVar := os.Getenv("SERVER")

	if hpGlVar != "" {
		hostPort = hpGlVar
	} else {
		hostPort = *hostPortCLI
	}

	log.Printf("Servers started at: %v", hostPort)
	// http.HandleFunc("/", HelloHandler)

	srv := &http.Server{
		Addr: hostPort,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	{% if cookiecutter.use_mongo == "y" and cookiecutter.use_postgres !="n" -%}
	//defer <mongo connection here>
	defer utl.SQLDB.Close()
	{% elif cookiecutter.use_mongo != "y" and cookiecutter.use_postgres !="n" -%}
	defer utl.SQLDB.Close()
	{% else -%}
	// DB correction not set
	{%- endif %}
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

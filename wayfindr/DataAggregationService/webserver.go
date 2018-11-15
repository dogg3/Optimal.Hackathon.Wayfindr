package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// WebServer initiates a HTTP webserver providing a RESTful API
//
// Any unrouted request will return 404 error
func WebServer() {
	r := mux.NewRouter()
	r.HandleFunc("/HandleCameraService/{timestamp}", HandleCameraService).Methods("POST")
	r.HandleFunc("/HandleBeaconService/{timestamp}", HandleBeaconService).Methods("POST")

	log.Println("Launching webserver at address: ", (viper.GetString("host.domain") + viper.GetString("host.port")))
	srv := &http.Server{
		Addr:         (viper.GetString("host.domain") + viper.GetString("host.port")),
		WriteTimeout: time.Second * viper.GetDuration("host.timeout.write"),
		ReadTimeout:  time.Second * viper.GetDuration("host.timeout.read"),
		IdleTimeout:  time.Second * viper.GetDuration("host.timeout.idle"),
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("webserver.timeout.cancel"))
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Gracefully shutting down")
	os.Exit(0)
}

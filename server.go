package main

import (
	"log"
	"net"
	"net/http"
	"origin-server/rozay.net/handlers"
	"time"
)

func SimpleHttpServer() {
	srv := &http.Server{
		Addr: "127.0.0.1:8081",
		Handler: http.TimeoutHandler(
			handlers.DefaultHandler(), 2*time.Minute, ""),
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := srv.Serve(l)
		if err != nil {
			log.Fatal(err)
		}
	}()

}

package www

import (
	"api_standard/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	port string
}

func ServerConfig(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Listener() {
	ctx := context.Background()

	server_channel := make(chan os.Signal, 1)
	signal.Notify(server_channel, os.Interrupt, syscall.SIGTERM)

	initServer := initServer(s.port)

	go initServer.ListenAndServe()

	log.Println("Server Start")

	<-server_channel

	initServer.Shutdown(ctx)

	log.Println("Server Stopped")
}

func initServer(addr string) *http.Server {
	router.Routes()

	return &http.Server{
		Addr: ":" + addr,
	}
}

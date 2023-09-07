package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"try-go-grpc/menu"

	pb "try-go-grpc/menu/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	menuRepository := menu.NewRepository()
	menuServer := menu.NewMenuServer(menuRepository)

	pb.RegisterMenuServer(server, menuServer)

	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

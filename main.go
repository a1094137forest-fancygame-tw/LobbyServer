package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"LobbyServer/config"
	"LobbyServer/constant"
	v1 "LobbyServer/controllers/v1"
	"LobbyServer/proto/backendLobby"
)

func init() {
	constant.ReadConfig(".env")
}
func main() {
	log.Println("[info] Create server:", config.LocalServerUrl)
	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	backendLobby.RegisterLobbyServer(s, &v1.LobbyServe{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

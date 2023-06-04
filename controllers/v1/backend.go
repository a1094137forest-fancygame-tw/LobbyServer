package v1

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"LobbyServer/config"
	"LobbyServer/proto/backendLobby"
	"LobbyServer/proto/user"
)

type LobbyServe struct {
	backendLobby.UnimplementedLobbyServer
}

func (s *LobbyServe) GetMemberList(ctx context.Context, req *backendLobby.EmptyReq) (*backendLobby.GetMemberListRes, error) {
	log.Println("lobby server get getMemberList Req")
	var resp backendLobby.GetMemberListRes

	conn, err := grpc.Dial(config.UserServerUrl, grpc.WithInsecure())
	if err != nil {
		log.Println("conn failed")
		return nil, err
	}
	defer conn.Close()

	userRes, err := user.NewUserClient(conn).GetMemberList(ctx, &user.EmptyReq{})
	if err != nil {
		log.Println("userRes failed")
		return nil, err
	}

	byteData, _ := proto.Marshal(userRes)
	_ = proto.Unmarshal(byteData, &resp)
	log.Println("getMemberList req success")
	return &resp, nil
}

func (s *LobbyServe) SetMemberData(ctx context.Context, req *backendLobby.SetMember) (*backendLobby.SetMemberRes, error) {
	var resp backendLobby.SetMemberRes

	byteData, _ := proto.Marshal(req)

	var userReq user.SetMember
	_ = proto.Unmarshal(byteData, &userReq)

	conn, err := grpc.Dial(config.UserServerUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	userRes, err := user.NewUserClient(conn).SetMemberData(ctx, &userReq)
	if err != nil {
		return nil, err
	}

	byteData, _ = proto.Marshal(userRes)
	_ = proto.Unmarshal(byteData, &resp)

	return &resp, nil
}

func (s *LobbyServe) KickOutMember(ctx context.Context, req *backendLobby.KickOutMemberInfo) (*backendLobby.KickOutMemberRes, error) {
	var resp backendLobby.KickOutMemberRes

	byteData, _ := proto.Marshal(req)

	var userReq user.KickOutMemberInfo
	_ = proto.Unmarshal(byteData, &userReq)

	conn, err := grpc.Dial(config.UserServerUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	userRes, err := user.NewUserClient(conn).KickOutMember(ctx, &userReq)
	if err != nil {
		return nil, err
	}

	byteData, _ = proto.Marshal(userRes)
	_ = proto.Unmarshal(byteData, &resp)

	return &resp, nil
}

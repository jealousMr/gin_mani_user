package main

import (
	"context"
	"gin_mani_user/conf"
	"gin_mani_user/handler"
	pb_mani "gin_mani_user/pb"
	logx "github.com/amoghe/distillog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)


type S struct {
}

func (s S) AddAndUpdateUserInfo(ctx context.Context, req *pb_mani.AddAndUpdateUserInfoReq) (*pb_mani.AddAndUpdateUserInfoResp, error) {
	logx.Infoln("AddAndUpdateUserInfo req: %v", req)
	resp, err := handler.AddAndUpdateUserInfo(ctx, req)
	logx.Infoln("AddAndUpdateUserInfo resp: %v", resp)
	return resp, err
}

func main() {
	logx.Infof("start mani user server")
	lis, err := net.Listen("tcp", conf.ServerPort)
	if err != nil {
		log.Fatal("failed to listen")
	}
	server := grpc.NewServer()
	pb_mani.RegisterGinUserServiceServer(server, &S{})
	reflection.Register(server)
	logx.Infof("run mani user server success...")
	if err := server.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}

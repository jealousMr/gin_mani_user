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

func (s S) AddUserRecord(ctx context.Context, req *pb_mani.AddUserRecordReq) (*pb_mani.AddUserRecordResp, error) {
	logx.Infoln("AddUserRecord req: %v", req)
	resp, err := handler.AddUserRecord(ctx, req)
	logx.Infoln("AddUserRecord resp: %v", resp)
	return resp, err
}

func (s S) QueryUserRecord(ctx context.Context, req *pb_mani.QueryUserRecordReq) (*pb_mani.QueryUserRecordResp, error) {
	logx.Infoln("QueryUserRecord req: %v", req)
	resp, err := handler.QueryUserRecord(ctx, req)
	logx.Infoln("QueryUserRecord resp: %v", resp)
	return resp, err
}

func (s S) QueryUserInfoByIds(ctx context.Context, req *pb_mani.QueryUserInfoByIdsReq) (*pb_mani.QueryUserInfoByIdsResp, error) {
	logx.Infoln("QueryUserInfoByIds req: %v", req)
	resp, err := handler.QueryUserInfoByIds(ctx, req)
	logx.Infoln("QueryUserInfoByIds resp: %v", resp)
	return resp, err
}

func (s S) AddAndUpdateUserInfo(ctx context.Context, req *pb_mani.AddAndUpdateUserInfoReq) (*pb_mani.AddAndUpdateUserInfoResp, error) {
	logx.Infoln("AddAndUpdateUserInfo req: %v", req)
	resp, err := handler.AddAndUpdateUserInfo(ctx, req)
	logx.Infoln("AddAndUpdateUserInfo resp: %v", resp)
	return resp, err
}

func main() {
	cf := conf.GetConfig()
	logx.Infof("start mani user server")
	lis, err := net.Listen("tcp", cf.Server.Port)
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

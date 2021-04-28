package test

import (
	"context"
	"fmt"
	"gin_mani_user/handler"
	pb_mani "gin_mani_user/pb"
	"testing"
)

func TestAddUser(t *testing.T) {
	req := &pb_mani.AddAndUpdateUserInfoReq{
		UserInfo: &pb_mani.UserInfo{
			UserId:    "test uid",
			NickName:  "nick",
			AvatarUrl: "ava2",
			Gender:    1,
			Country:   "ar",
			UserState: 1,
		},
	}
	resp, err := handler.AddAndUpdateUserInfo(context.Background(), req)
	fmt.Println(resp, err)
}

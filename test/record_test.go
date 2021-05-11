package test

import (
	"context"
	"fmt"
	"gin_mani_user/handler"
	pb_mani "gin_mani_user/pb"
	"testing"
)

func TestAddRecord(t *testing.T) {
	req := &pb_mani.AddUserRecordReq{
		UserRecord: &pb_mani.UserRecord{
			UserId: "test user",
			RuleId: "cascxasc",
			Extra:  "xxx",
		},
	}
	resp, err := handler.AddUserRecord(context.Background(), req)
	fmt.Println(resp, err)
}

func TestQueryRecord(t *testing.T) {
	req := &pb_mani.QueryUserRecordReq{
		UserId:   "test user",
		PageNo:   1,
		PageSize: 10,
	}
	resp, err := handler.QueryUserRecord(context.Background(), req)
	fmt.Println(resp, err)
}

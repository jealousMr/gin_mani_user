package handler

import (
	"context"
	"errors"
	"gin_mani_user/logic"
	pb_mani "gin_mani_user/pb"
	"gin_mani_user/util"
	logx "github.com/amoghe/distillog"
)

func checkAddUserRecordParam(req *pb_mani.AddUserRecordReq) error {
	if req.UserRecord == nil {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func AddUserRecord(ctx context.Context, req *pb_mani.AddUserRecordReq) (resp *pb_mani.AddUserRecordResp, err error) {
	resp = &pb_mani.AddUserRecordResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkAddUserRecordParam(req); err != nil {
		logx.Errorf("AddUserRecord checkAddUserRecordParam error:%v", err)
		return nil, err
	}
	if err = logic.AddUserRecord(ctx, req.UserRecord); err != nil {
		logx.Errorf("AddUserRecord error:%v", err)
		return nil, err
	}
	return
}

func checkQueryUserRecordParam(req *pb_mani.QueryUserRecordReq) error {
	if req.RuleId == "" && req.UserId == "" {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func QueryUserRecord(ctx context.Context, req *pb_mani.QueryUserRecordReq) (resp *pb_mani.QueryUserRecordResp, err error) {
	resp = &pb_mani.QueryUserRecordResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkQueryUserRecordParam(req); err != nil {
		logx.Errorf("QueryUserRecord checkQueryUserRecordParam error:%v", err)
		return nil, err
	}
	resp, err = logic.QueryUserRecord(ctx, req.RuleId, req.UserId, req.PageNo, req.PageSize)
	if err != nil {
		logx.Errorf("QueryUserRecord error:%v", err)
		return nil, err
	}
	return
}

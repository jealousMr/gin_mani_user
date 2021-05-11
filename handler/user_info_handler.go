package handler

import (
	"context"
	"gin_mani_user/logic"
	pb_mani "gin_mani_user/pb"
	"gin_mani_user/util"
	logx "github.com/amoghe/distillog"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
)

func checkAddAndUpdateUserInfoParam(req *pb_mani.AddAndUpdateUserInfoReq) error {
	if req.UserInfo == nil || req.UserInfo.UserId == "" {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func AddAndUpdateUserInfo(ctx context.Context, req *pb_mani.AddAndUpdateUserInfoReq) (resp *pb_mani.AddAndUpdateUserInfoResp, err error) {
	resp = &pb_mani.AddAndUpdateUserInfoResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkAddAndUpdateUserInfoParam(req); err != nil {
		logx.Errorf("AddAndUpdateUserInfo checkAddAndUpdateUserInfoParam error")
		return resp, err
	}
	if err = logic.AddAndUpdateUserInfo(ctx, req.UserInfo); err != nil {
		logx.Errorf("AddAndUpdateUserInfo logic AddAndUpdateUserInfo error")
		return resp, err
	}
	return resp, nil
}

func checkQueryUserInfoByIds(req *pb_mani.QueryUserInfoByIdsReq) error {
	if req.IdList == nil || len(req.IdList) == 0 {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func QueryUserInfoByIds(ctx context.Context, req *pb_mani.QueryUserInfoByIdsReq) (resp *pb_mani.QueryUserInfoByIdsResp, err error) {
	resp = &pb_mani.QueryUserInfoByIdsResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkQueryUserInfoByIds(req); err != nil {
		logx.Errorf("QueryUserInfoByIds checkQueryUserInfoByIds error")
		return resp, err
	}
	users, err := logic.GetUserByIds(ctx, req.IdList)
	if err != nil {
		logx.Errorf("QueryUserInfoByIds error:%v", err)
		return resp, err
	}
	resp.Users = users
	return
}

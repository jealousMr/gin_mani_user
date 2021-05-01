package logic

import (
	"context"
	"gin_mani_user/dal"
	"gin_mani_user/model"
	pb_mani "gin_mani_user/pb"
	logx "github.com/amoghe/distillog"
)

func AddAndUpdateUserInfo(ctx context.Context, user *pb_mani.UserInfo) error {
	userModel := &model.UserInfoModel{
		UserId:    user.UserId,
		NickName:  user.NickName,
		AvatarUrl: user.AvatarUrl,
		Gender:    user.Gender,
		Country:   user.Country,
		UserState: int64(user.UserState),
	}
	if err := dal.AddAndUpdateUserInfo(ctx, userModel); err != nil {
		logx.Errorf("logic AddAndUpdateUserInfo error:%v", err)
		return err
	}
	return nil
}

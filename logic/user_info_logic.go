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

func GetUserByIds(ctx context.Context, ids []string) ([]*pb_mani.UserInfo, error) {
	users, err := dal.GetUserByIds(ctx, ids)
	if err != nil {
		logx.Errorf("logic GetUserByIds err:%s", err)
		return nil, err
	}
	userList := make([]*pb_mani.UserInfo, 0)
	for _, u := range users {
		userList = append(userList, &pb_mani.UserInfo{
			UserId:    u.UserId,
			NickName:  u.NickName,
			AvatarUrl: u.AvatarUrl,
			Gender:    u.Gender,
			Country:   u.Country,
			UserState: pb_mani.UserState(u.UserState),
		})
	}
	return userList, nil
}

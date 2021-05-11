package dal

import (
	"context"
	"gin_mani_user/model"
	logx "github.com/amoghe/distillog"
	"gorm.io/gorm/clause"
)

func AddAndUpdateUserInfo(ctx context.Context, user *model.UserInfoModel) error {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("dal get db error:%v", err)
		return err
	}
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"nick_name", "avatar_url", "gender", "country", "user_state", "create_at", "update_at"}),
	})
	if err := db.Create(user).Error; err != nil {
		logx.Errorf("dal AddAndUpdateUserInfo error:%v", err)
		return err
	}
	return nil
}

func GetUserByIds(ctx context.Context, ids []string) ([]*model.UserInfoModel, error) {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("dal get db error:%v", err)
		return nil, err
	}
	userList := make([]*model.UserInfoModel, 0)
	if err := db.Table(model.UserInfoTableName()).Where("user_id in (?)", ids).Find(&userList).Error; err != nil {
		logx.Errorf("dal GetUserByIds error:%v", err)
		return nil, err
	}
	return userList, nil
}

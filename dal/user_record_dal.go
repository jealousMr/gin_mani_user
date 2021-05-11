package dal

import (
	"context"
	"gin_mani_user/model"
	logx "github.com/amoghe/distillog"
)

func AddUserRecord(ctx context.Context, record *model.UserRecordModel) error {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("dal get db error:%v", err)
		return err
	}
	if err := db.Table(model.UserRecordTableName()).Create(&record).Error; err != nil {
		logx.Errorf("dal AddUserRecord error:%v", err)
		return err
	}
	return nil

}

func GetUserRecord(ctx context.Context, ruleId, userId string, limit, offset int64) (records []*model.UserRecordModel, total int64, err error) {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("dal get db error:%v", err)
		return nil, 0, err
	}
	records = make([]*model.UserRecordModel, 0)
	db = db.Table(model.UserRecordTableName())
	if ruleId != "" {
		db = db.Where("rule_id in (?)", ruleId)
	}
	if userId != "" {
		db = db.Where("user_id in (?)", userId)
	}
	if err = db.Offset(int(offset)).Limit(int(limit)).Find(&records).Error; err != nil {
		logx.Errorf("dal GetUserRecord error:%v", err)
		return nil, 0, err
	}
	if err = db.Count(&total).Error; err != nil {
		logx.Errorf("dal GetUserRecord count error:%v", err)
		return nil, 0, err
	}
	return
}

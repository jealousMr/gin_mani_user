package model

import "time"

type UserRecordModel struct {
	OpenId     string    `json:"open_id"`
	RuleId     string    `json:"rule_id"`
	Extra      string    `json:"action_type"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func UserRecordTableName() string {
	return "user_record"
}
func (UserRecordModel) TableName() string {
	return UserRecordTableName()
}

package model

import "time"

type UserInfoModel struct {
	UserId string `json:"user_id"`
	NickName string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Gender int64 `json:"gender"`
	Country string `json:"country"`
	UserState int64 `json:"user_state"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func UserInfoTableName() string{
	return "user_info"
}

func (UserInfoModel) TableName() string{
	return UserInfoTableName()
}

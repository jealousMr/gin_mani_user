package model

type UserInfoModel struct {
	UserId string `json:"user_id"`
	NickName string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Gender int64 `json:"gender"`
	Country string `json:"country"`
	UserState int64 `json:"user_state"`
}

func UserInfoTableName() string{
	return "user_info"
}

func (UserInfoModel) TableName() string{
	return UserInfoTableName()
}

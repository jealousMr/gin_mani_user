package util

import pb_mani "gin_mani_user/pb"

func BuildBaseResp(err error,errMsg string) *pb_mani.BaseResp{
	if err == nil{
		return &pb_mani.BaseResp{
			State: StateSuccess,
			Msg: MsgSuccess,
		}
	}else{
		if errMsg != ""{
			return &pb_mani.BaseResp{
				State: StateError,
				Msg: errMsg,
			}
		}else{
			return &pb_mani.BaseResp{
				State: StateError,
				Msg: err.Error(),
			}
		}
	}
}

func GetLimitAndOffset(pageNo int64, pageSize int64) (int64, int64) {
	if pageNo < 1 || pageSize == 0 {
		return 0, 0
	}
	offset := (pageNo - 1) * pageSize
	limit := pageSize
	return limit, offset
}

func HasMore(realCount int64, totalCount int64, offset int64) bool {
	if realCount <= 0 || totalCount <= 0 || offset < 0 {
		return false
	}
	if totalCount < offset {
		offset = totalCount
	}
	if realCount < (totalCount - offset) {
		return true
	}
	return false
}

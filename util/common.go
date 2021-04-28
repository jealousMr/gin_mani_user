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

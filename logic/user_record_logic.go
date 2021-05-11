package logic

import (
	"context"
	"fmt"
	"gin_mani_user/clients"
	"gin_mani_user/dal"
	"gin_mani_user/model"
	pb_mani "gin_mani_user/pb"
	"gin_mani_user/util"
	logx "github.com/amoghe/distillog"
)

func AddUserRecord(ctx context.Context, user *pb_mani.UserRecord) error {
	center_rpc, err := clients.GetCenterClient()
	if err != nil {
		logx.Errorf("GetCenterClient error :%v", err)
		return err
	}
	ruleResp, err := center_rpc.GetRuleByCondition(ctx, &pb_mani.GetRuleByConditionReq{
		RuleId: user.RuleId,
	})
	if err != nil {
		logx.Errorf("AddUserRecord GetRuleByCondition error :%v", err)
		return err
	}
	if _, ok := ruleResp.Rules[user.RuleId]; !ok {
		logx.Errorf("AddUserRecord no rule")
		return nil
	}
	rule := ruleResp.Rules[user.RuleId]
	record := &model.UserRecordModel{
		OpenId: user.UserId,
		RuleId: user.RuleId,
		Extra:  fmt.Sprintf("%s-%s-%d", rule.User, rule.RuleId, rule.RuleType),
	}
	if err := dal.AddUserRecord(ctx, record); err != nil {
		logx.Errorf("logic AddUserRecord error:%v", err)
		return err
	}
	return nil
}

func QueryUserRecord(ctx context.Context, ruleId, userId string, pageNo, pageSize int64) (resp *pb_mani.QueryUserRecordResp, err error) {
	resp = &pb_mani.QueryUserRecordResp{}
	limit, offset := util.GetLimitAndOffset(pageNo, pageSize)
	records, total, err := dal.GetUserRecord(ctx, ruleId, userId, limit, offset)
	if err != nil {
		logx.Errorf("logic QueryUserRecord error:%v", err)
		return
	}
	recordList := make([]*pb_mani.UserRecord, 0)
	for _, r := range records {
		recordList = append(recordList, &pb_mani.UserRecord{
			UserId: r.OpenId,
			RuleId: r.RuleId,
			Extra:  r.Extra,
		})
	}
	resp.UserRecordList = recordList
	resp.Page = &pb_mani.PageStruct{
		PageSize: pageSize,
		PageNo:   pageNo,
		HasMore:  util.HasMore(int64(len(recordList)), int64(total), int64(offset)),
	}
	return
}

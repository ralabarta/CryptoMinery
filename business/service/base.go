package service

import (
	"github.com/IEatLemons/BaseGin/business/dao"
	resp "github.com/IEatLemons/GoHelper/helper/responses"
)

type UserService struct {
	D    *dao.BaseDao
	resp *resp.Resp
}

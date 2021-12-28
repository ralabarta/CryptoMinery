package model

import "time"

type Registered struct {
	RegType     string    `json:"registered_type" xorm:"-" binding:"required"`
	Username    string    `json:"username" xorm:"unique 'username'" binding:"required"`
	Password    string    `json:"password" xorm:"'password'" binding:"required"`
	CountryCode string    `json:"country_code" xorm:"'country_code'"`
	Phone       uint64    `json:"phone" xorm:"'phone'"`
	Email       string    `json:"email" xorm:"'email'"`
	InviteCode  string    `json:"invite_code" xorm:"-"`
	RegPlatform string    `json:"reg_platform" xorm:"'reg_platform'"`
	CreateTime  time.Time `json:"create_time" xorm:"'create_time' created"`
}

func (t *Registered) TableName() string {
	return "st_user"
}

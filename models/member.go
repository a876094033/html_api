package models

import "github.com/jinzhu/gorm"

type MemberInfo struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RealName     string `json:"real_name"`
	FirstCardNum string `json:"first_card_num"`
	LastCardNum  string `json:"last_card_num"`
}

func (MemberInfo) TableName() string {
	return "members"
}

func GetMemberInfo(id int) (*MemberInfo, error)  {
	var memberInfo MemberInfo
	err := db.Debug().Model(&MemberInfo{}).Where("id = ?", id).First(&memberInfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &memberInfo, nil
}

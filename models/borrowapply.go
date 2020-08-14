package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BorrowApply struct {
	Model

	Name string `json:"name"`                            //名称
	Phone string `json:"phone"`                          //手机号
	Birth string `json:"birth"`
	Sex int `json:"sex"`
	Email string `json:"email"`
	ApplyType int `json:"apply_type"`                    //借款类型
	PropertyType int `json:"property_type"`              //资产种类
	SeniorAmount float64 `json:"senior_amount"`          //市场价格
	ApplyAmount float64 `json:"apply_amount"`            //申请金额
	Period int `json:"period"`                           //借款期限
	CaseNumber string `json:"case_number"`               //房产证编号
	Postcode string `json:"postcode"`
	Address string `json:"address"`
	AddressDetail string `json:"address_detail"`

}

func GetBorrowApply(pageNum int, pageSize int, maps interface {}) (BorrowApply []BorrowApply) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&BorrowApply)

	return
}

func GetBorrowApplyTotal(maps interface {}) (count int){
	db.Model(&BorrowApply{}).Where(maps).Count(&count)

	return
}

func AddBorrowApply(data map[string]interface{}) error {
	borrow := BorrowApply{
		Name:           data["name"].(string),
		Phone:          data["phone"].(string),
		Birth :         data["birth"].(string),
		Sex :           data["sex"].(int),
		Email :         data["email"].(string),
		ApplyType :     data["apply_type"].(int),
		PropertyType :  data["property_type"].(int),
		SeniorAmount :  data["senior_amount"].(float64),
		ApplyAmount :   data["apply_amount"].(float64),
		Period :        data["period"].(int),
		CaseNumber :    data["case_number"].(string),
		Postcode :      data["postcode"].(string),
		Address :       data["address"].(string),
		AddressDetail : data["address_detail"].(string),
	}
	if err := db.Create(&borrow).Error; err != nil {
		return err
	}
	return nil
}

func (borrowapply *BorrowApply)BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

func (borrowapply *BorrowApply) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())

	return nil
}

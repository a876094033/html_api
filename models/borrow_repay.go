package models

import "github.com/jinzhu/gorm"

type BorrowRepay struct {
	ID            int     `json:"id" gorm:"primary_key"`
	BorrowId      int     `json:"borrow_name"`
	Period        int     `json:"period"`
	Capital       float64 `json:"capital"`
	Interest      float64 `json:"interest"`
	RepayStatus   int     `json:"repay_status"`
	RepayTime     string  `json:"repay_time"`
}

func AddBorrowRepay(data map[string]interface{}) error {
	borrowRepay := BorrowRepay{
		BorrowId:      data["borrow_id"].(int),
		Period:        data["period"].(int),
		Capital:       data["capital"].(float64),
		Interest:      data["interest"].(float64),
		RepayStatus:   data["repay_status"].(int),
		RepayTime:     data["repay_time"].(string),
	}
	if err := db.Create(&borrowRepay).Error; err != nil {
		return err
	}
	return nil
}

func GetBorrowRepays(id int) ([]*BorrowRepay, error) {
	var borrowRepay []*BorrowRepay
	err := db.Debug().Where("borrow_id = ?", id).Find(&borrowRepay).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return borrowRepay, nil
}

package models

import "github.com/jinzhu/gorm"

type BorrowSummary struct {
	ID                 int    `json:"id" gorm:"primary_key"`
	BorrowId           int    `json:"borrow_id"`
	Address            string `json:"address"`
	MerchantName       string `json:"merchant_name"`
	BusinessScale      string `json:"business_scale"`
	FundApply          string `json:"fund_apply"`
	ApplicationAmount  string `json:"application_amount"`
	Lcd                string `json:"lcd"`
	RedeemRanking      string `json:"redeem_ranking"`
	ConstructionPeriod string `json:"construction_period"`
	ExchangeRate       string `json:"exchange_rate"`
	Remark             string `json:"remark"`
}

func GetBorrowSummary(borrow_id int) (*BorrowSummary, error) {
	var borrowSummary BorrowSummary
	err := db.Debug().Where("borrow_id = ?", borrow_id).First(&borrowSummary).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &borrowSummary, nil
}

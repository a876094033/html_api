package models

import "github.com/jinzhu/gorm"

type Borrow struct {
	ID              int     `json:"id" gorm:"primary_key"`
	BorrowName      string  `json:"borrow_name"`
	InterestRate    float64 `json:"interest_rate"`
	Term            int     `json:"term"`
	TermType        int     `json:"term_type"`
	Amount          float64 `json:"amount"`
	RepayType       int     `json:"repay_type"`
	Purpose         string  `json:"purpose"`
	RepayName       string  `json:"repay_name"`
	Diya            string  `json:"diya"`
	AmountLimit     float64 `json:"amount_limit"`
	BorrowStatus    int     `json:"borrow_status"`
	Ramerk          string  `json:"ramerk"`
	BorrowImg       string  `json:"borrow_img"`
	BorrowStructImg string  `json:"borrow_struct_img"`
	ProvePdf        string  `json:"prove_pdf"`
}
type BorrowAll struct {
	Borrow
	InvestCount int
	InvestSum   float64
}

func GetBorrows(pageNum int, pageSize int, maps interface{}) ([]*Borrow, error) {
	var borrows []*Borrow
	err := db.Debug().Model(&Borrow{}).Where(maps).Where("borrow_status<>?", 1).Offset(pageNum / 10 * pageSize).Limit(pageSize).Find(&borrows).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return borrows, nil
}

func (Borrow) TableName() string {
	return "borrow"
}

func GetBorrowTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Debug().Model(&Borrow{}).Where(maps).Where("borrow_status <> ?", 1).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetBorrow(id int) (*Borrow, error) {
	var borrow Borrow
	err := db.Debug().Where("id = ?", id).First(&borrow).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &borrow, nil
}

func Update(id int, AmountLimit float64) error {
	if err := db.Debug().Raw("Update borrow set amount_limit = ? where id = ?", AmountLimit, id).Error; err != nil {
		return err
	}
	return nil
}

func GetBorrowsRepay(maps interface{}) ([]*Borrow, error) {
	var borrows []*Borrow
	err := db.Debug().Model(&Borrow{}).Where(maps).Find(&borrows).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return borrows, nil
}

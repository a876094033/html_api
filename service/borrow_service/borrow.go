package borrow_service

import (
	"encoding/json"
	"html_api/models"
	"html_api/pkg/gredis"
	"html_api/pkg/logging"
	"html_api/service/cache_service"
)

type Borrow struct {
	ID           int
	BorrowName   string
	InterestRate float64
	Term         int
	TermType     int
	Amount       float64
	RepayType    int
	purpose      string
	RepayName    string
	Diya         string
	AmountLimit  float64
	BorrowStatus int

	PageNum  int
	PageSize int
}

func (b *Borrow) GetBorrows() ([]*models.Borrow, error) {
	var (
		borrow, cacheBorrow []*models.Borrow
	)
	cache := cache_service.Borrow{
		PageNum:      b.PageNum,
		PageSize:     b.PageSize,
		BorrowStatus: b.BorrowStatus,
	}
	key := cache.GetBorrowsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheBorrow)
			return cacheBorrow, nil
		}
	}

	borrow, err := models.GetBorrows(b.PageNum, b.PageSize, b.getMaps())
	if err != nil {
		return nil, err
	}

	//gredis.Set(key, articles, 3600)
	return borrow, nil

}
func (b *Borrow) Count() (int, error) {
	return models.GetBorrowTotal(b.getMaps())
}
func (b *Borrow) GetBorrow() (*models.Borrow, error) {
	return models.GetBorrow(b.ID)
}
func (b *Borrow) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if b.BorrowStatus > 0 {
		maps["borrow_status"] = b.BorrowStatus
	}
	//if b.AmountLimit >= 0 {
	//	maps["amount_limit"] = b.AmountLimit
	//}
	return maps
}

func (b *Borrow) Update() error {
	return models.Update(b.ID, b.AmountLimit)
}
func (b *Borrow) GetBorrowsRepay() ([]*models.Borrow, error) {
	var (
		borrow []*models.Borrow
	)
	borrow, err := models.GetBorrowsRepay(b.getMaps())
	if err != nil {
		return nil, err
	}
	return borrow, nil

}

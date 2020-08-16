package repay_service

import (
	"html_api/models"
)

type BorrowRepay struct {
	BorrowId      int
	Period        int
	Capital       float64
	Interest      float64
	RepayStatus   int
	RepayTime     string
	RepayTimeTrue string
}

func (b *BorrowRepay) AddBorrowRepay() error {
	var borrowRepay = map[string]interface{}{
		"borrow_id":       b.BorrowId,
		"period":          b.Period,
		"capital":         b.Capital,
		"interest":        b.Interest,
		"repay_status":    b.RepayStatus,
		"repay_time":      b.RepayTime,
		"repay_time_true": b.RepayTimeTrue,
	}
	return models.AddBorrowRepay(borrowRepay)
}

func (b *BorrowRepay) GetBorrowRepays() ([]*models.BorrowRepay, error) {
	return models.GetBorrowRepays(b.BorrowId)
}
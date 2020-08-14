package invest_service

import (
	"html_api/models"
)

type InvestInfo struct {
	BorrowId int
	MemberId int
	Amount   float64
}

func (i *InvestInfo) Add() error {
	info := map[string]interface{}{
		"borrow_id" : i.BorrowId,
		"member_id" : i.MemberId,
		"amount" : i.Amount,
	}
	if err := models.AddInvest(info); err != nil {
		return err
	}
	return nil
}

func (i *InvestInfo) GetMemberInvest(page, page_size int)([]*models.InvestList,error) {
	return models.GetInvestList(i.MemberId, page, page_size)
}
package borrow_summary

import "html_api/models"

type BorrowSummary struct {
	BorrowID int
}

func (b *BorrowSummary) GetBorrowSummary() (*models.BorrowSummary, error) {
	return models.GetBorrowSummary(b.BorrowID)
}

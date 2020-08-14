package cache_service

import (
	"strconv"
	"strings"

	"html_api/pkg/e"
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

	PageNum int
	PageSize int
}

//func (a *Borrow) GetArticleKey() string {
//	return e.CACHE_ARTICLE + "_" + strconv.Itoa(a.ID)
//}
//
func (a *Borrow) GetBorrowsKey() string {
	keys := []string{
		e.CACHE_BORROW,
		"LIST",
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
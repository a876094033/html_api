package borrow_apply

import "html_api/models"

type BorrowApply struct {
	Name string
	Phone string
	Sex int
	Birth string
	Email string
	ApplyType int
	PropertyType int
	SeniorAmount float64
	ApplyAmount float64
	Period int
	CaseNumber string
	Postcode string
	Address string
	AddressDetail string
}
func (b *BorrowApply) Add() error {
	borrowapply := map[string]interface{}{
		"name" : b.Name,
		"phone" : b.Phone,
		"sex" : b.Sex,
		"birth" : b.Birth,
		"email" : b.Email,
		"apply_type" : b.ApplyType,
		"property_type" : b.PropertyType,
		"senior_amount" : b.SeniorAmount,
		"apply_amount" : b.ApplyAmount,
		"period" : b.Period,
		"case_number" : b.CaseNumber,
		"postcode" : b.Postcode,
		"address" : b.Address,
		"address_detail" : b.AddressDetail,
	}

	if err := models.AddBorrowApply(borrowapply); err != nil {
		return err
	}
	return nil
}
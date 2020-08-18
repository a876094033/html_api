package register_service

import (
	"html_api/models"
	"html_api/pkg/util"
)

type Register struct {
	Email      string
	Password   string
	InviteCode string
}

func (r *Register) Add() int {
	register := map[string]interface{}{
		"email":       r.Email,
		"password":    util.EncodeMD5(r.Password),
		"invite_code": r.InviteCode,
	}
	member_id := models.AddMember(register)
	if member_id <= 0 {
		return 0
	}
	return member_id
}

func (r *Register) CheckEmail() int {
	return models.CheckEmail(r.Email)
}

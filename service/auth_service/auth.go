package auth_service

import (
	"html_api/models"
	"html_api/pkg/util"
)

type Auth struct {
	ID           int
	Username     string
	Password     string
	AmountAll    float64
	AmountFrozen float64
	AmountUsed   float64
}

func (a *Auth) Check() (int, error) {
	return models.CheckAuth(a.Username, util.EncodeMD5(a.Password))
}

func (a *Auth) GetMember() (*models.Member, error) {
	return models.GetMember(a.ID)
}

func (a *Auth) UpdateMember() error {
	data := map[string]interface{}{
		"amount_all":    a.AmountAll,
		"amount_frozen": a.AmountFrozen,
		"amount_used":   a.AmountUsed,
	}
	return models.UpdateMember(a.ID, data)
}

func (a *Auth) GetMemberInfo() (*models.MemberInfo, error) {
	return models.GetMemberInfo(a.ID)
}

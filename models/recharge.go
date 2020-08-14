package models

type recharge struct {
	Model

	//RechargeId int     `json:"recharge_id" gorm:"index"`
	Amount     float64 `json:"amount"`
	MemberId   int `json:"member_id"`
	Status     int     `json:"status"`
}

func AddRecharge(maps map[string]interface{}) error {
	recharge := recharge{
		Amount : maps["amount"].(float64),
		MemberId: maps["member_id"].(int),
		Status: 0,
	}

	db.LogMode(true)
	if err := db.Debug().Create(&recharge).Error; err != nil {
		return err
	}
	return nil
}

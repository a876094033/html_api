package models

type Card struct {

	//RechargeId int     `json:"recharge_id" gorm:"index"`
	CardNumber string `json:"card_number"`
	MemberId   int    `json:"member_id"`
	Bank       string `json:"bank"`
}

func AddCard(maps map[string]interface{}) error {
	card := Card{
		CardNumber: maps["card_number"].(string),
		MemberId:   maps["member_id"].(int),
		Bank:       maps["bank"].(string),
	}

	db.LogMode(true)
	if err := db.Debug().Create(&card).Error; err != nil {
		return err
	}
	return nil
}

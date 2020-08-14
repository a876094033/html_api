package recharge_service

import "html_api/models"

type Recharge struct {
	Amount   float64
	MemberId int
	Status   int
}

type Card struct {
	MemberId   int
	CardNumber string
	Bank       string
}

func (r *Recharge) Add() error {
	recharge := map[string]interface{}{
		"amount":    r.Amount,
		"member_id": r.MemberId,
	}
	if err := models.AddRecharge(recharge); err != nil {
		return err
	}
	return nil
}

func (c *Card) Add() error {
	card := map[string]interface{}{
		"member_id":   c.MemberId,
		"card_number": c.CardNumber,
		"bank":        c.Bank,
	}
	if err := models.AddCard(card); err != nil {
		return err
	}
	return nil
}

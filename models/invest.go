package models

type Invest struct {
	ID           int     `json:"invest_id" gorm:"primary_key"`
	MamberId     int     `json:"mamber_id"`
	InvestAmount float64 `json:"invest_amount"`
	BorrowId     int     `json:"borrow_id"`
}

func AddInvest(data map[string]interface{}) error {
	invest := Invest{
		MamberId:     data["member_id"].(int),
		BorrowId:     data["borrow_id"].(int),
		InvestAmount: data["amount"].(float64),
	}
	if err := db.Debug().Create(&invest).Error; err != nil {
		return err
	}
	return nil
}

type InvestList struct {
	ID           int
	BorrowName   string
	InvestAmount float64
	InterestRate float64
	Term         int
	BorrowStatus int
	TermType     int
	BorrowId     int
}

func GetInvestList(member_id, page, page_size int) ([]*InvestList, error) {
	list := []*InvestList{}
	if err := db.Raw("select * from invest left join borrow on invest.borrow_id = borrow.id where invest.mamber_id = ? order by id desc limit ?,?", member_id, page, page_size).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func GetInvestCount(borrow_id int) (count int) {
	db.Where("borrow_id = ?", borrow_id).Count(&count)
	return
}

//func GetInvestSum(borrow_id int) (sum float64) {
//	db.Debug().Select("sum(id) as sum").Where("borrow_id = ?", borrow_id).First(&sum)
//	return sum
//}

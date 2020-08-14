package app

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type Calculation struct {
	Amount    float64
	Term      int
	TermType  int
	RepayType int
	Periods   int //期数
	Interest  float64
}
type RepayList struct {
	Capital   float64
	Interest  float64
	Period    int
	RepayTime string
}

func (calc *Calculation) CalcBorrowInterest() ([]RepayList, error) {
	if calc.TermType == 1 {
		//月
		calc.Periods = calc.Term
	} else {
		//天
		calc.Periods = 1
	}
	if calc.RepayType == 1 {
		//等额本息
		return calc.CapitalInterest()
	} else {
		//等额本金
		return calc.Capital()
	}
}

//等额本息
func (calc *Calculation) CapitalInterest() ([]RepayList, error) {
	var (
		less_capital = calc.Amount
		list         []RepayList
		nowTime      = time.Now()
		lis RepayList
	)
	if calc.TermType == 1 {
		for i := 1; i <= calc.Periods; i++ {
			//monthAmount := calc.Amount * calc.Interest /  12.0  * math.Pow(1.0 + calc.Interest / 12.0 , float64(calc.Term)) / (math.Pow( 1.0 + calc.Interest / 12.0, float64(calc.Term)) - 1.0)

			compound_rate := math.Pow(1 + calc.Interest/12, float64(calc.Term))
			pmt := float64(calc.Amount) * calc.Interest/12 * compound_rate / (compound_rate - 1)
			//获取利息
			interest := ((pmt * float64(calc.Term) ) - calc.Amount)/ 12
			//获取本金
			capital := calc.Amount / float64(calc.Term)
			//获取下个月时间
			getTime := nowTime.AddDate(0, i, 0)
			nextTime := getTime.Format("2006-01-02")
			lis.Interest = interest
			lis.RepayTime = nextTime
			lis.Capital = capital
			lis.Period = i
			list = append(list, lis)
			//less_capital -= capital
		}
	} else {
		interest, _ := strconv.ParseFloat(fmt.Sprintf("%2f", less_capital*calc.Interest), 64)
		getTime := nowTime.AddDate(0, 0, calc.Term)
		nextTime := getTime.Format("2006-01-02")
		lis.Interest = interest
		lis.RepayTime = nextTime
		lis.Capital = calc.Amount
		lis.Period = 1
		list = append(list, lis)
	}
	return list, nil
}

//等额本金
func (calc *Calculation) Capital() ([]RepayList, error) {
	var (
		less_capital = calc.Amount
		list         []RepayList
		nowTime      = time.Now()
		lis RepayList
	)
	if calc.TermType == 1 {
		for i := 1; i <= calc.Periods; i++ {
			//获取本金
			capital, _ := strconv.ParseFloat(fmt.Sprintf("%2f", calc.Amount/float64(calc.Periods)), 64)
			//获取利息
			interest, _ := strconv.ParseFloat(fmt.Sprintf("%2f", less_capital*calc.Interest/12), 64)
			//获取下个月时间

			getTime := nowTime.AddDate(0, i, 0)
			nextTime := getTime.Format("2006-01-02")
			lis.Interest = interest
			lis.RepayTime = nextTime
			lis.Capital = capital
			lis.Period = i
			list = append(list, lis)
			//list = []RepayList{
			//	lis,
			//}
			less_capital -= capital
		}
	} else {
		interest, _ := strconv.ParseFloat(fmt.Sprintf("%2f", less_capital*calc.Interest), 64)
		getTime := nowTime.AddDate(0, 0, calc.Term)
		nextTime := getTime.Format("2006-01-02")
		lis.Interest = interest
		lis.RepayTime = nextTime
		lis.Capital = calc.Amount
		lis.Period = 1
		list = append(list, lis)
	}
	return list, nil
}

func exponent(a float64, n int) float64 {
	var result float64 = 1
	for i := 1; i <= n; i++ {
		result *= a
	}
	return result
}

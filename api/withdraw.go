package api

import (
	err "2_Test/api/utility"
	"fmt"
)

func Returncash(mtype float32, num int) {
	cashier.Remove(mtype, num)
}

func Withd(amount float32) (map[float32]int, error) {
	orig := "Withd"

	withdm := make(map[float32]int)

	if amount < 0 {
		return withdm, &err.Myerr{
			Origin:  orig,
			Message: fmt.Errorf("invalid amount"),
		}
	}

	for amount > 0 {
		mtype, _, take, result := getmoney(amount)
		withdm[mtype] = take
		if result != nil {
			//return withdraw money to cassets
			for k, v := range withdm {
				err := refund(k, v)
				if err != nil {
					return withdm, err
				}
			}
			return withdm, result
		}

		amount -= (mtype * float32(take))
	}
	return withdm, nil
}

func getmoney(amount float32) (float32, int, int, error) {
	orig := "getmoney"

	take := 0
	num := 0
	var mtype float32
	var result error
	result = nil

	if amount/1000 >= 1 && cashier.Getnum(1000) > 0 {
		mtype = 1000
		num = int(amount / 1000)
		take = cashier.Remove(1000, num)
	} else if amount/500 >= 1 && cashier.Getnum(500) > 0 {
		mtype = 500
		num = int(amount / 500)
		take = cashier.Remove(500, num)
	} else if amount/100 >= 1 && cashier.Getnum(100) > 0 {
		mtype = 100
		num = int(amount / 100)
		take = cashier.Remove(100, num)
	} else if amount/50 >= 1 && cashier.Getnum(50) > 0 {
		mtype = 50
		num = int(amount / 50)
		take = cashier.Remove(50, num)
	} else if amount/20 >= 1 && cashier.Getnum(20) > 0 {
		mtype = 20
		num = int(amount / 20)
		take = cashier.Remove(20, num)
	} else if amount/10 >= 1 && cashier.Getnum(10) > 0 {
		mtype = 10
		num = int(amount / 10)
		take = cashier.Remove(10, num)
	} else if amount/5 >= 1 && cashier.Getnum(5) > 0 {
		mtype = 5
		num = int(amount / 5)
		take = cashier.Remove(5, num)
	} else if amount/1 >= 1 && cashier.Getnum(1) > 0 {
		mtype = 1
		num = int(amount / 1)
		take = cashier.Remove(1, num)
	} else {
		//Get 0.25 coin
		mtype = 0.25
		num = int(amount / 0.25)
		take = cashier.Remove(0.25, num)
	}

	//Insufficient cash in casset
	if take < num && mtype == 0.25 {
		result = &err.Myerr{
			Origin:  orig,
			Message: fmt.Errorf("insufficient money in casset for chance"),
		}
	}

	return mtype, num, take, result
}

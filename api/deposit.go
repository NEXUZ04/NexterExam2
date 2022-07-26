package api

import (
	err "2_Test/api/utility"
	"fmt"
)

func Dep(mtype float32) error {
	orig := "Dep"
	if cashier.Getnum(mtype) == cashier.Getmax(mtype) {
		return &err.Myerr{
			Origin:  orig,
			Message: fmt.Errorf("casset full"),
		}
	}

	_ = cashier.Add(mtype, 1)
	return nil
}

func refund(mtype float32, num int) error {
	orig := "refund"

	ex := cashier.Add(mtype, num)
	if ex > 0 {
		return &err.Myerr{
			Origin:  orig,
			Message: fmt.Errorf("refund money to casset error"),
		}
	}

	return nil
}

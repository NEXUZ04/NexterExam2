package api

import (
	core "2_Test/api/core"
	err "2_Test/api/utility"
	"fmt"
)

var cashier core.Cashier
var datas = make(map[float32]core.Notedata)

func Initcasset(mtype float32, n, maxn int) error {
	orig := "Initcasset"

	if mtype <= 0 || n < 0 || maxn < 0 {
		return &err.Myerr{
			Origin:  orig,
			Message: fmt.Errorf("input cannot be negative"),
		}
	}

	data := core.Notedata{
		Num:    n,
		Maxnum: maxn,
	}

	datas[mtype] = data
	return nil
}

func Startcashier() {
	cashier = core.Initstore(datas)
}

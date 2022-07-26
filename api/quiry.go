package api

func Getcasset(mtype float32) int {
	return cashier.Getnum(mtype)
}

func Getfull() []float32 {
	full := []float32{}
	if cashier.Getnum(1000) == cashier.Getmax(1000) {
		full = append(full, 1000)
	}

	if cashier.Getnum(500) == cashier.Getmax(500) {
		full = append(full, 500)
	}

	if cashier.Getnum(100) == cashier.Getmax(100) {
		full = append(full, 100)
	}

	if cashier.Getnum(50) == cashier.Getmax(50) {
		full = append(full, 50)
	}

	if cashier.Getnum(20) == cashier.Getmax(20) {
		full = append(full, 20)
	}

	if cashier.Getnum(10) == cashier.Getmax(10) {
		full = append(full, 10)
	}

	if cashier.Getnum(5) == cashier.Getmax(5) {
		full = append(full, 5)
	}

	if cashier.Getnum(1) == cashier.Getmax(1) {
		full = append(full, 1)
	}

	if cashier.Getnum(0.25) == cashier.Getmax(0.25) {
		full = append(full, 0.25)
	}

	return full
}

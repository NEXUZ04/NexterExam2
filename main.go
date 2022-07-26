package main

import (
	api "2_Test/api"
	"fmt"
	"strconv"
)

func main() {
	stop := false
	//Console app
	for !stop {
		fmt.Println("Please set 9 cashier's casset with limit and initial amount of cash number")
		fmt.Println("***************************** Press Q to exit ****************************")
		fmt.Println("--------------------------------------------------------------------------")

		cashtype := []float32{1000, 500, 100, 50, 20, 10, 5, 1, 0.25}
		var liminput string
		var initinput string
		success := false

		for i := 0; i < len(cashtype); {
			//Set limit of cash number
			fmt.Printf(">> For %.2f casset type, limit number of cash = ", cashtype[i])
			fmt.Scanln(&liminput)
			if liminput == "Q" || liminput == "q" {
				stop = true
				break
			}

			limint, err := strconv.Atoi(liminput)
			if err != nil || limint < 0 {
				fmt.Printf("Invalid input: %s\n", liminput)
				break
			}

			//Set initial amount of cash number
			fmt.Printf(">> For %.2f casset type, initial number of cash = ", cashtype[i])
			fmt.Scanln(&initinput)
			if initinput == "Q" || initinput == "q" {
				stop = true
				break
			}

			initint, err := strconv.Atoi(initinput)
			if err != nil || initint < 0 {
				fmt.Printf("Invalid input: %s\n", initinput)
				break
			}

			//Compare 2 input
			if limint < initint {
				fmt.Printf("Limit amount [%d] cannot less than initial amount [%d]\n", limint, initint)
				break
			}

			//Initial casset
			err = api.Initcasset(cashtype[i], initint, limint)
			if err != nil {
				fmt.Println(err)
				break
			}

			i++
			if i >= 8 {
				success = true
			}
		}

		//Start cashier
		if success {
			api.Startcashier()

			for !stop {
				stop = opencashier()
			}
		}

		fmt.Println("--------------------------------------------------------------------------")
		fmt.Printf("\n")
	}
	fmt.Println("Application stop!")
}

func opencashier() bool {
	var nextcust string
	var priceinput string

	//Waiting for customer input
	fmt.Println(">> Welcome next customer? [N: Close cashier] ")
	fmt.Scanln(&nextcust)
	if nextcust == "N" || nextcust == "n" {
		//End program
		return true
	}

	//Set product price
	fmt.Println(">> Product price = ")
	fmt.Scanln(&priceinput)
	price, err := strconv.ParseFloat(priceinput, 64)
	if err != nil || price < 0 {
		fmt.Printf("Invalid price: %s\n", priceinput)
		//Just cancel current transaction
		return false
	}

	fmt.Println("Casset status before transaction")
	getcassetstatus()

	starttrx(price)

	fmt.Println("Casset status after transaction")
	getcassetstatus()

	return false
}

func starttrx(price float64) {

	fmt.Println(">> Please insert note or coin one by one")
	endtrx := false
	cashinput := make(map[float32]int)
	var curramount float64

	//Get full warning
	full := api.Getfull()
	for _, v := range full {
		fmt.Printf("Warning: this cashier cannot receive cash denom %.2f\n", v)
	}

	for !endtrx {
		var denominput string

		//Get cash from customer
		fmt.Println(">> Customer insert cash denom = ")
		fmt.Scanln(&denominput)

		denom, err := strconv.ParseFloat(denominput, 64)
		if err != nil || (denom != 1000 &&
			denom != 500 &&
			denom != 100 &&
			denom != 50 &&
			denom != 20 &&
			denom != 10 &&
			denom != 5 &&
			denom != 1 &&
			denom != 0.25) {

			fmt.Printf("Invalid demon: %s\n", denominput)
		} else {
			//Deposit cash from customer
			if err = api.Dep(float32(denom)); err != nil {
				fmt.Println(err)
				fmt.Println("--- Return cash to customer ---")
			} else {
				cashinput[float32(denom)] += 1
				curramount += denom
				fmt.Printf("Current cash input = %.2f\n", curramount)
			}

			if curramount >= price {
				var withdm map[float32]int
				if withdm, err = api.Withd((float32(curramount - price))); err != nil {
					fmt.Println(err)
					for k, v := range cashinput {
						api.Returncash(k,v)
						fmt.Println("Return money to customer")
						fmt.Printf("Denom[%.2f] = %d\n", k, v)
					}
				} else {
					for k, v := range withdm {
						fmt.Println("Return change money to customer")
						fmt.Printf("Denom[%.2f] = %d\n", k, v)
					}
				}
				fmt.Println("--- Transaction end ---")
				endtrx = true
			} else {
				fmt.Println("Not enough!")
			}
		}
	}
}

func getcassetstatus() {
	fmt.Printf("denom 1000 = %d\n", api.Getcasset(1000))
	fmt.Printf("denom 500  = %d\n", api.Getcasset(500))
	fmt.Printf("denom 100  = %d\n", api.Getcasset(100))
	fmt.Printf("denom 50   = %d\n", api.Getcasset(50))
	fmt.Printf("denom 20   = %d\n", api.Getcasset(20))
	fmt.Printf("denom 10   = %d\n", api.Getcasset(10))
	fmt.Printf("denom 5    = %d\n", api.Getcasset(5))
	fmt.Printf("denom 1    = %d\n", api.Getcasset(1))
	fmt.Printf("denom 0.25 = %d\n", api.Getcasset(0.25))
}

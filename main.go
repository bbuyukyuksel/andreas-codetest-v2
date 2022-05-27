package main

import (
	"fmt"
	"time"
)

func Handler(csv *CSV) {
	last_n_second := 500
	tick_time := time.Now()

	var step int = 0
	for len(csv.customers) > 0 {
		step++
		fmt.Println("\n\n..............................................")
		fmt.Printf("STEP [ %2d ] elapsed time: %v", step, time.Now().Sub(tick_time))

		csv.Filter()
		csv.Sort()
		last_n := csv.GetLastNSecs(last_n_second, tick_time)

		PrintCustomerArray(last_n)
		PostRequest(last_n)

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Processes are done!")
}

func PostRequest(customer []*Customer) {
	for i := 0; i < len(customer); i++ {
		customer[i].ScheduleList = customer[i].ScheduleList[1:]
	}
}

func main() {
	var csv CSV
	csv.Parse("customers.csv")
	csv.Print()

	Handler(&csv)

	// CSV TEST
	// var csv CSV
	// csv.Parse("customers.csv")
	// csv.Print()
	// fmt.Println(csv.customers)

	// csv.Sort()
	// csv.Print()
	// for i, _ := range csv.customers {
	// 	csv.customers[i].ScheduleList = nil
	// }
	// csv.Filter()
	// fmt.Println("printing")
	// csv.Print()
}

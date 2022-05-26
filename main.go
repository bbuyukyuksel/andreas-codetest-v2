package main

import (
	"fmt"
)

func main() {
	var csv CSV
	csv.Parse("customers.csv")
	csv.Print()
	fmt.Println(csv.customers)

	csv.Sort()
	csv.Print()
	csv.customers[0].ScheduleList = nil
	csv.Filter()
	fmt.Println("printing")
	csv.Print()

}

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type CSV struct {
	customers []Customer
}

func (ptr *CSV) Parse(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	ptr.customers = createCustomers(data)
}

func createCustomers(data [][]string) []Customer {
	var customers []Customer
	for i, line := range data {
		if i > 0 { // omit header line
			var temp Customer
			for j, field := range line {
				if j == 0 {
					temp.Email = field
				} else if j == 1 {
					temp.Text = field
				} else if j == 2 {
					temp.Schedule = field

					// Convert schedule string to []int
					for _, value := range strings.Split(field, "-") {
						val, err := strconv.Atoi(strings.TrimRight(value, "s"))
						if err != nil {
							log.Fatal(err)
						}
						temp.ScheduleList = append(temp.ScheduleList, val)
					}
				}
			}
			customers = append(customers, temp)
		}
	}
	return customers
}

func (ptr *CSV) Sort() {
	sort.Slice(ptr.customers, func(i, j int) bool {
		return ptr.customers[i].ScheduleList[0] < ptr.customers[j].ScheduleList[0]
	})
}

func (ptr *CSV) Filter() {
	var customer []Customer

	for i, c := range ptr.customers {
		if len(c.ScheduleList) == 0 || c.Paid == true {

			var reason string
			if len(c.ScheduleList) == 0 {
				reason = color.GreenString("empty schedule")
			} else {
				reason = color.RedString("paid")
			}

			color.Yellow("[ %2d ] filtered, reason <%s>\n", i, reason)
			continue
		}
		customer = append(customer, c)
	}
	ptr.customers = nil
	ptr.customers = customer
}

func (o CSV) GetAllScheduleList() []int {
	var schedules []int
	for _, v := range o.customers {
		schedules = append(schedules, v.ScheduleList...)
	}
	sort.Slice(schedules, func(i, j int) bool {
		return schedules[i] < schedules[j]
	})
	return schedules
}

func (ptr *CSV) GetLastNSecs(last_n_time time.Duration, start_time time.Time) []*Customer {
	var customer []*Customer

	for i, c := range ptr.customers {
		if len(c.ScheduleList) == 0 {
			continue
		} else if TimeWithOffset(c.ScheduleList[0], start_time).Sub(time.Now()) < last_n_time {
			customer = append(customer, &ptr.customers[i])

		}
	}
	return customer
}

func (o CSV) Print() {
	fmt.Println("#####################CSV PRINT#####################")
	for i, c := range o.customers {
		fmt.Printf("- Data[ %-2d ]\n", i)
		fmt.Printf("\t\tEmail        : %s\n", c.Email)
		fmt.Printf("\t\tText         : %s\n", c.Text)
		fmt.Printf("\t\tSchedule     : %s\n", c.Schedule)
		fmt.Printf("\t\tScheduleList : %v\n", c.ScheduleList)
		fmt.Printf("\t\tPaid         : %v\n", c.Paid)
	}
	fmt.Println("#####################CSV PRINT#####################")
}

func PrintCustomerArray(o []*Customer) {
	if len(o) > 0 {
		fmt.Println("< CSV Print Function >")
		for i, c := range o {
			fmt.Printf("- Data[ %-2d ]\n", i)
			fmt.Printf("\t\tEmail        : %s\n", c.Email)
			fmt.Printf("\t\tText         : %s\n", c.Text)
			fmt.Printf("\t\tSchedule     : %s\n", c.Schedule)
			fmt.Printf("\t\tScheduleList : %v\n", c.ScheduleList)
			fmt.Printf("\t\tPaid         : %v\n", c.Paid)
		}
		fmt.Println("</ CSV Print Function >")
	}
}

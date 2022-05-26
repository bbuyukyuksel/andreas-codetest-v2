package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Customer struct {
	Email        string `json:"email"`
	Text         string `json:"text"`
	Schedule     string `json:"schedule"`
	Paid         bool   `json:"paid"`
	ScheduleList []int  `json:"-"`
}

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
					for i, value := range strings.Split(field, "-") {
						val, err := strconv.Atoi(strings.TrimRight(value, "s"))
						// calculate delay from previous schedule offset-time
						if i != 0 {
							val -= temp.ScheduleList[i-1]
						}
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
		if len(c.ScheduleList) == 0 {
			fmt.Println("Filtered index: ", i)
			continue
		}
		customer = append(customer, c)
	}
	ptr.customers = nil
	ptr.customers = customer
}

func (o CSV) Print() {
	for i, c := range o.customers {
		fmt.Printf("- Data[ %-2d ]\n", i)
		fmt.Printf("\t\tEmail        : %s\n", c.Email)
		fmt.Printf("\t\tText         : %s\n", c.Text)
		fmt.Printf("\t\tSchedule     : %s\n", c.Schedule)
		fmt.Printf("\t\tScheduleList : %v\n", c.ScheduleList)
		fmt.Printf("\t\tPaid         : %v\n", c.Paid)
	}
}

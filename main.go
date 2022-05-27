package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Handler(csv *CSV) {
	var schedules []int
	var adaptiveSleepTime time.Duration

	var last_n_time time.Duration = 10 * time.Millisecond
	tick_time := time.Now()

	var step int = 0
	for len(csv.customers) > 0 {
		step++

		csv.Filter()
		csv.Sort()
		last_n := csv.GetByLastNTime(last_n_time, tick_time)
		if len(last_n) > 0 {
			fmt.Println("\n\n..............................................")
			fmt.Printf("STEP [ %2d ] elapsed time: %v\n", step, time.Now().Sub(tick_time))
		}
		//PrintCustomerArray(last_n)
		PostRequest(last_n)

		schedules = csv.GetAllScheduleList()
		if len(schedules) > 0 {
			adaptiveSleepTime = TimeWithOffset(schedules[0], tick_time).Sub(time.Now().Add(5 * time.Millisecond))
			fmt.Println("Adaptive Sleep Time", adaptiveSleepTime)
			time.Sleep(adaptiveSleepTime)
		} else {
			time.Sleep(100 * time.Microsecond)
		}
	}

	fmt.Println("Processes are done!")
}

func PostRequest(customer []*Customer) {
	for i := 0; i < len(customer); i++ {

		customer[i].ScheduleList = customer[i].ScheduleList[1:]

		go func(customer *Customer) {
			data, _ := json.Marshal(customer)
			fmt.Println(string(data))

			client := http.Client{Timeout: time.Duration(1) * time.Second}
			payload := bytes.NewBuffer([]byte(string(data)))
			resp, err := client.Post("http://127.0.0.1:9090/messages", "application/json", payload)

			if err != nil {
				log.Fatal(err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			fmt.Printf("Reponse Status Code: %d, Body : %s", resp.StatusCode, body)

			// Unmarshall response
			var response Customer
			json.Unmarshal([]byte(body), &response)
			customer.Paid = response.Paid
		}(customer[i])
	}
}

func main() {
	var csv CSV
	csv.Parse("customers.csv")
	// csv.Parse("concurrent_customer.csv")
	csv.Print()

	Handler(&csv)
}

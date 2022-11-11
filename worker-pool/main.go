package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixMilli()
	processes := make(chan int, 5)
	results := make(chan *User, 5)

	quantityWorkers := 5
	quantityUsers := 5

	for i := 1; i <= quantityWorkers; i++ {
		go worker(processes, results, i)
	}

	for i := 1; i <= quantityUsers; i++ {
		processes <- i
	}

	close(processes)

	for i := 0; i < quantityUsers; i++ {
		fmt.Printf("User processed %+v \n", <-results)
	}

	end := time.Now().UnixMilli()
	fmt.Printf("Time elapsed: %ds\n", (end-start)/1000)
}

type User struct {
	FirstName string
	LastName  string
}

func worker(processes <-chan int, results chan<- *User, workerId int) {
	for process := range processes {
		user := request(process, workerId)
		calculate(process, user, workerId)
		save(process, user, workerId)

		results <- user
	}
}

func request(process int, workerId int) (user *User) {
	fmt.Printf("Init request [%d] - Worker [%d]\n", process, workerId)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finish request [%d] - Worker [%d]\n", process, workerId)
	return &User{
		FirstName: fmt.Sprintf("Vitor %d", process),
		LastName:  "Medeiro",
	}
}

func calculate(process int, user *User, workerId int) {
	fmt.Printf("Init calculate [%d] - Worker [%d]\n", process, workerId)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Finish calculate [%d] - Worker [%d]\n", process, workerId)
}

func save(process int, user *User, workerId int) {
	fmt.Printf("Init save [%d] - Worker [%d]\n", process, workerId)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finish save [%d] - Worker [%d]\n", process, workerId)
}

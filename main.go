package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	threads, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("Bad threads")
	}

	for i := 1; i < threads; i++ {
		go func() {
			start_time := time.Now()
			db, err := sql.Open("postgres", os.Args[1])
			if err != nil {
				fmt.Print(err.Error())
			}
			timeTrack(start_time, "Connecting")
			defer db.Close()
			start_time = time.Now()
			_, err = db.Exec("SELECT * FROM CashOrders")
			if err != nil {
				fmt.Print(err.Error())
			}
			timeTrack(start_time, "Executing")

		}()

	}
}

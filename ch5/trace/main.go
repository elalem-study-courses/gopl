// package description
package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(time.Second * 10)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("Enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}

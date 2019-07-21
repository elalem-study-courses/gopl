// package description
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type LocationTimeMap map[string]string

func (ltm LocationTimeMap) String() string {
	str := ""
	for location, time := range ltm {
		if str != "" {
			str += " "
		}
		str += fmt.Sprintf("%s:%s ", location, time)
	}
	return str
}

var (
	LocationsHost map[string]string
	LocationsTime LocationTimeMap
	wg            sync.WaitGroup
	mutex         sync.RWMutex
)

func init() {
	LocationsHost = make(map[string]string)
	LocationsTime = make(LocationTimeMap)
	loadLocations()
}

func main() {
	for location, address := range LocationsHost {
		wg.Add(1)
		go connect(location, address)
	}

	go printDates()

	wg.Wait()
}

func loadLocations() {
	for _, arg := range os.Args[1:] {
		loadLocation(arg)
	}
}

func loadLocation(arg string) {
	info := strings.Split(arg, "=")
	LocationsHost[info[0]] = info[1]
}

func connect(location, address string) {
	defer wg.Done()

	log.Printf("%s connecting to %s", location, address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	readDate(conn, location)
}

func readDate(conn net.Conn, location string) {
	reader := bufio.NewReader(conn)
	for {
		date, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		mutex.Lock()
		LocationsTime[location] = string(date)
		mutex.Unlock()
	}
}

func printDates() {
	for {
		mutex.Lock()
		fmt.Println(LocationsTime)
		mutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}

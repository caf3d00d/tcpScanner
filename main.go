package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func handleArgs(Args []string) (string, int, int) {
	if len(Args) != 4 {
		log.Fatal("Usage: tcpScanner <host> <range from> <range to>")
	}

	rf, err := strconv.Atoi(Args[2])
	if err != nil {
		log.Fatal("Error occurred in Atoi")
	}
	rt, err := strconv.Atoi(Args[3])
	if err != nil {
		log.Fatal("Error occurred in Atoi")
	}

	if rf > rt {
		log.Fatal("Invalid range")
	}

	return Args[1], rf, rt
}

func main() {
	host, rf, rt := handleArgs(os.Args)
	var openedPorts []string

	for rf <= rt {
		address := fmt.Sprintf("%s:%d", host, rf)
		// Didn't know port 0 is invalid
		if rf == 0 {
			rf++
			continue
		}
		connection, err := net.Dial("tcp", address)

		if err != nil {
			fmt.Printf("\033[0;31m* %s -> closed\n\u001B[0;37m", address)
			rf++
			continue
		}
		err = connection.Close()
		if err != nil {
			log.Fatal("\u001B[0;31mCouldn't close a connection\u001B[0;37m")
		}
		fmt.Printf("\033[0;32m* %s -> opened\n\u001B[0;37m", address)
		openedPorts = append(openedPorts, address)
		rf++
	}
	i := 0
	if len(openedPorts) > 0 {
		fmt.Printf("\u001B[0;32m\n\nSummary, Opened ports:\n")
		for range openedPorts {
			fmt.Printf("%d. %s\n\033[0;37m", i+1, openedPorts[i])
			i++
		}
	} else {
		fmt.Printf("\033[0;31m\n\nSummary, no opened port found\n\u001B[0;37m")
	}
}

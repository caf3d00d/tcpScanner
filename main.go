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

	for rf <= rt {
		address := fmt.Sprintf("%s:%d", host, rf)
		//fmt.Printf("rf -> %d\n", rf)
		connection, err := net.Dial("tcp", address)

		if err != nil {
			rf++
			continue
		}
		err = connection.Close()
		if err != nil {
			log.Fatal("Couldn't close a connection")
		}
		fmt.Printf("found open port at: %s\n", address)
		rf++
	}

}

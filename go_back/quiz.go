package main

import (
	"flag"
	"fmt"
	"os"

	srv "./server"
)

func displayUsage() {

	fmt.Println("Usage")
}

func getFlag() string {

	peerFalg := flag.String("peer", "8080", "peer")
	helpFlag := flag.Bool("h", false, "usage")

	flag.Parse()
	if *helpFlag == true {
		displayUsage()
		os.Exit(0)
	}
	return *peerFalg
}

func main() {

	peer := getFlag()
	srv := srv.InitServer(peer)
	srv.InitHandler()
	srv.RunServer()
}

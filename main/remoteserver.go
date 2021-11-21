package main

import (
	remoteserver "ip_reporter/remoteServer"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	rs := remoteserver.NewRemoteServer(port, os.Args[2])
	rs.Run()
}

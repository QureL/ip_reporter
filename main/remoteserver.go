package main

import (
	"fmt"
	remoteserver "ip_reporter/remoteServer"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("input %s listen_port outputPath.\n", os.Args[0])
		return
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	rs := remoteserver.NewRemoteServer(port, os.Args[2])
	rs.Run()
}

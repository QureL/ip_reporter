package main

import (
	"ip_reporter/reporter"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return
	}
	r := reporter.NewReporter(os.Args[1], port)
	r.Run(10)
}

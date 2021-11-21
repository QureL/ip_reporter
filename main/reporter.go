package main

import (
	"fmt"
	"ip_reporter/reporter"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("input %s remote_ip remote_port.\n", os.Args[0])
		return

	}
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return
	}
	r := reporter.NewReporter(os.Args[1], port)
	for {
		/* 半小时上报周期 */
		r.Run(60 * 30)
	}
}

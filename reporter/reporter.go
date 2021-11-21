package reporter

import (
	"ip_reporter/common"
	"ip_reporter/network/client"
	"log"
	"time"
)

type Reporter struct {
	client client.Client
}

func NewReporter(remote_ip string, remote_port int) *Reporter {
	return &Reporter{
		client: *client.NewClient(remote_ip, remote_port),
	}
}

func (r *Reporter) Run(cycle int) {
	if cycle <= 0 {
		log.Println("wrong cycle.")
		return
	}
	ret := r.client.Connect()
	if ret != common.RESULT_OK {
		return
	}
	defer r.client.Close()

	for {
		r.client.Send([]byte("hello"))
		time.Sleep(time.Second * time.Duration(cycle))
	}

}

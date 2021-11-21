package client

import (
	"fmt"
	"ip_reporter/common"
	"log"
	"net"
)

type Client struct {
	remote_ip   string
	remote_port int
	conn        net.Conn
}

func init() {
	log.SetPrefix("[IP_REPORTER]")
}
func NewClient(remote_ip string, remote_port int) *Client {
	return &Client{
		remote_ip:   remote_ip,
		remote_port: remote_port,
	}
}

func (c *Client) Connect() int {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", c.remote_ip, c.remote_port))
	if err != nil {
		log.Println(err)
		return common.RESULT_NOK
	}

	c.conn = conn
	return common.RESULT_OK
}

func (c *Client) Close() {
	c.conn.Close()
}

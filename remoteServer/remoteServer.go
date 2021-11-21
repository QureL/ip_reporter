package remoteserver

import (
	"io/ioutil"
	"ip_reporter/common"
	"ip_reporter/network/server"
	"log"
)

type RemoteServer struct {
	server     *server.Server
	outputPath string
}

func NewRemoteServer(port int, outputPath string) *RemoteServer {
	return &RemoteServer{
		server:     server.NewServer("0.0.0.0", port),
		outputPath: outputPath,
	}
}

func (rs *RemoteServer) Run() {
	if rs.server.Listen() == common.RESULT_NOK {
		return
	}
	defer rs.server.Close()
	buffer := make([]byte, 1024)
	for {
		_, addr, err := rs.server.RecvFrom(buffer)
		if err != nil {
			log.Println(err)
			continue
		}
		ioutil.WriteFile(rs.outputPath, []byte(addr.String()), 0664)
	}
}

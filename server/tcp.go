package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"supervisory-control/agreement"
	"time"
)

func TcpRun(tcpPort int32) {

	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", tcpPort))
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer func() {
		if err := tcpListener.Close(); err != nil {
			log.Warn(err)
		}
	}()

	log.Info("Server ready to read ...")
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Warn(err)
			continue
		}
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {

	defer func() {
		if err := conn.Close(); err != nil {
			log.Warn(err)
		}
	}()

	log.Debug(conn.RemoteAddr().String())

	local := 0
	buffer := make([]byte, 1024)
	for {
		if err := (*conn).SetReadDeadline(time.Now().Add(3 * time.Minute)); err != nil {
			log.Error(err)
			break
		}

		count, err := (*conn).Read(buffer[local:])
		if err != nil {
			log.Debug(err)
			break
		}

		if local += count; local >= 12 {

			if _, err := (*conn).Write(agreement.AnalysisAgreement(buffer[0:12])); err != nil {
				log.Error(err)
			}
			buffer = make([]byte, 1024)
			local = 0
		}
	}
}

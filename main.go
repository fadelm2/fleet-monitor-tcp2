package main

import (
	"encoding/hex"
	"fleet-monior/logger"
	"fleet-monior/parser"
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	log := logger.New(logrus.InfoLevel)
	parser := parser.New(log.Logger)

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.WithError(err).Fatal("failed to start server")
	}

	log.WithField("port", 9000).
		Info("tcp server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.WithError(err).Warn("accept failed")
			continue
		}

		log.WithField("remote", conn.RemoteAddr().String()).
			Info("client connected")

		go handleConn(conn, log, parser)
	}
}

func handleConn(conn net.Conn, log *logger.Logger, parser *parser.Parser) {
	defer func() {
		log.WithField("remote", conn.RemoteAddr().String()).
			Info("client disconnected")
		conn.Close()
	}()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.WithError(err).
				WithField("remote", conn.RemoteAddr().String()).
				Warn("read error")
			return
		}

		data := buf[:n]

		log.WithFields(logrus.Fields{
			"bytes": n,
			"raw":   hex.EncodeToString(data),
		}).Debug("packet received")

		parser.Parse(data)
	}
}

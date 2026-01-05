package main

import (
	"fleet-monior/logger"
	"fleet-monior/parser"
	"net"
)

func main() {
	logger.Init()

	logger.Log.Info("🚀 TCP Server listening on :9000")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		logger.Log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		logger.Log.Infof("🔌 CONNECTED %s", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			logger.Log.Warnf("❌ DISCONNECTED %s", conn.RemoteAddr())
			return
		}

		logger.Log.Infof("📥 RECV %d bytes from %s", n, conn.RemoteAddr())
		parser.ParseAndLog(buf[:n])
	}
}

package main

import (
	"fleet-monior/logger"
	"fleet-monior/parser"
	"net"
)

func main() {
	logger.Init() // ⬅️ WAJIB DIPANGGIL

	logger.Log.Info("TCP Server starting on :9000")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		logger.Log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		logger.Log.Info("Connected: ", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			logger.Log.Warn("Disconnected: ", conn.RemoteAddr())
			return
		}
		parser.ParseAndLog(buf[:n])
	}
}

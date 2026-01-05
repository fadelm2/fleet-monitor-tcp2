package parser

import (
	"encoding/hex"
	"fleet-monior/logger"
	"github.com/sirupsen/logrus"
)

func ParseAndLog(data []byte) {

	if len(data) < 5 {
		logger.Log.Warn("Packet too short")
		return
	}

	// dump raw hex supaya yakin data masuk
	logger.Log.Infof("RAW HEX: % X", data)

	protocol := data[3]
	logger.Log.Infof("PROTO=0x%X LEN=%d", protocol, len(data))

	switch protocol {

	// LOGIN
	case 0x01:
		if len(data) < 12 {
			logger.Log.Warn("LOGIN packet too short")
			return
		}

		imei := hex.EncodeToString(data[4:12])
		logger.Log.WithField("imei", imei).Info("LOGIN PACKET")

	// GPS DATA (beberapa device beda protocol)
	case 0x10, 0x12, 0x22:
		if len(data) < 18 {
			logger.Log.Warn("GPS packet too short")
			return
		}

		lat := parseCoordinate(data[10:14])
		lon := parseCoordinate(data[14:18])

		logger.Log.WithFields(logrus.Fields{
			"lat":  lat,
			"lon":  lon,
			"maps": "https://maps.google.com/?q=",
		}).Info("GPS DATA")

	default:
		logger.Log.Warnf("UNKNOWN PROTOCOL 0x%X", protocol)
	}
}

func parseCoordinate(b []byte) float64 {
	raw := uint32(b[0])<<24 |
		uint32(b[1])<<16 |
		uint32(b[2])<<8 |
		uint32(b[3])

	return float64(raw) / 30000.0 / 60.0
}

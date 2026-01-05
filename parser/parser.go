package parser

import (
	"encoding/hex"
	"fleet-monior/logger"
)

func ParseAndLog(data []byte) {
	if len(data) < 20 {
		logger.Log.Warnf("Packet too short: %d bytes", len(data))
		return
	}

	protocol := data[3]

	switch protocol {

	case 0x01:
		imei := hex.EncodeToString(data[4:12])
		logger.Log.Info("LOGIN | IMEI=", imei)

	case 0x12:
		lat := parseCoordinate(data[10:14])
		lon := parseCoordinate(data[14:18])

		logger.Log.WithFields(map[string]interface{}{
			"lat":  lat,
			"lon":  lon,
			"maps": "https://maps.google.com/?q=",
		}).Info("GPS DATA")
	}
}
func parseCoordinate(b []byte) float64 {
	raw := uint32(b[0])<<24 |
		uint32(b[1])<<16 |
		uint32(b[2])<<8 |
		uint32(b[3])

	return float64(raw) / 30000.0 / 60.0
}

func parseIMEI(b []byte) string {
	return hex.EncodeToString(b)
}

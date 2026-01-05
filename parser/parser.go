package parser

import (
	"fmt"
	"strconv"
	"strings"

	"fleet-monior/logger"

	"github.com/sirupsen/logrus"
)

func ParseAndLog(data []byte) {

	if len(data) < 20 {
		logger.Log.Warn("Packet too short")
		return
	}

	// GT02A = ASCII
	payload := string(data)
	logger.Log.Infof("RAW ASCII: %s", payload)

	// basic validation
	if !strings.HasPrefix(payload, "(") || !strings.HasSuffix(payload, ")") {
		logger.Log.Warn("INVALID PACKET FORMAT")
		return
	}

	// Protocol (ASCII)
	proto := payload[1:3]

	// IMEI (ASCII)
	// (02 8 044400735 BR ...)
	imei := payload[5:14]

	logger.Log.Infof("PROTO=%s IMEI=%s LEN=%d", proto, imei, len(payload))

	switch proto {

	case "02": // LOCATION PACKET
		parseLocation(payload, imei)

	default:
		logger.Log.Warnf("UNKNOWN PROTOCOL %s", proto)
	}
}

func parseLocation(p, imei string) {

	/*
	   Example packet:
	   (028044400735BR00260105A0610.2215S10643.9911E000.0141322181.830100000L00000000)
	*/

	// GPS Status
	status := p[23:24]
	if status != "A" {
		logger.Log.WithField("imei", imei).Warn("GPS NOT VALID")
		return
	}

	latRaw := p[24:33] // 0610.2215
	latDir := p[33:34] // N/S
	lonRaw := p[34:44] // 10643.9911
	lonDir := p[44:45] // E/W

	lat := convertGT02ACoord(latRaw, latDir)
	lon := convertGT02ACoord(lonRaw, lonDir)

	logger.Log.WithFields(logrus.Fields{
		"imei": imei,
		"lat":  lat,
		"lon":  lon,
		"maps": fmt.Sprintf("https://www.google.com/maps?q=%f,%f", lat, lon),
	}).Info("📍 GPS LOCATION")
}

func convertGT02ACoord(raw, dir string) float64 {
	// raw = DDMM.MMMM or DDDMM.MMMM
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0
	}

	deg := float64(int(v / 100))
	min := v - (deg * 100)

	dec := deg + (min / 60)

	if dir == "S" || dir == "W" {
		dec = -dec
	}
	return dec
}

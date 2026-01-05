//package test
//
//import (
//	"bytes"
//	"fleet-monior/parser"
//	"os"
//	"strings"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//)
//
///*
//Helper untuk capture stdout
//*/
//func captureOutput(t *testing.T, f func()) string {
//	t.Helper()
//
//	r, w, err := os.Pipe()
//	require.NoError(t, err)
//
//	stdout := os.Stdout
//	os.Stdout = w
//
//	f()
//
//	_ = w.Close()
//	os.Stdout = stdout
//
//	var buf bytes.Buffer
//	_, err = buf.ReadFrom(r)
//	require.NoError(t, err)
//
//	return buf.String()
//}
//
//func TestParseIMEI(t *testing.T) {
//	data := []byte{
//		0x35, 0x39, 0x30, 0x31,
//		0x32, 0x33, 0x34, 0x35,
//	}
//
//	imei := parser.parseIMEI(data)
//
//	assert.Equal(t, "3539303132333435", imei)
//}
//
//func TestParseCoordinate(t *testing.T) {
//	// raw = 0x001C2000 = 1843200
//	data := []byte{0x00, 0x1C, 0x20, 0x00}
//
//	coord := parser.parseCoordinate(data)
//	expected := float64(1843200) / 30000.0 / 60.0
//
//	assert.InDelta(t, expected, coord, 0.000001)
//}
//
//func TestParseAndLog_LoginPacket(t *testing.T) {
//	packet := []byte{
//		0x78, 0x78, 0x0D, 0x01, // LOGIN
//		0x35, 0x39, 0x30, 0x31,
//		0x32, 0x33, 0x34, 0x35, // IMEI
//		0x00, 0x00, 0x00, 0x00,
//		0x00, 0x00, 0x00, 0x00, // ✅ padding
//	}
//
//	require.GreaterOrEqual(t, len(packet), 20)
//
//	output := captureOutput(t, func() {
//		ParseAndLog(packet)
//	})
//
//	assert.Contains(t, output, "LOGIN")
//	assert.Contains(t, output, "IMEI")
//	assert.Contains(t, output, "3539303132333435")
//}
//
//func TestParseAndLog_GPSPacket(t *testing.T) {
//	packet := []byte{
//		0x78, 0x78, 0x1F, 0x12, // protocol GPS
//		0x00, 0x00, 0x00, 0x00, // filler
//		0x00, 0x1C, 0x20, 0x00, // latitude
//		0x00, 0x1C, 0x20, 0x00, // longitude
//		0x00, 0x00, 0x00, 0x00, // ✅ padding
//	}
//
//	require.GreaterOrEqual(t, len(packet), 20)
//
//	output := captureOutput(t, func() {
//		ParseAndLog(packet)
//	})
//
//	assert.Contains(t, output, "GPS DATA")
//	assert.Contains(t, output, "Latitude")
//	assert.Contains(t, output, "Longitude")
//	assert.Contains(t, output, "maps.google.com")
//}
//func TestParseAndLog_InvalidPacket(t *testing.T) {
//	packet := []byte{0x01, 0x02, 0x03} // < 20 byte
//
//	output := captureOutput(t, func() {
//		ParseAndLog(packet)
//	})
//
//	assert.True(t, strings.TrimSpace(output) == "")
//}

package math

import (
	"bytes"
	"encoding/binary"
	"strings"
)

// Uint8ToLower return Lower case letters with uint8
//
// e.g. "D" == 68 and "d" == 100, Uint8ToLower(68) == 100
func Uint8ToLower(n uint8) uint8 {
	return []byte(strings.ToLower(string(n)))[0]
}

func Int2Bytes(i int) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.BigEndian, int64(i))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Bytes2Int64(b []byte) (int64, error) {
	var i int64
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.BigEndian, &i)
	if err != nil {
		return 0, err
	}
	return i, nil
}

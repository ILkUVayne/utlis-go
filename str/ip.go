package str

import (
	"errors"
	"strconv"
	"strings"
)

// IPStrToHost string type host to []byte host
//
// e.g. "127.0.0.1" -> []byte{127,0,0,1}
func IPStrToHost(host string) ([4]byte, error) {
	var h [4]byte
	hosts := strings.Split(host, ".")
	if len(hosts) != 4 {
		return h, errors.New("str2host error: host is bad, host: " + host + " len != 4")
	}
	for idx, v := range hosts {
		i, err := strconv.Atoi(v)
		if err != nil {
			return h, errors.New("str2host error: strconv.Atoi err " + err.Error())
		}
		h[idx] = uint8(i)
	}
	return h, nil
}

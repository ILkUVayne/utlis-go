package str

import "testing"

func TestIPStrToHost(t *testing.T) {
	ip := "192.168.55.1"
	host, err := IPStrToHost(ip)
	if err != nil {
		t.Error(err)
	}
	if host != [4]byte{192, 168, 55, 1} {
		t.Error("StrToHost err")
	}
	ip = "192.168.55.1.2"
	host, err = IPStrToHost(ip)
	if err == nil {
		t.Error("StrToHost err: ip = ", ip)
	}
	ip = "192.168.s55.1"
	host, err = IPStrToHost(ip)
	if err == nil {
		t.Error("StrToHost err: ip = ", ip)
	}
}

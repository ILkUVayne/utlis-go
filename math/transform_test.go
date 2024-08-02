package math

import "testing"

func TestUint8ToLower(t *testing.T) {
	s1 := "ABC"
	s2 := "abc"
	if u := Uint8ToLower(s1[0]); u != s2[0] {
		t.Errorf("Uint8ToLower err : u = %d, s2[0] = %d", u, s2[0])
	}
	if u := Uint8ToLower(s1[1]); u != s2[1] {
		t.Errorf("Uint8ToLower err : u = %d, s2[1] = %d", u, s2[1])
	}
	if u := Uint8ToLower(s1[2]); u != s2[2] {
		t.Errorf("Uint8ToLower err : u = %d, s2[2] = %d", u, s2[2])
	}
}

func TestInt2Bytes(t *testing.T) {
	buf, err := Int2Bytes(1555)
	if err != nil {
		t.Error(err)
	}
	i, err := Bytes2Int64(buf)
	if err != nil {
		t.Error(err)
	}
	if i != 1555 {
		t.Error("Bytes2Int64 err : i = ", i)
	}
}

package str

import "testing"

func TestCamelCaseToUnderscore(t *testing.T) {
	ts := "DoSomething"
	if res := CamelCaseToUnderscore(ts); res != "do_something" {
		t.Error("CamelCaseToUnderscore err result = ", res)
	}
	ts = "doSomething"
	if res := CamelCaseToUnderscore(ts); res != "do_something" {
		t.Error("CamelCaseToUnderscore err result = ", res)
	}
}

func TestString2Int64(t *testing.T) {
	var i int64
	s := "65"
	err := String2Int64(s, &i)
	if err != nil {
		t.Error(err)
	}
	s = "a65"
	err = String2Int64(s, &i)
	if err == nil {
		t.Error(err)
	}
}

func TestString2Float64(t *testing.T) {
	var f float64
	s := "65.2"
	err := String2Float64(s, &f)
	if err != nil {
		t.Error(err)
	}
	s = "a65"
	err = String2Float64(s, &f)
	if err == nil {
		t.Error(err)
	}
}

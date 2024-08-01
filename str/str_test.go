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

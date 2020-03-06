package strings_test

import (
	"testing"

	"github.com/x-punch/go-strings"
)

func TestRandomString(t *testing.T) {
	if len(strings.GetRandomString(0)) != 0 {
		t.Fail()
	}
	if len(strings.GetRandomString(10)) != 10 {
		t.Fail()
	}
}

func TestRandomNumber(t *testing.T) {
	if len(strings.GetRandomNumber(0)) != 0 {
		t.Fail()
	}
	if len(strings.GetRandomNumber(5)) != 5 {
		t.Fail()
	}
}

func TestSubString(t *testing.T) {
	if "12345" != strings.SubString("123456789", 0, 5) {
		t.Fail()
	}
	if "123456789" != strings.SubString("123456789", 0, 20) {
		t.Fail()
	}
	if "456" != strings.SubString("123456789", 3, 3) {
		t.Fail()
	}
}

func TestContainString(t *testing.T) {
	if !strings.ContainString("", ",", "") {
		t.Fail()
	}
	if !strings.ContainString("123,456,abc,ddd", ",", "456") {
		t.Fail()
	}
	if strings.ContainString("12.34.56,aa.admin.guest", ".", "aa") {
		t.Fail()
	}
	if strings.ContainString("12.34.56,aa.admin.guest", ".", "") {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	if !strings.ContainsInArray([]string{"12", "aa", "bb"}, "aa") {
		t.Fail()
	}
	if strings.ContainsInArray([]string{"12", "aa", "334"}, "34") {
		t.Fail()
	}
	if strings.ContainsInArray([]string{}, "34") {
		t.Fail()
	}
}

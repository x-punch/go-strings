package strings_test

import (
	"math/rand"
	"testing"

	"github.com/x-punch/go-strings"
)

func TestEncodeToBase62(t *testing.T) {
	if "0" != strings.EncodeToBase62(0) {
		t.Fail()
	}
	if "r" != strings.EncodeToBase62(27) {
		t.Fatal(strings.EncodeToBase62(27))
		t.Fail()
	}
	if "7M" != strings.EncodeToBase62(2983) {
		t.Fatal(strings.EncodeToBase62(2983))
		t.Fail()
	}
	if "yy53" != strings.EncodeToBase62(736346) {
		t.Fatal(strings.EncodeToBase62(736346))
		t.Fail()
	}
	if "kiWFb" != strings.EncodeToBase62(172535232) {
		t.Fatal(strings.EncodeToBase62(172535232))
		t.Fail()
	}
}

func TestDncodeBase62(t *testing.T) {
	if 0 != strings.DecodeBase62("0") {
		t.Fatal(strings.DecodeBase62("0"))
		t.Fail()
	}
	if 27 != strings.DecodeBase62("r") {
		t.Fatal(strings.DecodeBase62("r"))
		t.Fail()
	}
	if 2983 != strings.DecodeBase62("7M") {
		t.Fatal(strings.DecodeBase62("7M"))
		t.Fail()
	}
	if 736346 != strings.DecodeBase62("yy53") {
		t.Fatal(strings.DecodeBase62("yy53"))
		t.Fail()
	}
	if 172535232 != strings.DecodeBase62("kiWFb") {
		t.Fatal(strings.DecodeBase62("kiWFb"))
		t.Fail()
	}
}

func BenchmarkBase62_EncodeThenDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Uint64()
		t := strings.EncodeToBase62(n)
		if n != strings.DecodeBase62(t) {
			b.Fail()
		}
	}
}

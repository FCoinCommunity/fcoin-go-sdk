package fcoin

import (
	"testing"
)

func TestSign(t *testing.T) {
	// according to: https://developer.fcoin.com/en.html?python#05e32f581f
	method := "POST"
	uri := "https://api.fcoin.com/v2/orders"
	ts := "1523069544359"
	args := "amount=100.0&price=100.0&side=buy&symbol=btcusdt&type=limit"
	key := "3600d0a74aa3410fb3b1996cca2419c8"
	output := Sign(method, uri, ts, args, key)

	if output != "DeP6oftldIrys06uq3B7Lkh3a0U=" {
		t.Fatalf("output is wrong %v", output)
	}
}

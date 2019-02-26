package btcvanity

import (
	"strings"
	"testing"
	"time"
)

func TestBTCVanity(t *testing.T) {

	b := New(&Config{
		Buffer: 5,
	})

	wallet, err := b.Find("ab")
	if err != nil {
		t.Error(err)
		return
	}

	pubKey := strings.ToLower(wallet.PublicKey())
	expected := "1ab"

	hasPrefix := strings.HasPrefix(pubKey, expected)
	if !hasPrefix {
		t.Errorf("Wrong generated wallet, got %v expected %v\n", pubKey, expected)
	}

	t.Logf("Generated address %v\n", pubKey)

	done := make(chan bool)
	// this is going to take billion of years
	go func() {
		wallet, err = b.Find("helloworld")
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if wallet != nil {
			t.Error("We got generated wallet! Should not happen")
			t.Fail()
		}
		done <- true
	}()

	// wait and stop the process
loop:
	for {
		select {
		case <-done:
			t.Log("Process stopped")
			break loop
		case <-time.After(2 * time.Second):
			b.Stop()
			break
		}
	}

}

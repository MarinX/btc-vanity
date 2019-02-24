package btcvanity

import (
	"regexp"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
)

func TestGenerator(t *testing.T) {
	g := &Generator{
		params: &chaincfg.MainNetParams,
	}

	w, err := g.Generate()
	if err != nil {
		t.Error(err)
		return
	}

	if w == nil {
		t.Error("No address generated")
		return
	}

	pubKye := w.PublicKey()
	privKey := w.PrivateKey()
	t.Logf("%v\n", pubKye)
	t.Logf("%v\n", privKey)

	if len(privKey) != 51 {
		t.Errorf("Wrong size of private key, expected 51 got %v\n", len(privKey))
		return
	}

	// regex to see if this is bitcoin adddress
	matched, err := regexp.MatchString("^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$", pubKye)
	if err != nil {
		t.Error(err)
		return
	}
	if !matched {
		t.Errorf("This is not generated bitcoin address %v\n", pubKye)
		return
	}

}

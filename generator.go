package btcvanity

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// Generator hold the logic for generating bitcoin address
// Depending on chain configuration, it can generate main/test net versions of wallet
type Generator struct {
	params *chaincfg.Params
}

// IGenerator is interface for generator
type IGenerator interface {
	Generate() (IWallet, error)
}

// Generate generates bitcoin wallet interface
func (g *Generator) Generate() (IWallet, error) {
	wallet := &Wallet{}
	var err error

	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}

	wallet.privKey, err = btcutil.NewWIF(privKey, g.params, false)
	if err != nil {
		return nil, err
	}

	wallet.pubKey, err = btcutil.NewAddressPubKey(
		wallet.privKey.PrivKey.PubKey().SerializeUncompressed(), g.params)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

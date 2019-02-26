package btcvanity

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// BTCVanity library holder
type BTCVanity struct {
	config *Config
	stop   chan bool
}

// New creates a new BTCVanity service
func New(config *Config) *BTCVanity {
	return &BTCVanity{
		config: config,
		stop:   make(chan bool),
	}
}

// Find runs a service to find matching pattern
func (b *BTCVanity) Find(pattern string) (IWallet, error) {
	var resWallet IWallet
	var resError error

	cWallet := make(chan IWallet, b.config.Threads)
	cErr := make(chan error)
	var chainParams *chaincfg.Params
	if b.config.TestNet {
		chainParams = &chaincfg.TestNet3Params
	} else {
		chainParams = &chaincfg.MainNetParams
	}

	btcWorker := &worker{generator: &Generator{params: chainParams}}

loop:
	for {
		select {
		case wallet := <-cWallet:
			if isMatch(pattern, wallet.PublicKey()) {
				resWallet = wallet
				break loop
			}
			break

		case err := <-cErr:
			resError = err
			break loop

		case <-b.stop:
			break loop

		default:
			btcWorker.Work(cWallet, cErr)
		}
	}

	close(cWallet)
	close(cErr)
	return resWallet, resError
}

// Stop stops the process
func (b *BTCVanity) Stop() {
	b.stop <- true
}

package btcvanity

import (
	"sync"

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

	var chainParams *chaincfg.Params
	if b.config.TestNet {
		chainParams = &chaincfg.TestNet3Params
	} else {
		chainParams = &chaincfg.MainNetParams
	}

	btcWorker := &worker{gen: &Generator{params: chainParams}}

	mutex := sync.Mutex{}
	for i := 0; i < b.config.Buffer; i++ {
		go func() {
			for {
				wallet, err := btcWorker.Work()
				if err != nil {
					mutex.Lock()
					resError = err
					break
				}
				if isMatch(pattern, wallet.PublicKey()) {
					mutex.Lock()
					if resWallet != nil {
						break
					}
					resWallet = wallet
					break
				}
			}

			b.stop <- true
			mutex.Unlock()
		}()
	}

	<-b.stop

	return resWallet, resError
}

// Stop stops the process
func (b *BTCVanity) Stop() {
	b.stop <- true
}

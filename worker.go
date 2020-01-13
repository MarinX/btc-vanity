package btcvanity

type worker struct {
	gen IGenerator
}

// Work generates bitcoin wallet and pushes through channel
func (w *worker) Work() (result IWallet, erri error) {
	wallet, err := w.gen.Generate()
	if err != nil {
		erri = err
		return
	}
	result = wallet
	return
}

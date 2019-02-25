package btcvanity

type worker struct {
	generator *Generator
	result    chan IWallet
	err       chan error
}

// Work generates bitcoin wallet and pushes through channel
func (w *worker) Work(result chan IWallet, erri chan error) {
	wallet, err := w.generator.Generate()
	if err != nil {
		erri <- err
		return
	}
	result <- wallet
}

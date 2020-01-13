package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MarinX/btc-vanity"
)

var buffer = flag.Int("threads", 16, "How many threads you want to spawn")
var testnet = flag.Bool("testnet", false, "Use testnet")
var usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] pattern\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Example: %s Kid\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

	cfg := &btcvanity.Config{
		Buffer:  *buffer,
		TestNet: *testnet,
	}

	btc := btcvanity.New(cfg)

	fmt.Fprintf(os.Stdout,
		"Testnet: %t\nThreads: %d\nPattern: %s\nWorking...please wait\n",
		*testnet, *buffer, flag.Arg(0),
	)

	address, err := btc.Find(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Fprintf(os.Stdout, "Public key\n%s\n", address.PublicKey())
	fmt.Fprintf(os.Stdout, "Private key\n%s\n", address.PrivateKey())

}

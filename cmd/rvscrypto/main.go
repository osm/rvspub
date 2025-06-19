package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/osm/rvspub/internal/crypto"
)

func main() {
	decrypt := flag.Bool("d", false, "Decrypt instead of encrypt")
	version := flag.Uint("v", 2, "Crypto version")
	key := flag.String("k", "Cproc%2u", "8-byte encryption key (required)")
	flag.Parse()

	if len(*key) != crypto.KeySize {
		fmt.Fprintln(os.Stderr, "Error: -k key must be exactly 8 bytes")
		os.Exit(1)
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read stdin:", err)
		os.Exit(1)
	}

	var output []byte
	switch {
	case *decrypt:
		output, err = crypto.Decrypt(crypto.Version(*version-1), []byte(*key), data)
	default:
		output, err = crypto.Encrypt(crypto.Version(*version-1), []byte(*key), data)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Crypto error:", err)
		os.Exit(1)
	}

	os.Stdout.Write(output)
}

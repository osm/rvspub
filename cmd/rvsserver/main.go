package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/osm/rvspub/internal/crypto"
)

func main() {
	encryptionKey := flag.String("encryption-key", "Cproc%2u", "8-byte encryption key (required)")
	listenAddr := flag.String("listen-addr", "0.0.0.0:80", "Listen address")
	message := flag.String("message", "Remote Virtual Spy", "Message sent to clients")
	rvsServerAddr := flag.String("rvs-server-addr", "", "RVS server address")
	flag.Parse()

	if len(*encryptionKey) != crypto.KeySize {
		fmt.Fprintln(os.Stderr, "Error: -encryption-key key must be exactly 8 bytes")
		os.Exit(1)
	}

	if *rvsServerAddr == "" {
		fmt.Fprintln(os.Stderr, "Error: -rvs-server-addr can't be empty")
		os.Exit(1)
	}

	encryptedAddr, err := crypto.Encrypt(crypto.V3, []byte(*encryptionKey), []byte(*rvsServerAddr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to encrypt RVS server address: %v\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/rvssystem/test.txt",
		loggingHandler(requestHandler(fmt.Sprintf("RVS SERVER LIST\n%s\n", encryptedAddr))))

	http.HandleFunc("/rvssystem/msg.txt",
		loggingHandler(requestHandler(fmt.Sprintf("MSGVER3 '%s'", *message))))

	fmt.Printf("Server is listening on %s\n", *listenAddr)
	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to start server: %v\n", err)
		os.Exit(1)
	}
}

func loggingHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		h(w, r)
	}
}

func requestHandler(payload string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		if _, err := io.WriteString(w, payload); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write data: %v\n", err)
		}
	}
}

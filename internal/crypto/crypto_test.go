package crypto_test

import (
	"bytes"
	"testing"

	"github.com/osm/rvspub/internal/crypto"
)

var key = []byte("Cproc%2u")

func TestCrypto(t *testing.T) {
	type testCase struct {
		name    string
		version crypto.Version
		input   []byte
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "Encrypt and decrypt V1",
			version: crypto.V1,
			input:   []byte("RVHK_test_data_v1"),
		},
		{
			name:    "Encrypt and decrypt V2",
			version: crypto.V2,
			input:   []byte("RVHK_test_data_v2"),
		},
		{
			name:    "Unsupported version",
			version: crypto.Version(99),
			input:   []byte("test"),
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ciphertext, err := crypto.Encrypt(tc.version, key, tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected encrypt error: %v", err)
			}

			plaintext, err := crypto.Decrypt(tc.version, key, ciphertext)
			if err != nil {
				t.Fatalf("unexpected decrypt error: %v", err)
			}

			if !bytes.Equal(plaintext, tc.input) {
				t.Errorf("decrypted data does not match original.\n got: %q\nwant: %q", plaintext, tc.input)
			}
		})
	}
}

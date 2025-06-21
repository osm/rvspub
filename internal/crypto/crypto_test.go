package crypto_test

import (
	"bytes"
	"fmt"
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
			name:    "Encrypt and decrypt V3",
			version: crypto.V3,
			input:   []byte("RVHK_test_data_v3"),
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

var v3Samples = []struct {
	plaintext  []byte
	ciphertext []byte
}{
	{[]byte("GET %s/%s HTTP/1.0\nConnection: Keep-Alive\nAccept: text/html\n\n"), []byte("AEDFCGEPEGFGBNFADAFADKDLDHHFBNEEGNEAHICMAMELFMBACAAEBLAAANBPBCDOCGBFACECCCEJFLADCGHKDDAMAAEAECABHJFAAGAKBLFBBNBNDHBNBOGFGJ")},
	{[]byte("RVS SERVER LIST"), []byte("BBCGCBEPDAGAGACDAGCCFCCDCKHGGG")},
	{[]byte("test.txt"), []byte("DHBFABBLENFBEKAB")},
	{[]byte("200 OK"), []byte("HBEAECEPCMGO")},
	{[]byte("209.1.224.18"), []byte("HBEAELEBFCALAAEHHHFOEDFH")},
	{[]byte("216.234.161.84"), []byte("HBEBEEEBFBBGAGFLHCEGEDEBFLBB")},
	{[]byte("206.253.222.119"), []byte("HBEAEEEBFBBAABFLHBECEAEBFCBEAL")},
	{[]byte("209.90.125.196"), []byte("HBEAELEBFKBFBMEEHBEFFMFOFKBD")},
	{[]byte("/rvssystem"), []byte("GMACAEBMBAFMEBABCGBN")},
	{[]byte("GET %s/%s HTTP/1.0\nAccept: text/html\n\n"), []byte("AEDFCGEPEGFGBNFADAFADKDLDHHFBNEEGNEAHICOAAEGFHAFDHEKFCBLAGFNEGFKCLAEBPADGJCP")},
	{[]byte("/~rvs"), []byte("GMAOAABJBA")},
	{[]byte("/rvs"), []byte("GMACAEBM")},
	{[]byte("msg.txt"), []byte("COADBFEBBHFNEG")},
	{[]byte("www.geocities.com"), []byte("DEAHAFEBAEEAFNBGCKAEBLAKBAALFBBKCO")},
	{[]byte("server28.hypermart.net"), []byte("DABFAABJAGFHAAENGNBIALBPAGFHFPBEDBAEFMABAGFB")},
	{[]byte("lightning.prohosting.com"), []byte("CPBJBFAHBHELFLBLCEFOACBNAMENFNAGDHBJBMAIENEGFNBI")},
	{[]byte("uranus.spaceports.com"), []byte("DGACBDABBGFGBMAGDDBBBBAKBDEKEAABDAFOBBAAAO")},
}

func TestV3Samples(t *testing.T) {
	for i, sample := range v3Samples {
		t.Run(fmt.Sprintf("V3_Sample_%d", i), func(t *testing.T) {
			encrypted, err := crypto.Encrypt(crypto.V3, key, sample.plaintext)
			if err != nil {
				t.Fatalf("Encrypt failed: %v", err)
			}

			if !bytes.Equal(encrypted, sample.ciphertext) {
				t.Errorf("Encryption mismatch:\n")
				t.Errorf("Plaintext: %q\n", sample.plaintext)
				t.Errorf("Expected: %q\n", sample.ciphertext)
				t.Errorf("Got: %q\n", encrypted)
			}

			decrypted, err := crypto.Decrypt(crypto.V3, key, sample.ciphertext)
			if err != nil {
				t.Fatalf("Decrypt failed: %v", err)
			}

			if !bytes.Equal(decrypted, sample.plaintext) {
				t.Errorf("Decryption mismatch:\n")
				t.Errorf("Ciphertext: %q\n", sample.ciphertext)
				t.Errorf("Expected: %q\n", sample.plaintext)
				t.Errorf("Got: %q", decrypted)
			}

			roundTrip, err := crypto.Decrypt(crypto.V3, key, encrypted)
			if err != nil {
				t.Fatalf("Round-trip decrypt failed: %v", err)
			}

			if !bytes.Equal(roundTrip, sample.plaintext) {
				t.Errorf("Round-trip failed:\n")
				t.Errorf("Original: %q\n", sample.plaintext)
				t.Errorf("Round-trip: %q", roundTrip)
			}
		})
	}
}

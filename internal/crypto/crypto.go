package crypto

import (
	"errors"
)

type Version uint8

const (
	V1 Version = iota
	V2
	V3
)

const KeySize = 8

var (
	ErrInvalidKeyLength   = errors.New("expecting key to be eight bytes")
	ErrUnsupportedVersion = errors.New("unsupported crypto version")
)

func Encrypt(v Version, key, plaintext []byte) ([]byte, error) {
	if len(key) != KeySize {
		return nil, ErrInvalidKeyLength
	}

	switch v {
	case V1:
		buf := append([]byte(nil), plaintext...)
		encryptV1(key, buf)
		return buf, nil
	case V2:
		buf := append([]byte(nil), plaintext...)
		encryptV2(key, buf)
		return buf, nil
	case V3:
		buf := make([]byte, len(plaintext)*2)
		copy(buf[:len(plaintext)], plaintext)
		encryptV3(key, buf, len(plaintext)) // Pass original length
		return buf, nil
	}

	return nil, ErrUnsupportedVersion
}

func Decrypt(v Version, key, ciphertext []byte) ([]byte, error) {
	if len(key) != KeySize {
		return nil, ErrInvalidKeyLength
	}

	switch v {
	case V1:
		buf := append([]byte(nil), ciphertext...)
		decryptV1(key, buf)
		return buf, nil
	case V2:
		buf := append([]byte(nil), ciphertext...)
		decryptV2(key, buf)
		return buf, nil
	case V3:
		buf := append([]byte(nil), ciphertext...)
		decryptV3(key, buf)
		return buf[:len(buf)/2], nil
	}

	return nil, ErrUnsupportedVersion
}

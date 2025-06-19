package crypto

import (
	"errors"
)

type Version uint8

const (
	V1 Version = iota
	V2
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

	buf := append([]byte(nil), plaintext...)

	switch v {
	case V1:
		encryptV1(key, buf)
	case V2:
		encryptV2(key, buf)
	default:
		return nil, ErrUnsupportedVersion
	}

	return buf, nil
}

func Decrypt(v Version, key, ciphertext []byte) ([]byte, error) {
	if len(key) != KeySize {
		return nil, ErrInvalidKeyLength
	}

	buf := append([]byte(nil), ciphertext...)

	switch v {
	case V1:
		decryptV1(key, buf)
	case V2:
		decryptV2(key, buf)
	default:
		return nil, ErrUnsupportedVersion
	}

	return buf, nil
}

func encryptV1(key, buf []byte) {
	xorWithKeyAndIndex(key, buf)
}

func decryptV1(key, buf []byte) {
	xorWithKeyAndIndex(key, buf)
}

func encryptV2(key, buf []byte) {
	xorWithKey(key, buf)
	swapByteNibbles(buf)
	xorWithIndexStar(buf)
}

func decryptV2(key, buf []byte) {
	xorWithIndexStar(buf)
	swapByteNibbles(buf)
	xorWithKey(key, buf)
}

func xorWithKeyAndIndex(key, buf []byte) {
	for i := range buf {
		buf[i] ^= key[i%8] ^ byte(i)
	}
}

func xorWithIndexStar(buf []byte) {
	for i := range buf {
		buf[i] ^= byte((i * 0x2a) & 0xff)
	}
}

func swapByteNibbles(buf []byte) {
	for i := range buf {
		b := buf[i]
		buf[i] = ((b >> 4) | (b << 4)) & 0xff
	}
}

func xorWithKey(key, buf []byte) {
	for i := range buf {
		buf[i] ^= key[i%len(key)]
	}
}

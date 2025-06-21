package crypto

func encryptV1(key, buf []byte) {
	xorWithKeyAndIndex(key, buf)
}

func decryptV1(key, buf []byte) {
	xorWithKeyAndIndex(key, buf)
}

func xorWithKeyAndIndex(key, buf []byte) {
	for i := range buf {
		buf[i] ^= key[i%8] ^ byte(i)
	}
}

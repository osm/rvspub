package crypto

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

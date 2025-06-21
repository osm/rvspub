package crypto

func encryptV3(key, buf []byte, originalLen int) {
	xorWithKey(key, buf[:originalLen])
	bytesToHex(buf, originalLen)
}

func decryptV3(key, buf []byte) {
	hexToBytes(buf)
	xorWithKey(key, buf[:len(buf)/2])
}

func bytesToHex(buf []byte, originalLen int) {
	for i := originalLen - 1; i >= 0; i-- {
		b := buf[i]
		highNibble := (b >> 4) & 0x0f
		lowNibble := b & 0x0f
		buf[i*2] = byte(highNibble + 'A')
		buf[i*2+1] = byte(lowNibble + 'A')
	}
}

func hexToBytes(buf []byte) {
	for i := 0; i+1 < len(buf); i += 2 {
		c1 := int(buf[i] - 'A')
		c2 := int(buf[i+1] - 'A')
		combined := (c1 << 4) | c2
		buf[i/2] = byte(combined)
	}
}

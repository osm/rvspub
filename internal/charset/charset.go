package charset

import "strings"

var special = [32]string{
	"", "", "", "", "", ".", "", "",
	"", "", "", "", "", ">", ".", ".",
	"[", "]", "0", "1", "2", "3", "4", "5",
	"6", "7", "8", "9", ".", "<", "=", ">",
}

func Parse(input []byte) string {
	n := 0
	for _, c := range input {
		if c == 0 {
			break
		}
		n++
	}

	var result strings.Builder
	for _, b := range input[:n] {
		b &= 0x7f

		if b == 0x7f {
			continue
		}

		if b < 32 {
			result.WriteString(special[b])
		} else {
			result.WriteByte(b)
		}
	}

	return result.String()
}

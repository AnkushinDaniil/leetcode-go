package addBinary

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	bytes := make([]byte, len(a)+1)
	i := 1
	var bit byte
	for ; i <= len(b); i++ {
		bit += (a[len(a)-i] - '0') + (b[len(b)-i] - '0')
		bytes[len(bytes)-i] = '0' + bit&1
		bit >>= 1
	}
	for ; i <= len(a); i++ {
		bit += (a[len(a)-i] - '0')
		bytes[len(bytes)-i] = '0' + bit&1
		bit >>= 1
	}
	if bit == 1 {
		bytes[0] = '1'
	} else {
		bytes = bytes[1:]
	}
	return string(bytes)
}

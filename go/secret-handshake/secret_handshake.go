package secret

var secrets = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(handshakeCode uint) []string {
	handshake := make([]string, 0)
	for _, secret := range secrets {
		if handshakeCode&1 == 1 {
			handshake = append(handshake, secret)
		}
		handshakeCode = handshakeCode >> 1
	}

	if handshakeCode == 1 {
		for i, j := 0, len(handshake)-1; i < j; i, j = i+1, j-1 {
			handshake[i], handshake[j] = handshake[j], handshake[i]
		}
	}

	return handshake
}

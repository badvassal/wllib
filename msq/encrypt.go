package msq

// CalcChecksum calculates the checksum over an MSQ block's secure section.
// secSection must already be decrypted when passed to this function.
func CalcChecksum(secSection []byte) uint16 {
	csum := uint16(0)

	for _, b := range secSection {
		csum -= uint16(b)
	}

	return csum
}

// Encrypt encrypts an MSQ block's secure section.  xor0 and xor1 are the two
// checksum bytes in the block header.  See
// <https://wasteland.gamepedia.com/Encryption_Decryption>.
func Encrypt(plaintext []byte, xor0 byte, xor1 byte) []byte {
	cbytes := make([]byte, 0, len(plaintext))

	enc := xor0 ^ xor1
	for _, pb := range plaintext {
		cb := pb ^ enc
		enc += 0x1f

		cbytes = append(cbytes, cb)
	}

	return cbytes
}

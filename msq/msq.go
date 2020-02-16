package msq

const (
	BlockPrefix = "msq"
	HeaderLen   = 6
)

// Header is a decoded MSQ block header.  See
// <https://wasteland.gamepedia.com/Encryption_Decryption>.
type Header struct {
	GameIdx int // 0 or 1
	Xor0    byte
	Xor1    byte
}

// Body is a decoded and decrypted MSQ block body.  See
// <https://wasteland.gamepedia.com/MSQ_Block>.
type Body struct {
	SecSection   []byte
	PlainSection []byte
}

// Desc is a descriptor for an MSQ block.  It contains the block itself as well
// as some metadata obtained when the block was read.
type Desc struct {
	Offset int
	Hdr    Header
	Body   Body
}

// Clone performs a deep copy of a decoded MSQ block.
func (body *Body) Clone() *Body {
	dup := &Body{
		SecSection:   make([]byte, len(body.SecSection)),
		PlainSection: make([]byte, len(body.PlainSection)),
	}

	copy(dup.SecSection, body.SecSection)
	copy(dup.PlainSection, body.PlainSection)

	return dup
}

// EncodeMsqHeader encodes an MSQ block header to a byte sequence.
func EncodeMsqHeader(hdr Header) []byte {
	var out []byte

	out = append(out, []byte(BlockPrefix)...)
	out = append(out, byte(hdr.GameIdx+0x30))
	out = append(out, hdr.Xor0)
	out = append(out, hdr.Xor1)

	return out
}

// EncodeMsqBlock encodes an MSQ body to a byte sequence.
func EncodeMsqBlock(body Body, gameIdx int) []byte {
	csum := CalcChecksum(body.SecSection)

	hdr := Header{
		GameIdx: gameIdx,
		Xor0:    byte(csum & 0xff),
		Xor1:    byte(csum >> 8),
	}

	hdrBytes := EncodeMsqHeader(hdr)
	encBytes := Encrypt(body.SecSection, hdr.Xor0, hdr.Xor1)

	out := append(hdrBytes, encBytes...)
	out = append(out, body.PlainSection...)

	return out
}

package msq

const (
	BlockPrefix = "msq"
	HeaderLen   = 6
)

// A decoded MSQ block header.  See
// <https://wasteland.gamepedia.com/Encryption_Decryption>.
type Header struct {
	GameIdx byte
	Xor0    byte
	Xor1    byte
}

// A decoded and decrypted MSQ block.  See
// <https://wasteland.gamepedia.com/MSQ_Block>.
// XXX: Remove offset and header?  These two fields can get out of sync with
// the two "section" fields.
type Block struct {
	Offset       int
	Hdr          Header
	EncSection   []byte
	PlainSection []byte
}

// Clone performs a deep copy of a decoded MSQ block.
func (block *Block) Clone() *Block {
	dup := &Block{
		Offset:       block.Offset,
		Hdr:          block.Hdr,
		EncSection:   make([]byte, len(block.EncSection)),
		PlainSection: make([]byte, len(block.PlainSection)),
	}

	copy(dup.EncSection, block.EncSection)
	copy(dup.PlainSection, block.PlainSection)

	return dup
}

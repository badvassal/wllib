package decode

import (
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
	log "github.com/sirupsen/logrus"
)

// CentralDirLen is the size, in bytes, of the character table at the start of
// a strings area.
const StringsCharacterTableLen = 60

// StringsArea contains most of the text for a single MSQ block.  See
// <https://wasteland.gamepedia.com/String_Compressor>.
// XXX: This structure should contain only decompressed strings.  The character
// table and pointers should be derived when the area gets re-encoded.  There
// are currently some issues with string decoding, so for now we store the
// encoded contents.
type StringsArea struct {
	CharTable  []byte
	Pointers   []int
	StringData []byte
}

// DecodeCentralDir parses a stringsd area from a sequence of bytes.  It
// returns the decoded strings area and the size of the area, in bytes.
func DecodeStringsArea(data []byte) (*StringsArea, int, error) {
	onErr := wlerr.MakeWrapper("failed to decode strings area")

	chars, err := gen.ExtractBlob(data, 0, StringsCharacterTableLen)
	if err != nil {
		return nil, 0, onErr(err, "failed to extract character table")
	}

	off := StringsCharacterTableLen
	readPtr := func() int {
		p, err := gen.ReadUint16(data[off : off+2])
		if err != nil {
			panic(err.Error())
		}
		off += 2
		return p
	}

	// Read the first pointer to determine the offset of the first string.
	// From this, we can derive the end of the pointer list.
	p := readPtr()
	ptrAreaSize := p
	if ptrAreaSize > len(data)-off {
		return nil, 0, onErr(nil,
			"pointer area truncated: p1=%d off=%d len(data)=%d",
			p, off, len(data))
	}
	if ptrAreaSize%2 != 0 {
		return nil, 0, onErr(nil,
			"pointer region has invalid length: have=%d want%%2",
			ptrAreaSize)
	}

	pointers := make([]int, ptrAreaSize/2)
	pointers[0] = p
	for i := 1; i < len(pointers); i++ {
		pointers[i] = readPtr()
	}

	// Sometimes the final pointer is garbage for some reason?
	// XXX: Just exclude the final string group for now.
	if len(pointers) > 0 {
		pointers = pointers[:len(pointers)-1]
	}

	dataStart := StringsCharacterTableLen + pointers[0]
	// XXX: We don't know where the strings area ends.  Just assume the last
	// string is 10 bytes long for now.
	dataEnd := StringsCharacterTableLen + pointers[len(pointers)-1] + 10
	log.Debugf("decoding strings area: start=%d end=%d len(data)=%d",
		dataStart, dataEnd, len(data))
	stringData, err := gen.ExtractBlob(data, dataStart, dataEnd)
	if err != nil {
		return nil, 0, onErr(err, "")
	}

	return &StringsArea{
		CharTable:  chars,
		Pointers:   pointers,
		StringData: stringData,
	}, dataEnd, nil
}

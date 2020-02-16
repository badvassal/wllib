package msq

import (
	"bytes"
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	invRawMsg = "invalid raw MSQ block"
)

var stringPrefix = []byte{0x20, 0x65}

type readState int

const (
	readSecSection readState = iota
	readPlainSection
)

type reader struct {
	// Constant:
	hdr       Header
	src       []byte
	startOff  int
	isMap     bool
	finalCsum uint16

	// Variable
	off          int
	sec          byte
	csum         uint16
	state        readState
	secSection   []byte
	plainSection []byte
}

// newReader constructs a reader for the specified GAMEx file contents.
func newReader(hdr Header, game []byte, startOff int, isMap bool) *reader {
	return &reader{
		hdr:       hdr,
		src:       game,
		startOff:  startOff,
		isMap:     isMap,
		finalCsum: uint16(hdr.Xor1)<<8 + uint16(hdr.Xor0),

		off:   startOff,
		sec:   hdr.Xor0 ^ hdr.Xor1,
		state: readSecSection,
	}
}

// bytesFollow indicates whether a given string of bytes follows the current
// reader offset.
func (r *reader) bytesFollow(b []byte) bool {
	sub := r.src[r.off:]

	if len(sub) < len(b) {
		return false
	}

	mine := sub[:len(b)]
	return bytes.Compare(mine, b) == 0
}

// isDone indicates whether the MSQ block has been fully read.
func (r *reader) isDone() bool {
	if r.off == len(r.src) {
		// End of file.
		return true
	}

	return r.bytesFollow([]byte(BlockPrefix))
}

func (r *reader) readByte() (bool, error) {
	if r.off >= len(r.src) {
		return false, fmt.Errorf(
			"%s: truncated: start-off=%d", invRawMsg, r.startOff)
	}

	rawbyte := r.src[r.off]

	if r.state == readSecSection {
		pbyte := rawbyte ^ r.sec
		newCsum := r.csum - uint16(pbyte)

		atEnd := func() bool {
			// If our current checksum doesn't match the value indicated in the
			// header then there is more data to decrypt.
			if r.csum != r.finalCsum {
				return false
			}

			if r.isMap {
				// This is a bit of a hack.  The secure section ends when
				// the checksum equals the value specified in the MSQ header.
				// This leads to false positives.  To ignore false positives,
				// ensure the first plaintext section (strings section)
				// immediately follows.
				return r.bytesFollow(stringPrefix)
			} else {
				// We don't know what follows the secure section in non-map
				// blocks.  Just rely on the checksum.
				return newCsum != r.finalCsum
			}
		}

		if atEnd() {
			log.Debugf("done reading secure map data")
			r.state = readPlainSection
		} else {
			r.secSection = append(r.secSection, pbyte)
			r.sec += 0x1f
			r.csum = newCsum
		}
	}

	if r.state == readPlainSection {
		r.plainSection = append(r.plainSection, rawbyte)
	}

	r.off++

	if r.isDone() {
		return true, nil
	}

	return false, nil
}

func parseBody(game []byte, startOff int, hdr Header,
	isMap bool) (*Body, error) {

	r := newReader(hdr, game, startOff, isMap)

	for {
		done, err := r.readByte()
		if err != nil {
			return nil, err
		}

		if done {
			break
		}
	}

	return &Body{
		SecSection:   r.secSection,
		PlainSection: r.plainSection,
	}, nil
}

func parseHeader(game []byte, startOff int) (*Header, error) {
	rem := len(game) - startOff
	if rem < HeaderLen {
		return nil, fmt.Errorf(
			"%s: too few bytes: off=%d: have=%d want>=%d",
			invRawMsg, startOff, rem, HeaderLen)
	}
	sub := game[startOff:]

	prefix := sub[0:3]
	if string(prefix) != BlockPrefix {
		return nil, fmt.Errorf(
			"%s: invalid 3-byte prefix: off=%d (0x%x): want=\"%s\" have=%v",
			invRawMsg, startOff, startOff, BlockPrefix, prefix)
	}

	var idx int
	idxByte := sub[3]

	switch rune(idxByte) {
	case '0':
		idx = 0

	case '1':
		idx = 1

	default:
		return nil, fmt.Errorf(
			"%s: invalid sub idx: have=0x%02x want='0' or '1'",
			invRawMsg, idxByte)
	}

	hdr := &Header{
		GameIdx: idx,
		Xor0:    sub[4],
		Xor1:    sub[5],
	}

	log.Debugf("read header: %#v\n", hdr)

	return hdr, nil
}

// parseBlock decodes a single MSQ block.  isMap indicates whether the block to
// be parsed is a map block (as opposed to e.g. character data).
func parseBlock(game []byte, startOff int, isMap bool) (*Desc, error) {
	hdr, err := parseHeader(game, startOff)
	if err != nil {
		return nil, err
	}

	body, err := parseBody(game, startOff+HeaderLen, *hdr, isMap)
	if err != nil {
		return nil, err
	}

	return &Desc{
		Offset: startOff,
		Hdr:    *hdr,
		Body:   *body,
	}, nil
}

// ParseGame converts the contents of a GAMEx file into a set of decrypted MSQ
// blocks.  numMapBlocks is the number of blocks in the input which represent
// maps (as opposed to e.g. character data).
func ParseGame(game []byte, numMapBlocks int) ([]Desc, error) {
	var descs []Desc
	for off := 0; off < len(game); {
		isMap := len(descs) < numMapBlocks

		log.Debugf("parsing block %d at offset %d\n", len(descs), off)
		desc, err := parseBlock(game, off, isMap)
		if err != nil {
			return nil, err
		}
		descs = append(descs, *desc)

		off += HeaderLen +
			len(desc.Body.SecSection) +
			len(desc.Body.PlainSection)
	}

	return descs, nil
}

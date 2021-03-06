package action

import (
	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
)

const (
	transitionMinLen = 5
	transitionMaxLen = 6

	// If a transition's ToClass field is greater than or equal to
	// transitionToClassNoneMin, the ToSelector field is not present.
	transitionToClassNoneMin = 0xfd
)

// Transition represents a transition (teleport action) in an MSQ block.
type Transition struct {
	Relative   bool // B0b7
	Prompt     bool // B0b6
	StringPtr  int  // B0b0,5
	LocX       int  // B1
	LocY       int  // B2
	Location   int  // B3b0,7
	ToClass    int  // B4b0,3
	ToSelector int  // B4b4,7|B5b0,3 (optional)
}

// DecodeTransition decodes a transition from a sequence of bytes.  It returns
// the decoded transition and its length (in bytes).
func DecodeTransition(data []byte) (*Transition, int, error) {
	wrapErr := wlerr.MakeWrapper("failed to decode action transition")

	if len(data) < transitionMinLen {
		return nil, 0, wrapErr(nil,
			"data length too short: have=%d want>=%d: data=%+v",
			len(data), transitionMinLen, data)
	}

	at := &Transition{}
	off := 0

	at.Relative = data[off]&0x80 != 0
	at.Prompt = data[off]&0x40 != 0
	at.StringPtr = int(data[off] & 0x3f)
	off++

	at.LocX = int(data[off])
	off++

	at.LocY = int(data[off])
	off++

	at.Location = int(data[off])
	off++

	at.ToClass = int(data[off])
	off++

	if at.ToClass < transitionToClassNoneMin {
		if len(data) < transitionMaxLen {
			return nil, 0, wrapErr(nil,
				"data length too short: have=%d want>=%d: data=%+v",
				len(data), transitionMaxLen, data)
		}

		at.ToSelector = int(data[off])
		off++
	}

	return at, off, nil
}

// DecodeTransitionTables decodes a set of transitions from a table of byte
// buffers.
func DecodeTransitionTable(table gen.Table) ([]*Transition, error) {
	var ts []*Transition

	for i, elem := range table.Elems {
		if len(elem) == 0 {
			ts = append(ts, nil)
		} else {
			t, _, err := DecodeTransition(elem)
			if err != nil {
				return nil, wlerr.Wrapf(err, "transidx=%d", i)
			}
			ts = append(ts, t)
		}
	}

	return ts, nil
}

// MakeAbsolute converts a relative transition to an absolute one.  absCoords
// are the destination coordinates of the absolute transition.
func (t *Transition) MakeAbsolute(absCoords gen.Point) {
	t.Relative = false
	t.LocX = absCoords.X
	t.LocY = absCoords.Y
}

// IsDerelict indicates whether a transition leads to a derelict building.
func (t *Transition) IsDerelict() bool {
	return t.Location != defs.LocationPrevious && t.Location >= 128
}

// EncodeActionTransition encodes a single transition to a byte sequence.
func EncodeActionTransition(at Transition) []byte {
	var buf []byte

	b0 := byte(0)
	if at.Relative {
		b0 |= 0x80
	}
	if at.Prompt {
		b0 |= 0x40
	}
	b0 |= byte(at.StringPtr)
	buf = append(buf, b0)

	buf = append(buf, byte(at.LocX))

	buf = append(buf, byte(at.LocY))

	buf = append(buf, byte(at.Location))

	buf = append(buf, byte(at.ToClass))
	if at.ToClass < transitionToClassNoneMin {
		buf = append(buf, byte(at.ToSelector))
	}

	return buf
}

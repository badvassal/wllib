package action

import (
	"fmt"

	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
)

const (
	TransitionMinLen = 5
	TransitionMaxLen = 6

	// If a transition's ToClass field is greater than or equal to
	// TransitionToClassNoneMin, the ToSelector field is not present.
	TransitionToClassNoneMin = 0xfd
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
	wrapErr := func(err error, format string, args ...interface{}) error {
		return wlerr.Wrapf(err, "failed to decode action transition: %s",
			fmt.Sprintf(format, args...))
	}

	if len(data) < TransitionMinLen {
		return nil, 0, wrapErr(nil,
			"data length too short: have=%d want>=%d: data=%+v",
			len(data), TransitionMinLen, data)
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

	if at.ToClass < TransitionToClassNoneMin {
		if len(data) < TransitionMaxLen {
			return nil, 0, wrapErr(nil,
				"data length too short: have=%d want>=%d: data=%+v",
				len(data), TransitionMaxLen, data)
		}

		at.ToSelector = int(data[off])
		off++
	}

	return at, off, nil
}

// DecodeTransition decodes a set of transitions from a sequence of bytes.
// baseOff is the offset of the start of the transition action table relative
// to the start of the MSQ block's encoded section.
func DecodeTransitions(table []byte, baseOff int) ([]*Transition, error) {
	subPtrs, firstPtr, err := gen.ReadPointers(table, baseOff)
	if err != nil {
		return nil, err
	}
	tdata := table[len(subPtrs)*2:]

	var ts []*Transition
	for i, p := range subPtrs {
		var t *Transition

		if p == 0 {
			t = nil
		} else if i < len(subPtrs)-1 && subPtrs[i+1] == p {
			t = nil
		} else {
			blob, err := gen.ExtractBlob(tdata, p-firstPtr, -1)
			if err != nil {
				return nil, wlerr.Wrapf(err,
					"failed to extract action transition data: "+
						"p=%d firstPtr=%d",
					p, firstPtr)
			}

			t, _, err = DecodeTransition(blob)
			if err != nil {
				return nil, err
			}
		}

		ts = append(ts, t)
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

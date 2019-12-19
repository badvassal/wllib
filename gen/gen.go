package gen

import (
	"fmt"
	"sort"

	"github.com/badvassal/wllib/gen/wlerr"
)

// Point represents 2d coordinates or a width,height pair.
type Point struct {
	X int
	Y int
}

// ExtractBlob copies a subsequence of bytes from a larger sequence.  off is
// the offset within section to start the copy.  end is the offset at which to
// end the copy, or -1 to copy until the end of section.
func ExtractBlob(section []byte, off int, end int) ([]byte, error) {
	if off < 0 {
		return nil, fmt.Errorf("invalid offset: %d", off)
	}

	if end < 0 {
		end = len(section)
	}

	size := end - off
	if size < 0 {
		return nil, wlerr.Errorf(
			"data has negative length: off=%d end=%d", off, end)
	}

	if size == 0 {
		return nil, nil
	}

	if end > len(section) {
		return nil, wlerr.Errorf(
			"data extends beyond end of section: data=%d section=%d",
			end, len(section))
	}

	return section[off:end], nil
}

// ReadString parses a null-terminated UTF8 string from a sequence of bytes.
func ReadString(b []byte) (string, error) {
	for i, c := range b {
		if c == 0 {
			return string(b[:i]), nil
		}
	}

	return "", wlerr.Errorf("failed to parse string: no null terminator")
}

// ReadString parses a little-endian uint16 from a sequence of bytes.
func ReadUint16(b []byte) (int, error) {
	if len(b) != 2 {
		return 0, wlerr.Errorf(
			"cannot decode uint16: wrong number of bytes: have=%d want=2",
			len(b))
	}

	return int(b[1])<<8 + int(b[0]), nil
}

// ReadString parses a little-endian uint24 from a sequence of bytes.
func ReadUint24(b []byte) (int, error) {
	if len(b) != 3 {
		return 0, wlerr.Errorf(
			"cannot decode uint24: wrong number of bytes: have=%d want=3",
			len(b))
	}

	return int(b[2])<<16 + int(b[1])<<8 + int(b[0]), nil
}

// ReadPointers reads the pointers at the start of a table.  These pointers are
// offsets of data contained by the table.  The offsets are relative to the
// start of the table baseOff is the offset of the start of the table relative
// to the start of the MSQ block's encoded section.  It returns the set of
// pointers and its size, in bytes.
func ReadPointers(data []byte, baseOff int) ([]int, int, error) {
	off := 0
	readPtr := func() (int, error) {
		b, err := ExtractBlob(data, off, off+2)
		if err != nil {
			return 0, err
		}
		p, err := ReadUint16(b)
		if err != nil {
			return 0, err
		}

		off += 2

		return p, nil
	}

	var firstPtr int
	var ptrAreaLen int
	var ptrs []int
	for {
		p, err := readPtr()
		if err != nil {
			return nil, 0, wlerr.ToWLError(err)
		}
		ptrs = append(ptrs, p)

		if p != 0 {
			firstPtr = p
			ptrAreaLen = p - baseOff
			break
		}
	}

	if ptrAreaLen < 0 {
		return nil, 0, wlerr.Errorf(
			"pointer area has invalid size: p=%d have=%d want>=0",
			firstPtr, ptrAreaLen)
	}
	if ptrAreaLen%2 != 0 {
		return nil, 0, wlerr.Errorf(
			"pointer area has invalid size: have=%d want%%2: ptrs=%+v baseOff=%d",
			ptrAreaLen, ptrs, baseOff)
	}
	numPtrs := ptrAreaLen / 2

	for i := len(ptrs); i < numPtrs; i++ {
		p, err := readPtr()
		if err != nil {
			return nil, 0, wlerr.ToWLError(err)
		}
		ptrs = append(ptrs, p)
	}

	return ptrs, firstPtr, nil
}

// WriteUint16 converts a uint16 to its byte representation (little endian).
func WriteUint16(u16 uint16) []byte {
	return []byte{
		byte(u16 & 0xff),
		byte(u16 >> 8),
	}
}

// WriteUint16 converts the lower 24 bits of an int to their byte
// representation (little endian).  If any higher bits are set they are
// ignored.
func WriteUint24(u24 int) []byte {
	return []byte{
		byte((u24 & 0xff) >> 0),
		byte((u24 & 0xff00) >> 8),
		byte((u24 & 0xff0000) >> 16),
	}
}

// BoolToByte converts a boolean to its byte representation
// (false=0x00, true=0x01).
func BoolToByte(b bool) byte {
	if b {
		return byte(1)
	} else {
		return byte(0)
	}
}

// SortedUniqueInts sorts a slice of ints and removes duplicates.
func SortedUniqueInts(vals []int) []int {
	m := map[int]struct{}{}

	for _, v := range vals {
		m[v] = struct{}{}
	}

	var s []int
	for k, _ := range m {
		s = append(s, k)
	}

	sort.Ints(s)

	return s
}

// NextInt searches for the "next" value from a slice of sorted unique
// integers.  The "next" value is the first value greater than cur.  If no such
// next value is present in the list, max is returned instead.
func NextInt(cur int, sortedUnique []int, max int) int {
	for i := 0; i < len(sortedUnique)-1; i++ {
		if sortedUnique[i] >= cur {
			return sortedUnique[i+1]
		}
	}

	return max
}

// Assert panics if the given condition is false.
func Assert(expr bool) {
	if !expr {
		panic("ASSERTION FAILED")
	}
}

func FilterIDs(numIDs int, shouldKeep func(id int) bool) []int {
	ids := make([]int, 0, numIDs)

	for i := 0; i < numIDs; i++ {
		if shouldKeep(i) {
			ids = append(ids, i)
		}
	}

	return ids
}

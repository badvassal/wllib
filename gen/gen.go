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

func ReadUint16(b []byte) (int, error) {
	if len(b) != 2 {
		return 0, wlerr.Errorf(
			"cannot decode pointer: wrong number of bytes: have=%d want=2",
			len(b))
	}

	return int(b[1])<<8 + int(b[0]), nil
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

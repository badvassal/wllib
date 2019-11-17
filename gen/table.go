package gen

import (
	"github.com/badvassal/wllib/gen/wlerr"
)

// An encoded table has the following structure:
// <pointer-1>
// [pointer-2...]
// <element-1>
// [element-2...]
//
// Each pointer specifies the offset of its corresponding element.  Offsets may
// be relative to the start of the table or relative to something else (it
// depends on the specific table).

// Table represents a decoded table.  It contains a sequence of byte-slice
// elements.
type Table struct {
	Elems [][]byte
}

// ParseTable parses a table from a sequence of bytes.  baseOff the value that
// should be subtracted from a pointer to yield its corresponding element
// relative to the start of the table.  For example, if the pointers are all
// relative to the start of the table, then a baseOff value of 0 is
// appropriate.
func ParseTable(data []byte, baseOff int) (*Table, error) {
	ptrs, _, err := ReadPointers(data, baseOff)
	if err != nil {
		return nil, err
	}

	// The first element starts immediately after the last pointer.
	dataStart := len(ptrs) * 2

	// Detect invalid pointers.
	for i, p := range ptrs {
		if ptrs[i] != 0 {
			if ptrs[i] < dataStart {
				return nil, wlerr.Errorf(
					"table pointer points before data: p=%d data-start=%d ptrs=%+v",
					p, baseOff+dataStart, ptrs)
			} else if ptrs[i] >= baseOff+len(data) {
				return nil, wlerr.Errorf(
					"table pointer points after data: p=%d data-end=%d ptrs=%+v",
					p, baseOff+len(data), ptrs)
			} else if i > 0 && p < ptrs[i-1] {
				return nil, wlerr.Errorf(
					"table contains unsorted pointer list: ptrs=%+v", ptrs)
			}
		}
	}

	// Read the elements.

	var elems [][]byte

	for i, p := range ptrs {
		if p == 0 {
			elems = append(elems, nil)
		} else {
			next := baseOff + len(data)
			for _, n := range ptrs[i+1:] {
				if n != 0 {
					next = n
					break
				}
			}

			elem, err := ExtractBlob(data, p-baseOff, next-baseOff)
			if err != nil {
				return nil, wlerr.Wrapf(err,
					"failed to extract element from table: "+
						"i=%d baseOff=%d len(data)=%d ptrs=%+v",
					i, baseOff, len(data), ptrs)
			}
			elems = append(elems, elem)
		}
	}

	return &Table{
		Elems: elems,
	}, nil
}

// Pointers calculates the set of pointers that would be written at the start
// of an encoded table.
func (t *Table) Pointers(baseOff int) []int {
	var ptrs []int

	cur := baseOff + len(t.Elems)*2
	for _, e := range t.Elems {
		var p int
		if len(e) == 0 {
			p = 0
		} else {
			p = cur
			cur += len(e)
		}
		ptrs = append(ptrs, p)
	}

	return ptrs
}

// Encode encodes a table into a byte sequence.  baseOff is the value that
// should be added to each pointer at the start of the table.  If the pointers
// should be relative to the start of the table, then a baseOff value of 0 is
// appropriate.
func (t *Table) Encode(baseOff int) []byte {
	var b []byte

	ptrs := t.Pointers(baseOff)
	for _, p := range ptrs {
		b = append(b, WriteUint16(uint16(p))...)
	}

	for _, e := range t.Elems {
		b = append(b, e...)
	}

	return b
}

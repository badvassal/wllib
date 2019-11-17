package wlstrings

import "fmt"

const (
	StringCodeCapital   = 0x1e
	StringCodeShiftChar = 0x1f

	StringShiftAmount = 0x1e
)

func toBits(data []byte) []bool {
	bits := make([]bool, 0, len(data)*8)

	for _, val := range data {
		for i := 0; i < 8; i++ {
			bit := val&(1<<uint(i)) != 0
			bits = append(bits, bit)
		}
	}

	return bits
}

func bitsTo5b(bits []bool) []int {
	var vals []int
	for off := 0; off+5 < len(bits); off += 5 {
		val := 0
		blob := bits[off : off+5]
		for i, b := range blob {
			if b {
				val += 1 << uint(i)
			}
		}

		vals = append(vals, val)
	}

	return vals
}

// DecompressStringGroup converts a character table and a set of compressed
// strings into ASCII text.
func DecompressStringGroup(charTable []byte, data []byte) ([]byte, error) {
	bits := toBits(data)
	fiveb := bitsTo5b(bits)

	var raw []byte

	var shift bool
	var capital bool

	for _, v := range fiveb {
		switch v {
		case StringCodeCapital:
			if capital {
				return nil, fmt.Errorf("two adjacent capital control codes")
			}
			capital = true

		case StringCodeShiftChar:
			if shift {
				return nil, fmt.Errorf("two adjacent shift control codes")
			}
			shift = true

		default:
			idx := v
			if shift {
				idx += StringShiftAmount
			}

			c := charTable[idx]
			if capital {
				c -= 0x20
			}
			raw = append(raw, c)

			shift = false
			capital = false
		}
	}

	return raw, nil
}

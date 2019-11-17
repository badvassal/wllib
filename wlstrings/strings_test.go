package wlstrings

import "testing"

func intsToBools(ints []int) []bool {
	bools := make([]bool, len(ints))

	for i, val := range ints {
		switch val {
		case 0:
			bools[i] = false
		case 1:
			bools[i] = true
		default:
			panic("cannot convert non0 non1 integer to bool")
		}
	}

	return bools
}

func testToBitsOnce(t *testing.T, data []byte, want []int) {
	have := toBits(data)
	wantBools := intsToBools(want)

	if len(have) != len(want) {
		t.Fatalf("incorrect bit count: have=%d want=%d", len(have), len(want))
	}

	for i, h := range have {
		w := wantBools[i]

		if h != w {
			t.Fatalf("incorrect bit at offset %d: have=%v want=%v", i, h, w)
		}
	}
}

func testTo5bOnce(t *testing.T, data []byte, want []int) {
	bits := toBits(data)
	have := bitsTo5b(bits)

	if len(have) != len(want) {
		t.Fatalf("incorrect bit count: have=%d want=%d", len(have), len(want))
	}

	for i, h := range have {
		w := want[i]

		if h != w {
			t.Fatalf("incorrect 5b at offset %d: have=%v want=%v", i, h, w)
		}
	}
}

func TestToBits(t *testing.T) {
	b := []byte{0x01, 0x02, 0x03, 0xff}
	testToBitsOnce(t, b, []int{
		1, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 0, 0, 0, 0, 0, 0,
		1, 1, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1, 1, 1,
	})
}

func TestTo5b(t *testing.T) {
	b := []byte{0x01, 0x02, 0x03, 0xff}
	testTo5bOnce(t, b, []int{
		0x01,
		0x10,
		0x00,
		0x06,
		0x10,
		0x1f,
	})
}

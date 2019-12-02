package decode

import (
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
)

// MapData represents the effect each tile has when the player steps on it.
// See <https://wasteland.gamepedia.com/Map_Tile_Action_Classes>.
type MapData struct {
	ActionClasses   [][]int // [y][x] (0-15)
	ActionSelectors [][]int // [y][x] (0-255)
}

// MapDataLen calculates the size, in bytes, of a single block's map data.  dim
// is the map dimensions.
func MapDataLen(dim gen.Point) int {
	return dim.X * dim.Y * 3 / 2
}

// DecodeMapData decodes map data from a sequence of bytes.  dim is the map
// dimensions.
func DecodeMapData(data []byte, dim gen.Point) (*MapData, error) {
	onErr := wlerr.MakeWrapper("failed to decode map data")

	if err := ValidateMapDim(dim); err != nil {
		return nil, onErr(err, "")
	}

	numTiles := dim.X * dim.Y
	reqLen := numTiles/2 + numTiles
	if len(data) < reqLen {
		return nil, onErr(nil,
			"map data truncated: have=%dB want>=%dB",
			len(data), reqLen)
	}

	off := 0

	actclass := make([][]int, dim.Y)
	for y := 0; y < dim.Y; y++ {
		actclass[y] = make([]int, dim.X)
		for x := 0; x < dim.X; x += 2 {
			b := data[off]
			off++

			actclass[y][x] = int(b & 0x0f)
			actclass[y][x+1] = int(b >> 4)
		}
	}

	actsel := make([][]int, dim.Y)
	for y := 0; y < dim.Y; y++ {
		actsel[y] = make([]int, dim.X)
		for x := 0; x < dim.X; x++ {
			b := data[off]
			off++

			actsel[y][x] = int(b)
		}
	}

	return &MapData{
		ActionClasses:   actclass,
		ActionSelectors: actsel,
	}, nil
}

// EncodeMapData encodes map data to a byte sequence.
func EncodeMapData(md MapData) []byte {
	dim := gen.Point{
		X: len(md.ActionClasses[0]),
		Y: len(md.ActionClasses),
	}

	b := make([]byte, MapDataLen(dim))

	off := 0
	for y := 0; y < dim.Y; y++ {
		for x := 0; x < dim.X; x++ {
			ac1 := md.ActionClasses[y][x]
			ac2 := md.ActionClasses[y][x+1]
			b[off] = byte(ac2)<<4 | byte(ac1)

			off++
		}
	}

	for y := 0; y < dim.Y; y++ {
		for x := 0; x < dim.X; x++ {
			b[off] = byte(md.ActionSelectors[y][x])
			off++
		}
	}

	gen.Assert(off == len(b))

	return b
}

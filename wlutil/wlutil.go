package wlutil

import (
	"io/ioutil"

	"github.com/badvassal/wllib/decode"
	"github.com/badvassal/wllib/defs"
	"github.com/badvassal/wllib/gen"
	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wllib/modify"
	"github.com/badvassal/wllib/msq"
	"github.com/badvassal/wllib/serialize"
)

func GameIdxToFilename(gameIdx int) (string, error) {
	switch gameIdx {
	case 0:
		return "GAME1", nil
	case 1:
		return "GAME2", nil
	default:
		return "", wlerr.Errorf(
			"invalid game index: have=%d want=0or1", gameIdx)
	}
}

// ReadGames reads the WL.EXE file from disk and returns its contents.
func ReadExe(inDir string) ([]byte, error) {
	wl, err := ioutil.ReadFile(inDir + "/WL.EXE")
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to read WL.EXE file")
	}

	return wl, nil
}

// ReadGames reads the GAME1 and GAME2 files from disk and returns their
// contents.
func ReadGames(inDir string) ([]byte, []byte, error) {
	g0, err := ioutil.ReadFile(inDir + "/GAME1")
	if err != nil {
		return nil, nil, wlerr.Wrapf(err, "failed to read GAME1 file")
	}

	g1, err := ioutil.ReadFile(inDir + "/GAME2")
	if err != nil {
		return nil, nil, wlerr.Wrapf(err, "failed to read GAME2 file")
	}

	return g0, g1, nil
}

// ParseGames parses the contents of the GAME1 and GAME2 files into sequences
// of decrypted MSQ blocks.
func ParseGames(game0 []byte, game1 []byte) ([]msq.Desc, []msq.Desc, error) {
	d0, err := msq.ParseGame(game0, defs.Block0NumBlocks)
	if err != nil {
		return nil, nil, err
	}

	d1, err := msq.ParseGame(game1, defs.Block1NumBlocks)
	if err != nil {
		return nil, nil, err
	}

	return d0, d1, nil
}

// ReadAndParseGames reads the GAME1 and GAME2 files from disk and converts the
// contents to sequences of decrypted MSQ blocks.
func ReadAndParseGames(dir string) ([]msq.Desc, []msq.Desc, error) {
	g0, g1, err := ReadGames(dir)
	if err != nil {
		return nil, nil, err
	}

	d0, d1, err := ParseGames(g0, g1)
	if err != nil {
		return nil, nil, err
	}

	return d0, d1, nil
}

// DecodeGame decodes a sequence of decrypted MSQ blocks.
func DecodeGame(blocks []msq.Body, gameIdx int) ([]decode.Block, error) {
	var dbs []decode.Block

	for i, b := range blocks {
		db, err := decode.DecodeBlock(b, defs.MapDims[gameIdx][i])
		if err != nil {
			return nil, wlerr.Wrapf(err, "block=%d", i)
		}

		dbs = append(dbs, *db)
	}

	return dbs, nil
}

// CommitDecodeState writes the given decode state to a set of MSQ blocks.
// After decodede blocks are modified, the modifications are transferred to MSQ
// blocks with this function.
func CommitDecodeState(state decode.DecodeState,
	bodies1 []msq.Body, bodies2 []msq.Body) error {

	commitGame := func(dbs []decode.Block, bodies []msq.Body,
		dims []gen.Point) error {

		for i, db := range dbs {
			m := modify.NewBlockModifier(bodies[i], dims[i])
			if err := m.ReplaceActionTransitions(
				db.ActionTables.Transitions); err != nil {

				return wlerr.Wrapf(err, "block=%d", i)
			}

			bodies[i] = m.Body()
		}

		return nil
	}

	if err := commitGame(state.Blocks[0], bodies1, defs.MapDims[0]); err != nil {
		return err
	}

	if err := commitGame(state.Blocks[1], bodies2, defs.MapDims[1]); err != nil {
		return err
	}

	return nil
}

// DecodeGames converts a pair of MSQ block sequences (read from the GAME1 and
// GAME2 files) into a DecodeState.
func DecodeGames(bs1 []msq.Body,
	bs2 []msq.Body) (*decode.DecodeState, error) {

	dbs1, err := DecodeGame(bs1[:defs.Block0NumBlocks], 0)
	if err != nil {
		return nil, err
	}

	dbs2, err := DecodeGame(bs2[:defs.Block1NumBlocks], 1)
	if err != nil {
		return nil, err
	}

	return &decode.DecodeState{
		[][]decode.Block{
			dbs1,
			dbs2,
		},
	}, nil
}

func WriteGame(gameIdx int, data []byte, outDir string) error {
	filename, err := GameIdxToFilename(gameIdx)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(outDir+"/"+filename, data, 0644); err != nil {
		return wlerr.Wrapf(err, "failed to write game file")
	}

	return nil
}

// WriteGames writes the GAME1 and GAME2 files to disk.
func WriteGames(data0 []byte, data1 []byte, outDir string) error {
	if err := WriteGame(0, data0, outDir); err != nil {
		return err
	}

	if err := WriteGame(1, data1, outDir); err != nil {
		return err
	}

	return nil
}

// SerailizeAndWriteGames encrypts a pair of MSQ block sequences and writes
// them to disk as the GAME1 and GAME2.
func SerializeAndWriteGames(blocks0 []msq.Body, blocks1 []msq.Body,
	outDir string) error {

	g0 := serialize.SerializeGame(blocks0, 0)
	g1 := serialize.SerializeGame(blocks1, 1)

	if err := WriteGames(g0, g1, outDir); err != nil {
		return err
	}

	return nil
}

// DescsToBodies converts a slice of block descriptors to a slice of block
// bodies.
func DescsToBodies(descs []msq.Desc) []msq.Body {
	bodies := make([]msq.Body, len(descs))
	for i, d := range descs {
		bodies[i] = d.Body
	}

	return bodies
}

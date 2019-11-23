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

// ReadGames reads the GAME1 and GAME2 files from disk and returns their
// contents.
func ReadGames(inDir string) ([]byte, []byte, error) {
	g0, err := ioutil.ReadFile(inDir + "/GAME1")
	if err != nil {
		return nil, nil, wlerr.Wrapf(err, "failed to read GAME file")
	}

	g1, err := ioutil.ReadFile(inDir + "/GAME2")
	if err != nil {
		return nil, nil, wlerr.Wrapf(err, "failed to read GAME file")
	}

	return g0, g1, nil
}

// ParseGames parses the contents of the GAME1 and GAME2 files into sequences
// of decrypted MSQ blocks.
func ParseGames(game0 []byte, game1 []byte) ([]msq.Block, []msq.Block, error) {
	b0, err := msq.ParseGame(game0, defs.Block0NumBlocks)
	if err != nil {
		return nil, nil, err
	}

	b1, err := msq.ParseGame(game1, defs.Block1NumBlocks)
	if err != nil {
		return nil, nil, err
	}

	return b0, b1, nil
}

// ReadAndParseGames reads the GAME1 and GAME2 files from disk and converts the
// contents to sequences of decrypted MSQ blocks.
func ReadAndParseGames(dir string) ([]msq.Block, []msq.Block, error) {
	g0, g1, err := ReadGames(dir)
	if err != nil {
		return nil, nil, err
	}

	b0, b1, err := ParseGames(g0, g1)
	if err != nil {
		return nil, nil, err
	}

	return b0, b1, nil
}

// DecodeGame decodes a sequence of decrypted MSQ blocks.
func DecodeGame(blocks []msq.Block, gameIdx int) ([]decode.Block, error) {
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
	blocks1 []msq.Block, blocks2 []msq.Block) error {

	commitGame := func(dbs []decode.Block, blocks []msq.Block,
		dims []gen.Point) error {

		for i, db := range dbs {
			m := modify.NewBlockModifier(blocks[i], dims[i])
			if err := m.ReplaceActionTransitions(
				db.ActionTables.Transitions); err != nil {

				return wlerr.Wrapf(err, "block=%d", i)
			}

			blocks[i] = m.Block()
		}

		return nil
	}

	if err := commitGame(state.Blocks[0], blocks1, defs.MapDims[0]); err != nil {
		return err
	}

	if err := commitGame(state.Blocks[1], blocks2, defs.MapDims[1]); err != nil {
		return err
	}

	return nil
}

// DecodeGames converts a pair of MSQ block sequences (read from the GAME1 and
// GAME2 files) into a DecodeState.
func DecodeGames(bs1 []msq.Block,
	bs2 []msq.Block) (*decode.DecodeState, error) {

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

// WriteGames writes the GAME1 and GAME2 files to disk.
func WriteGames(data0 []byte, data1 []byte, outDir string) error {
	if err := ioutil.WriteFile(outDir+"/GAME1", data0, 0644); err != nil {
		return wlerr.Wrapf(err, "failed to write game file")
	}

	if err := ioutil.WriteFile(outDir+"/GAME2", data1, 0644); err != nil {
		return wlerr.Wrapf(err, "failed to write game file")
	}

	return nil
}

// SerailizeAndWriteGames encrypts a pair of MSQ block sequences and writes
// them to disk as the GAME1 and GAME2.
func SerializeAndWriteGames(blocks0 []msq.Block, blocks1 []msq.Block,
	outDir string) error {

	g0 := serialize.SerializeGame(blocks0)
	g1 := serialize.SerializeGame(blocks1)

	if err := WriteGames(g0, g1, outDir); err != nil {
		return err
	}

	return nil
}

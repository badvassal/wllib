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

// ReadOneGame reads a single GAMEx file from disk and converts it to a
// sequence of decrypted MSQ blocks.
func ReadOneGame(inPath string, numMapBlocks int) ([]msq.Block, error) {
	g, err := ioutil.ReadFile(inPath)
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to read GAME file")
	}

	blocks, err := msq.ParseGame(g, numMapBlocks)
	if err != nil {
		return nil, err
	}

	return blocks, nil
}

// ReadGames reads the GAME1 and GAME2 files and converts them into sequences
// of decrypted MSQ blocks.
func ReadGames(inDir string) ([]msq.Block, []msq.Block, error) {
	b1, err := ReadOneGame(inDir+"/GAME1", defs.Block0NumBlocks)
	if err != nil {
		return nil, nil, err
	}

	b2, err := ReadOneGame(inDir+"/GAME2", defs.Block1NumBlocks)
	if err != nil {
		return nil, nil, err
	}

	return b1, b2, nil
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
func DecodeGames(bs1 []msq.Block, bs2 []msq.Block) (*decode.DecodeState, error) {
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

// WriteOneGame encrypts and writes a sequence of decrypted MSQ blocks to disk.
func WriteOneGame(blocks []msq.Block, filename string) error {
	game := serialize.SerializeGame(blocks)

	if err := ioutil.WriteFile(filename, game, 0644); err != nil {
		return wlerr.Wrapf(err, "failed to write game file")
	}

	return nil
}

// WriteOneGame encrypts and writes a pair of decrypted MSQ block sequences to
// disk.
func WriteGames(blocks1 []msq.Block, blocks2 []msq.Block, outDir string) error {
	if err := WriteOneGame(blocks1, outDir+"/GAME1"); err != nil {
		return err
	}

	if err := WriteOneGame(blocks2, outDir+"/GAME2"); err != nil {
		return err
	}

	return nil
}

// Copyright (c) 2014-2014 PPCD developers.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ppcutil_test

import (
	"testing"

	"github.com/ppcsuite/ppcd/chaincfg"
	"github.com/ppcsuite/ppcd/database"
	_ "github.com/ppcsuite/ppcd/database/ldb" // init only
	"github.com/ppcsuite/ppcutil"
)

func TestPoWTargetCalculation(t *testing.T) {
	params := chaincfg.MainNetParams
	db, err := database.OpenDB("leveldb", "testdata/db_512")
	if err != nil {
		t.Errorf("db error %v", err)
		return
	}
	defer db.Close()

	lastBlock, _ := db.FetchBlockBySha(params.GenesisHash)
	for height := 1; height < 512; height++ {
		sha, _ := db.FetchBlockShaByHeight(int64(height))
		block, _ := db.FetchBlockBySha(sha)
		if !block.MsgBlock().IsProofOfStake() {
			targetRequired := ppcutil.GetNextTargetRequired(params, db, lastBlock, false)
			if targetRequired != block.MsgBlock().Header.Bits {
				t.Errorf("bad target for block #%d %v, have %x want %x", height, sha, targetRequired, block.MsgBlock().Header.Bits)
				return
			}
		}
		lastBlock = block
	}
	if lastBlock.Height() != 511 {
		t.Error("test ended too early")
	}
	return
}

func TestReadCBlockIndex(t *testing.T) {
	r := ppcutil.ReadCBlockIndex("testdata/blkindex.csv")
	if r.Height != 0 {
		t.Errorf("bad root height, have %d, want %d", r.Height, 0)
	}
	for r.Next != nil {
		r = r.Next
	}
	if r.Height != 131325 {
		t.Errorf("bad head height, have %d, want %d", r.Height, 131325)
	}
}

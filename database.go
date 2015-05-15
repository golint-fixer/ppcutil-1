// Copyright (c) 2014-2014 PPCD developers.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ppcutil

import (
	"github.com/ppcsuite/btcutil"
	"github.com/ppcsuite/ppcd/database"
)

// https://github.com/ppcoin/ppcoin/blob/v0.4.0ppc/src/main.cpp#L894
// ppc: find last block index up to pindex
func GetLastBlockIndex(db database.Db, last *btcutil.Block, proofOfStake bool) (block *btcutil.Block) {
	block = last
	for true {
		if block == nil {
			break
		}
		//TODO dirty workaround, ppcoin doesn't point to genesis block
		if block.Height() == 0 {
			return nil
		}
		prevExists, err := db.ExistsSha(&block.MsgBlock().Header.PrevBlock)
		if err != nil || !prevExists {
			break
		}
		if block.MsgBlock().IsProofOfStake() == proofOfStake {
			break
		}
		block, _ = db.FetchBlockBySha(&block.MsgBlock().Header.PrevBlock)
	}
	return block
}

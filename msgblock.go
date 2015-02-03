// Copyright (c) 2014-2014 PPCD developers.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ppcutil

import (
	"github.com/ppcsuite/btcwire"
)

// IsMsgBlockProofOfStake checks if MsgBlock is of proof of stake type
// https://github.com/ppcoin/ppcoin/blob/v0.4.0ppc/src/main.h#L962
// ppc: two types of block: proof-of-work or proof-of-stake
func IsMsgBlockProofOfStake(msg *btcwire.MsgBlock) bool {
	return len(msg.Transactions) > 1 &&
		msg.Transactions[1].IsCoinStake()
}

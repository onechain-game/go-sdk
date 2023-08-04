package types

import (
	ntypes "github.com/onechain-game/go-sdk/common/types"
	"github.com/onechain-game/go-sdk/types/tx"
	types "github.com/tendermint/tendermint/rpc/core/types"
)

func NewCodec() *amino.Codec {
	cdc := amino.NewCodec()
	types.RegisterAmino(cdc)
	ntypes.RegisterWire(cdc)
	tx.RegisterCodec(cdc)
	return cdc
}

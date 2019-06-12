package genutil

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/andrecronje/cosmos-sdk/codec"
	sdk "github.com/andrecronje/cosmos-sdk/types"
)

// initialize accounts and deliver genesis transactions
func InitGenesis(ctx sdk.Context, cdc *codec.Codec, stakingKeeper StakingKeeper,
	deliverTx deliverTxfn, genesisState GenesisState) []abci.ValidatorUpdate {

	var validators []abci.ValidatorUpdate
	if len(genesisState.GenTxs) > 0 {
		validators = DeliverGenTxs(ctx, cdc, genesisState.GenTxs, stakingKeeper, deliverTx)
	}
	return validators
}

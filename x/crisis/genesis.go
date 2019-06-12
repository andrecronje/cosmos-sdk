package crisis

import (
	sdk "github.com/andrecronje/cosmos-sdk/types"
	"github.com/andrecronje/cosmos-sdk/x/crisis/types"
)

// new crisis genesis
func InitGenesis(ctx sdk.Context, keeper Keeper, data types.GenesisState) {
	keeper.SetConstantFee(ctx, data.ConstantFee)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) types.GenesisState {
	constantFee := keeper.GetConstantFee(ctx)
	return types.NewGenesisState(constantFee)
}

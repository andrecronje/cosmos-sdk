package types

import (
	sdk "github.com/andrecronje/cosmos-sdk/types"
	"github.com/andrecronje/cosmos-sdk/x/params"
)

// Default parameter namespace
const (
	DefaultParamspace = ModuleName
)

var (
	// key for constant fee parameter
	ParamStoreKeyConstantFee = []byte("ConstantFee")
)

// type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyConstantFee, sdk.Coin{},
	)
}

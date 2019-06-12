package genutil

import (
	"github.com/andrecronje/cosmos-sdk/codec"
	sdk "github.com/andrecronje/cosmos-sdk/types"
	"github.com/andrecronje/cosmos-sdk/x/auth"
	"github.com/andrecronje/cosmos-sdk/x/staking"
)

// generic sealed codec to be used throughout this module
var moduleCdc *codec.Codec

// TODO abstract genesis transactions registration back to staking
// required for genesis transactions
func init() {
	moduleCdc = codec.New()
	staking.RegisterCodec(moduleCdc)
	auth.RegisterCodec(moduleCdc)
	sdk.RegisterCodec(moduleCdc)
	codec.RegisterCrypto(moduleCdc)
	moduleCdc.Seal()
}

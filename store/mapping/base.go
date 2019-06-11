package mapping

import (
	//	"github.com/tendermint/tendermint/crypto/merkle"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Base struct {
	cdc     *codec.Codec
	storefn func(Context) KVStore
	prefix  []byte
	//	keypath keypath // temporal
}

func EmptyBase() Base {
	return NewBase(nil, nil)
}

func NewBase(cdc *codec.Codec, key sdk.StoreKey) Base {
	return Base{
		cdc:     cdc,
		storefn: func(ctx Context) KVStore { return ctx.KVStore(key) },
	}
}

func NewBaseWithGetter(cdc *codec.Codec, storefn func(Context) KVStore) Base {
	return Base{
		cdc:     cdc,
		storefn: storefn,
	}
}

func (base Base) Store(ctx Context) KVStore {
	return prefix.NewStore(base.storefn(ctx), base.prefix)
}

func (base Base) Prefix(prefix []byte) (res Base) {
	res = Base{
		cdc:     base.cdc,
		storefn: base.storefn,
	}
	res.prefix = join(base.prefix, prefix)
	return
}

func (base Base) Cdc() *codec.Codec {
	return base.cdc
}

func (base Base) key(key []byte) []byte {
	return join(base.prefix, key)
}

func join(a, b []byte) (res []byte) {
	res = make([]byte, len(a)+len(b))
	copy(res, a)
	copy(res[len(a):], b)
	return
}

/*
func (base Base) KeyPath() merkle.KeyPath {
	if len(base.keypath.prefix) != 0 {
		return base.keypath.keypath.AppendKey(base.keypath.prefix, merkle.KeyEncodingHex)
	}
	return base.keypath.keypath
}
*/

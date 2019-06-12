package tags

import (
	sdk "github.com/andrecronje/cosmos-sdk/types"
)

// Tag keys and values
var (
	ActionUndelegateCoins = "undelegateCoins"
	ActionDelegateCoins   = "delegateCoins"
	TxCategory            = "bank"

	Action    = sdk.TagAction
	Category  = sdk.TagCategory
	Recipient = "recipient"
	Sender    = "sender"
)

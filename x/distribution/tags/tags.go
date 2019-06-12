package tags

import (
	sdk "github.com/andrecronje/cosmos-sdk/types"
)

// Distribution tx tags
var (
	Rewards    = "rewards"
	Commission = "commission"
	TxCategory = "distribution"

	Validator = sdk.TagSrcValidator
	Category  = sdk.TagCategory
	Sender    = sdk.TagSender
)

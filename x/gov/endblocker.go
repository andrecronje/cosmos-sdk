package gov

import (
	"fmt"

	sdk "github.com/andrecronje/cosmos-sdk/types"
	"github.com/andrecronje/cosmos-sdk/x/gov/tags"
)

// EndBlocker called every block, process inflation, update validator set
func EndBlocker(ctx sdk.Context, keeper Keeper) sdk.Tags {
	logger := keeper.Logger(ctx)
	resTags := sdk.NewTags()

	// delete inactive proposal from store and its deposits
	keeper.IterateInactiveProposalsQueue(ctx, ctx.BlockHeader().Time, func(proposal Proposal) bool {
		keeper.DeleteProposal(ctx, proposal.ProposalID)
		keeper.DeleteDeposits(ctx, proposal.ProposalID)

		resTags = resTags.AppendTag(tags.ProposalID, fmt.Sprintf("%d", proposal.ProposalID))
		resTags = resTags.AppendTag(tags.ProposalResult, tags.ActionProposalDropped)

		logger.Info(
			fmt.Sprintf("proposal %d (%s) didn't meet minimum deposit of %s (had only %s); deleted",
				proposal.ProposalID,
				proposal.GetTitle(),
				keeper.GetDepositParams(ctx).MinDeposit,
				proposal.TotalDeposit,
			),
		)
		return false
	})

	// fetch active proposals whose voting periods have ended (are passed the block time)
	keeper.IterateActiveProposalsQueue(ctx, ctx.BlockHeader().Time, func(proposal Proposal) bool {
		var tagValue, logMsg string

		passes, burnDeposits, tallyResults := tally(ctx, keeper, proposal)

		if burnDeposits {
			keeper.DeleteDeposits(ctx, proposal.ProposalID)
		} else {
			keeper.RefundDeposits(ctx, proposal.ProposalID)
		}

		if passes {
			handler := keeper.router.GetRoute(proposal.ProposalRoute())
			cacheCtx, writeCache := ctx.CacheContext()

			// The proposal handler may execute state mutating logic depending
			// on the proposal content. If the handler fails, no state mutation
			// is written and the error message is logged.
			err := handler(cacheCtx, proposal.Content)
			if err == nil {
				proposal.Status = StatusPassed
				tagValue = tags.ActionProposalPassed
				logMsg = "passed"

				// write state to the underlying multi-store
				writeCache()
			} else {
				proposal.Status = StatusFailed
				tagValue = tags.ActionProposalFailed
				logMsg = fmt.Sprintf("passed, but failed on execution: %s", err.ABCILog())
			}
		} else {
			proposal.Status = StatusRejected
			tagValue = tags.ActionProposalRejected
			logMsg = "rejected"
		}

		proposal.FinalTallyResult = tallyResults

		keeper.SetProposal(ctx, proposal)
		keeper.RemoveFromActiveProposalQueue(ctx, proposal.ProposalID, proposal.VotingEndTime)

		logger.Info(
			fmt.Sprintf(
				"proposal %d (%s) tallied; result: %s",
				proposal.ProposalID, proposal.GetTitle(), logMsg,
			),
		)

		resTags = resTags.AppendTag(tags.ProposalID, fmt.Sprintf("%d", proposal.ProposalID))
		resTags = resTags.AppendTag(tags.ProposalResult, tagValue)

		return false
	})

	return resTags
}

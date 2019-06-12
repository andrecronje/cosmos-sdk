package client

import (
	"github.com/andrecronje/cosmos-sdk/x/distribution/client/cli"
	"github.com/andrecronje/cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/andrecronje/cosmos-sdk/x/gov/client"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)

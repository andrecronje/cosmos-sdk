package rest

import (
	"net/http"

	"github.com/andrecronje/cosmos-sdk/client/context"
	clientrest "github.com/andrecronje/cosmos-sdk/client/rest"
	sdk "github.com/andrecronje/cosmos-sdk/types"
	"github.com/andrecronje/cosmos-sdk/types/rest"
	"github.com/andrecronje/cosmos-sdk/x/ibc"

	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/ibc/{destchain}/{address}/send", TransferRequestHandlerFn(cliCtx)).Methods("POST")
}

type transferReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Amount  sdk.Coins    `json:"amount"`
}

// TransferRequestHandler - http request handler to transfer coins to a address
// on a different chain via IBC.
func TransferRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		destChainID := vars["destchain"]
		bech32Addr := vars["address"]

		to, err := sdk.AccAddressFromBech32(bech32Addr)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var req transferReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		from, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		packet := ibc.NewIBCPacket(from, to, req.Amount, req.BaseReq.ChainID, destChainID)
		msg := ibc.MsgIBCTransfer{IBCPacket: packet}

		clientrest.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/client/context
// ALIASGEN: github.com/cosmos/cosmos-sdk/client/flags
// ALIASGEN: github.com/cosmos/cosmos-sdk/client/keys
// ALIASGEN: github.com/cosmos/cosmos-sdk/client/lcd
// ALIASGEN: github.com/cosmos/cosmos-sdk/client/rest
// ALIASGEN: github.com/cosmos/cosmos-sdk/client/rpc
// ALIASGEN: github.com/cosmos/cosmos-sdk/client/tx
// ALIASGEN: github.com/cosmos/cosmos-sdk/client/utils
// ALIASGEN: github.com/cosmos/cosmos-sdk/client/input
package client

import (
	"github.com/andrecronje/cosmos-sdk/client/context"
	"github.com/andrecronje/cosmos-sdk/client/flags"
	"github.com/andrecronje/cosmos-sdk/client/input"
	"github.com/andrecronje/cosmos-sdk/client/keys"
	"github.com/andrecronje/cosmos-sdk/client/lcd"
	"github.com/andrecronje/cosmos-sdk/client/rest"
	"github.com/andrecronje/cosmos-sdk/client/rpc"
	"github.com/andrecronje/cosmos-sdk/client/tx"
	"github.com/andrecronje/cosmos-sdk/client/utils"
)

const (
	DefaultGasAdjustment   = flags.DefaultGasAdjustment
	DefaultGasLimit        = flags.DefaultGasLimit
	GasFlagAuto            = flags.GasFlagAuto
	BroadcastBlock         = flags.BroadcastBlock
	BroadcastSync          = flags.BroadcastSync
	BroadcastAsync         = flags.BroadcastAsync
	FlagHome               = flags.FlagHome
	FlagUseLedger          = flags.FlagUseLedger
	FlagChainID            = flags.FlagChainID
	FlagNode               = flags.FlagNode
	FlagHeight             = flags.FlagHeight
	FlagGasAdjustment      = flags.FlagGasAdjustment
	FlagTrustNode          = flags.FlagTrustNode
	FlagFrom               = flags.FlagFrom
	FlagName               = flags.FlagName
	FlagAccountNumber      = flags.FlagAccountNumber
	FlagSequence           = flags.FlagSequence
	FlagMemo               = flags.FlagMemo
	FlagFees               = flags.FlagFees
	FlagGasPrices          = flags.FlagGasPrices
	FlagBroadcastMode      = flags.FlagBroadcastMode
	FlagPrintResponse      = flags.FlagPrintResponse
	FlagDryRun             = flags.FlagDryRun
	FlagGenerateOnly       = flags.FlagGenerateOnly
	FlagIndentResponse     = flags.FlagIndentResponse
	FlagListenAddr         = flags.FlagListenAddr
	FlagMaxOpenConnections = flags.FlagMaxOpenConnections
	FlagRPCReadTimeout     = flags.FlagRPCReadTimeout
	FlagRPCWriteTimeout    = flags.FlagRPCWriteTimeout
	FlagOutputDocument     = flags.FlagOutputDocument
	FlagSkipConfirmation   = flags.FlagSkipConfirmation
	DefaultKeyPass         = keys.DefaultKeyPass
	FlagAddress            = keys.FlagAddress
	FlagPublicKey          = keys.FlagPublicKey
	FlagBechPrefix         = keys.FlagBechPrefix
	FlagDevice             = keys.FlagDevice
	OutputFormatText       = keys.OutputFormatText
	OutputFormatJSON       = keys.OutputFormatJSON
	MinPassLength          = input.MinPassLength
)

var (
	// functions aliases
	NewCLIContextWithFrom              = context.NewCLIContextWithFrom
	NewCLIContext                      = context.NewCLIContext
	GetAccountDecoder                  = context.GetAccountDecoder
	GetFromFields                      = context.GetFromFields
	ErrInvalidAccount                  = context.ErrInvalidAccount
	ErrVerifyCommit                    = context.ErrVerifyCommit
	GetCommands                        = flags.GetCommands
	PostCommands                       = flags.PostCommands
	RegisterRestServerFlags            = flags.RegisterRestServerFlags
	ParseGas                           = flags.ParseGas
	NewCompletionCmd                   = flags.NewCompletionCmd
	MarshalJSON                        = keys.MarshalJSON
	UnmarshalJSON                      = keys.UnmarshalJSON
	Commands                           = keys.Commands
	NewAddNewKey                       = keys.NewAddNewKey
	NewRecoverKey                      = keys.NewRecoverKey
	NewUpdateKeyReq                    = keys.NewUpdateKeyReq
	NewDeleteKeyReq                    = keys.NewDeleteKeyReq
	GetKeyInfo                         = keys.GetKeyInfo
	GetPassphrase                      = keys.GetPassphrase
	ReadPassphraseFromStdin            = keys.ReadPassphraseFromStdin
	NewKeyBaseFromHomeFlag             = keys.NewKeyBaseFromHomeFlag
	NewKeyBaseFromDir                  = keys.NewKeyBaseFromDir
	NewInMemoryKeyBase                 = keys.NewInMemoryKeyBase
	NewRestServer                      = lcd.NewRestServer
	ServeCommand                       = lcd.ServeCommand
	WriteGenerateStdTxResponse         = rest.WriteGenerateStdTxResponse
	BlockCommand                       = rpc.BlockCommand
	GetChainHeight                     = rpc.GetChainHeight
	BlockRequestHandlerFn              = rpc.BlockRequestHandlerFn
	LatestBlockRequestHandlerFn        = rpc.LatestBlockRequestHandlerFn
	RegisterRPCRoutes                  = rpc.RegisterRPCRoutes
	CLIVersionRequestHandler           = rpc.CLIVersionRequestHandler
	NodeVersionRequestHandler          = rpc.NodeVersionRequestHandler
	StatusCommand                      = rpc.StatusCommand
	NodeInfoRequestHandlerFn           = rpc.NodeInfoRequestHandlerFn
	NodeSyncingRequestHandlerFn        = rpc.NodeSyncingRequestHandlerFn
	ValidatorCommand                   = rpc.ValidatorCommand
	GetValidators                      = rpc.GetValidators
	ValidatorSetRequestHandlerFn       = rpc.ValidatorSetRequestHandlerFn
	LatestValidatorSetRequestHandlerFn = rpc.LatestValidatorSetRequestHandlerFn
	BroadcastTxRequest                 = tx.BroadcastTxRequest
	GetBroadcastCommand                = tx.GetBroadcastCommand
	EncodeTxRequestHandlerFn           = tx.EncodeTxRequestHandlerFn
	GetEncodeCommand                   = tx.GetEncodeCommand
	SearchTxCmd                        = tx.SearchTxCmd
	QueryTxCmd                         = tx.QueryTxCmd
	QueryTxsByTagsRequestHandlerFn     = tx.QueryTxsByTagsRequestHandlerFn
	QueryTxRequestHandlerFn            = tx.QueryTxRequestHandlerFn
	RegisterTxRoutes                   = tx.RegisterTxRoutes
	SearchTxs                          = tx.SearchTxs
	ValidateTxResult                   = tx.ValidateTxResult
	GenerateOrBroadcastMsgs            = utils.GenerateOrBroadcastMsgs
	CompleteAndBroadcastTxCLI          = utils.CompleteAndBroadcastTxCLI
	EnrichWithGas                      = utils.EnrichWithGas
	CalculateGas                       = utils.CalculateGas
	PrintUnsignedStdTx                 = utils.PrintUnsignedStdTx
	SignStdTx                          = utils.SignStdTx
	SignStdTxWithSignerAddress         = utils.SignStdTxWithSignerAddress
	ReadStdTxFromFile                  = utils.ReadStdTxFromFile
	GetTxEncoder                       = utils.GetTxEncoder
	PrepareTxBuilder                   = utils.PrepareTxBuilder
	BufferStdin                        = input.BufferStdin
	OverrideStdin                      = input.OverrideStdin
	GetPassword                        = input.GetPassword
	GetCheckPassword                   = input.GetCheckPassword
	GetConfirmation                    = input.GetConfirmation
	GetString                          = input.GetString
	PrintPrefixed                      = input.PrintPrefixed

	// variable aliases
	LineBreak  = flags.LineBreak
	GasFlagVar = flags.GasFlagVar
)

type (
	CLIContext             = context.CLIContext
	GasSetting             = flags.GasSetting
	AddNewKey              = keys.AddNewKey
	RecoverKey             = keys.RecoverKey
	UpdateKeyReq           = keys.UpdateKeyReq
	DeleteKeyReq           = keys.DeleteKeyReq
	RestServer             = lcd.RestServer
	ValidatorOutput        = rpc.ValidatorOutput
	ResultValidatorsOutput = rpc.ResultValidatorsOutput
	BroadcastReq           = tx.BroadcastReq
	EncodeResp             = tx.EncodeResp
	GasEstimateResponse    = utils.GasEstimateResponse
)

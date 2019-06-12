// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/staking/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/staking/types
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/staking/exported
package staking

import (
	"github.com/andrecronje/cosmos-sdk/x/staking/exported"
	"github.com/andrecronje/cosmos-sdk/x/staking/keeper"
	"github.com/andrecronje/cosmos-sdk/x/staking/types"
)

const (
	DefaultParamspace                  = keeper.DefaultParamspace
	DefaultCodespace                   = types.DefaultCodespace
	CodeInvalidValidator               = types.CodeInvalidValidator
	CodeInvalidDelegation              = types.CodeInvalidDelegation
	CodeInvalidInput                   = types.CodeInvalidInput
	CodeValidatorJailed                = types.CodeValidatorJailed
	CodeInvalidAddress                 = types.CodeInvalidAddress
	CodeUnauthorized                   = types.CodeUnauthorized
	CodeInternal                       = types.CodeInternal
	CodeUnknownRequest                 = types.CodeUnknownRequest
	ModuleName                         = types.ModuleName
	StoreKey                           = types.StoreKey
	TStoreKey                          = types.TStoreKey
	QuerierRoute                       = types.QuerierRoute
	RouterKey                          = types.RouterKey
	DefaultUnbondingTime               = types.DefaultUnbondingTime
	DefaultMaxValidators               = types.DefaultMaxValidators
	DefaultMaxEntries                  = types.DefaultMaxEntries
	QueryValidators                    = types.QueryValidators
	QueryValidator                     = types.QueryValidator
	QueryDelegatorDelegations          = types.QueryDelegatorDelegations
	QueryDelegatorUnbondingDelegations = types.QueryDelegatorUnbondingDelegations
	QueryRedelegations                 = types.QueryRedelegations
	QueryValidatorDelegations          = types.QueryValidatorDelegations
	QueryValidatorRedelegations        = types.QueryValidatorRedelegations
	QueryValidatorUnbondingDelegations = types.QueryValidatorUnbondingDelegations
	QueryDelegation                    = types.QueryDelegation
	QueryUnbondingDelegation           = types.QueryUnbondingDelegation
	QueryDelegatorValidators           = types.QueryDelegatorValidators
	QueryDelegatorValidator            = types.QueryDelegatorValidator
	QueryPool                          = types.QueryPool
	QueryParameters                    = types.QueryParameters
	MaxMonikerLength                   = types.MaxMonikerLength
	MaxIdentityLength                  = types.MaxIdentityLength
	MaxWebsiteLength                   = types.MaxWebsiteLength
	MaxDetailsLength                   = types.MaxDetailsLength
	DoNotModifyDesc                    = types.DoNotModifyDesc
)

var (
	// functions aliases
	RegisterInvariants                 = keeper.RegisterInvariants
	AllInvariants                      = keeper.AllInvariants
	SupplyInvariants                   = keeper.SupplyInvariants
	NonNegativePowerInvariant          = keeper.NonNegativePowerInvariant
	PositiveDelegationInvariant        = keeper.PositiveDelegationInvariant
	DelegatorSharesInvariant           = keeper.DelegatorSharesInvariant
	NewKeeper                          = keeper.NewKeeper
	ParamKeyTable                      = keeper.ParamKeyTable
	NewQuerier                         = keeper.NewQuerier
	RegisterCodec                      = types.RegisterCodec
	NewCommissionRates                 = types.NewCommissionRates
	NewCommission                      = types.NewCommission
	NewCommissionWithTime              = types.NewCommissionWithTime
	NewDelegation                      = types.NewDelegation
	MustMarshalDelegation              = types.MustMarshalDelegation
	MustUnmarshalDelegation            = types.MustUnmarshalDelegation
	UnmarshalDelegation                = types.UnmarshalDelegation
	NewUnbondingDelegation             = types.NewUnbondingDelegation
	NewUnbondingDelegationEntry        = types.NewUnbondingDelegationEntry
	MustMarshalUBD                     = types.MustMarshalUBD
	MustUnmarshalUBD                   = types.MustUnmarshalUBD
	UnmarshalUBD                       = types.UnmarshalUBD
	NewRedelegation                    = types.NewRedelegation
	NewRedelegationEntry               = types.NewRedelegationEntry
	MustMarshalRED                     = types.MustMarshalRED
	MustUnmarshalRED                   = types.MustUnmarshalRED
	UnmarshalRED                       = types.UnmarshalRED
	NewDelegationResp                  = types.NewDelegationResp
	NewRedelegationResponse            = types.NewRedelegationResponse
	NewRedelegationEntryResponse       = types.NewRedelegationEntryResponse
	ErrNilValidatorAddr                = types.ErrNilValidatorAddr
	ErrBadValidatorAddr                = types.ErrBadValidatorAddr
	ErrNoValidatorFound                = types.ErrNoValidatorFound
	ErrValidatorOwnerExists            = types.ErrValidatorOwnerExists
	ErrValidatorPubKeyExists           = types.ErrValidatorPubKeyExists
	ErrValidatorPubKeyTypeNotSupported = types.ErrValidatorPubKeyTypeNotSupported
	ErrValidatorJailed                 = types.ErrValidatorJailed
	ErrBadRemoveValidator              = types.ErrBadRemoveValidator
	ErrDescriptionLength               = types.ErrDescriptionLength
	ErrCommissionNegative              = types.ErrCommissionNegative
	ErrCommissionHuge                  = types.ErrCommissionHuge
	ErrCommissionGTMaxRate             = types.ErrCommissionGTMaxRate
	ErrCommissionUpdateTime            = types.ErrCommissionUpdateTime
	ErrCommissionChangeRateNegative    = types.ErrCommissionChangeRateNegative
	ErrCommissionChangeRateGTMaxRate   = types.ErrCommissionChangeRateGTMaxRate
	ErrCommissionGTMaxChangeRate       = types.ErrCommissionGTMaxChangeRate
	ErrSelfDelegationBelowMinimum      = types.ErrSelfDelegationBelowMinimum
	ErrMinSelfDelegationInvalid        = types.ErrMinSelfDelegationInvalid
	ErrMinSelfDelegationDecreased      = types.ErrMinSelfDelegationDecreased
	ErrNilDelegatorAddr                = types.ErrNilDelegatorAddr
	ErrBadDenom                        = types.ErrBadDenom
	ErrBadDelegationAddr               = types.ErrBadDelegationAddr
	ErrBadDelegationAmount             = types.ErrBadDelegationAmount
	ErrNoDelegation                    = types.ErrNoDelegation
	ErrBadDelegatorAddr                = types.ErrBadDelegatorAddr
	ErrNoDelegatorForAddress           = types.ErrNoDelegatorForAddress
	ErrInsufficientShares              = types.ErrInsufficientShares
	ErrDelegationValidatorEmpty        = types.ErrDelegationValidatorEmpty
	ErrNotEnoughDelegationShares       = types.ErrNotEnoughDelegationShares
	ErrBadSharesAmount                 = types.ErrBadSharesAmount
	ErrBadSharesPercent                = types.ErrBadSharesPercent
	ErrNotMature                       = types.ErrNotMature
	ErrNoUnbondingDelegation           = types.ErrNoUnbondingDelegation
	ErrMaxUnbondingDelegationEntries   = types.ErrMaxUnbondingDelegationEntries
	ErrBadRedelegationAddr             = types.ErrBadRedelegationAddr
	ErrNoRedelegation                  = types.ErrNoRedelegation
	ErrSelfRedelegation                = types.ErrSelfRedelegation
	ErrVerySmallRedelegation           = types.ErrVerySmallRedelegation
	ErrBadRedelegationDst              = types.ErrBadRedelegationDst
	ErrTransitiveRedelegation          = types.ErrTransitiveRedelegation
	ErrMaxRedelegationEntries          = types.ErrMaxRedelegationEntries
	ErrDelegatorShareExRateInvalid     = types.ErrDelegatorShareExRateInvalid
	ErrBothShareMsgsGiven              = types.ErrBothShareMsgsGiven
	ErrNeitherShareMsgsGiven           = types.ErrNeitherShareMsgsGiven
	ErrMissingSignature                = types.ErrMissingSignature
	NewGenesisState                    = types.NewGenesisState
	DefaultGenesisState                = types.DefaultGenesisState
	NewMultiStakingHooks               = types.NewMultiStakingHooks
	GetValidatorKey                    = types.GetValidatorKey
	GetValidatorByConsAddrKey          = types.GetValidatorByConsAddrKey
	AddressFromLastValidatorPowerKey   = types.AddressFromLastValidatorPowerKey
	GetValidatorsByPowerIndexKey       = types.GetValidatorsByPowerIndexKey
	GetLastValidatorPowerKey           = types.GetLastValidatorPowerKey
	ParseValidatorPowerRankKey         = types.ParseValidatorPowerRankKey
	GetValidatorQueueTimeKey           = types.GetValidatorQueueTimeKey
	GetDelegationKey                   = types.GetDelegationKey
	GetDelegationsKey                  = types.GetDelegationsKey
	GetUBDKey                          = types.GetUBDKey
	GetUBDByValIndexKey                = types.GetUBDByValIndexKey
	GetUBDKeyFromValIndexKey           = types.GetUBDKeyFromValIndexKey
	GetUBDsKey                         = types.GetUBDsKey
	GetUBDsByValIndexKey               = types.GetUBDsByValIndexKey
	GetUnbondingDelegationTimeKey      = types.GetUnbondingDelegationTimeKey
	GetREDKey                          = types.GetREDKey
	GetREDByValSrcIndexKey             = types.GetREDByValSrcIndexKey
	GetREDByValDstIndexKey             = types.GetREDByValDstIndexKey
	GetREDKeyFromValSrcIndexKey        = types.GetREDKeyFromValSrcIndexKey
	GetREDKeyFromValDstIndexKey        = types.GetREDKeyFromValDstIndexKey
	GetRedelegationTimeKey             = types.GetRedelegationTimeKey
	GetREDsKey                         = types.GetREDsKey
	GetREDsFromValSrcIndexKey          = types.GetREDsFromValSrcIndexKey
	GetREDsToValDstIndexKey            = types.GetREDsToValDstIndexKey
	GetREDsByDelToValDstIndexKey       = types.GetREDsByDelToValDstIndexKey
	NewMsgCreateValidator              = types.NewMsgCreateValidator
	NewMsgEditValidator                = types.NewMsgEditValidator
	NewMsgDelegate                     = types.NewMsgDelegate
	NewMsgBeginRedelegate              = types.NewMsgBeginRedelegate
	NewMsgUndelegate                   = types.NewMsgUndelegate
	NewParams                          = types.NewParams
	DefaultParams                      = types.DefaultParams
	MustUnmarshalParams                = types.MustUnmarshalParams
	UnmarshalParams                    = types.UnmarshalParams
	InitialPool                        = types.InitialPool
	MustUnmarshalPool                  = types.MustUnmarshalPool
	UnmarshalPool                      = types.UnmarshalPool
	NewQueryDelegatorParams            = types.NewQueryDelegatorParams
	NewQueryValidatorParams            = types.NewQueryValidatorParams
	NewQueryBondsParams                = types.NewQueryBondsParams
	NewQueryRedelegationParams         = types.NewQueryRedelegationParams
	NewQueryValidatorsParams           = types.NewQueryValidatorsParams
	NewValidator                       = types.NewValidator
	MustMarshalValidator               = types.MustMarshalValidator
	MustUnmarshalValidator             = types.MustUnmarshalValidator
	UnmarshalValidator                 = types.UnmarshalValidator
	NewDescription                     = types.NewDescription

	// variable aliases
	ModuleCdc                        = types.ModuleCdc
	PoolKey                          = types.PoolKey
	LastValidatorPowerKey            = types.LastValidatorPowerKey
	LastTotalPowerKey                = types.LastTotalPowerKey
	ValidatorsKey                    = types.ValidatorsKey
	ValidatorsByConsAddrKey          = types.ValidatorsByConsAddrKey
	ValidatorsByPowerIndexKey        = types.ValidatorsByPowerIndexKey
	DelegationKey                    = types.DelegationKey
	UnbondingDelegationKey           = types.UnbondingDelegationKey
	UnbondingDelegationByValIndexKey = types.UnbondingDelegationByValIndexKey
	RedelegationKey                  = types.RedelegationKey
	RedelegationByValSrcIndexKey     = types.RedelegationByValSrcIndexKey
	RedelegationByValDstIndexKey     = types.RedelegationByValDstIndexKey
	UnbondingQueueKey                = types.UnbondingQueueKey
	RedelegationQueueKey             = types.RedelegationQueueKey
	ValidatorQueueKey                = types.ValidatorQueueKey
	KeyUnbondingTime                 = types.KeyUnbondingTime
	KeyMaxValidators                 = types.KeyMaxValidators
	KeyMaxEntries                    = types.KeyMaxEntries
	KeyBondDenom                     = types.KeyBondDenom
)

type (
	Keeper                    = keeper.Keeper
	Commission                = types.Commission
	CommissionRates           = types.CommissionRates
	DVPair                    = types.DVPair
	DVVTriplet                = types.DVVTriplet
	Delegation                = types.Delegation
	Delegations               = types.Delegations
	UnbondingDelegation       = types.UnbondingDelegation
	UnbondingDelegationEntry  = types.UnbondingDelegationEntry
	UnbondingDelegations      = types.UnbondingDelegations
	Redelegation              = types.Redelegation
	RedelegationEntry         = types.RedelegationEntry
	Redelegations             = types.Redelegations
	DelegationResponse        = types.DelegationResponse
	DelegationResponses       = types.DelegationResponses
	RedelegationResponse      = types.RedelegationResponse
	RedelegationEntryResponse = types.RedelegationEntryResponse
	RedelegationResponses     = types.RedelegationResponses
	CodeType                  = types.CodeType
	AccountKeeper             = types.AccountKeeper
	GenesisState              = types.GenesisState
	LastValidatorPower        = types.LastValidatorPower
	MultiStakingHooks         = types.MultiStakingHooks
	MsgCreateValidator        = types.MsgCreateValidator
	MsgEditValidator          = types.MsgEditValidator
	MsgDelegate               = types.MsgDelegate
	MsgBeginRedelegate        = types.MsgBeginRedelegate
	MsgUndelegate             = types.MsgUndelegate
	Params                    = types.Params
	Pool                      = types.Pool
	QueryDelegatorParams      = types.QueryDelegatorParams
	QueryValidatorParams      = types.QueryValidatorParams
	QueryBondsParams          = types.QueryBondsParams
	QueryRedelegationParams   = types.QueryRedelegationParams
	QueryValidatorsParams     = types.QueryValidatorsParams
	Validator                 = types.Validator
	Validators                = types.Validators
	Description               = types.Description
	DelegationI               = exported.DelegationI
	ValidatorI                = exported.ValidatorI
)

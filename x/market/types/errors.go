package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	errCodeInvalidOrder uint32 = iota + 1
	errCodeEmptyProvider
	errCodeSameAccount
	errCodeInternal
	errCodeOverOrder
	errCodeAttributeMismatch
	errCodeUnknownBid
	errCodeUnknownLease
	errCodeUnknownLeaseForOrder
	errCodeUnknownOrderForBid
	errCodeLeaseNotActive
	errCodeBidNotActive
	errCodeUnknownOrder
	errCodeNoLeaseForOrder
	errCodeOrderNotFound
	errCodeBidNotFound
	errCodeBidZeroPrice
	errCodeLeaseNotFound
	errCodeBidExists
	errCodeInvalidPrice
	errCodeOrderActive
	errCodeOrderClosed
	errCodeOrderExists
	errCodeOrderDurationExceeded
	errCodeOrderTooEarly
	errInvalidDeposit
	errInvalidParam
)

var (
	// ErrInvalidOrder is the error when order id is invalid
	ErrInvalidOrder = sdkerrors.Register(ModuleName, errCodeInvalidOrder, "invalid: order id")
	// ErrEmptyProvider is the error when provider is empty
	ErrEmptyProvider = sdkerrors.Register(ModuleName, errCodeEmptyProvider, "empty provider")
	// ErrSameAccount is the error when owner and provider are the same account
	ErrSameAccount = sdkerrors.Register(ModuleName, errCodeSameAccount, "owner and provider are the same account")
	// ErrInternal is the error for internal error
	ErrInternal = sdkerrors.Register(ModuleName, errCodeInternal, "internal error")
	// ErrBidOverOrder is the error when bid price is above max order price
	ErrBidOverOrder = sdkerrors.Register(ModuleName, errCodeOverOrder, "bid price above max order price")
	// ErrAttributeMismatch is the error for attribute mismatch
	ErrAttributeMismatch = sdkerrors.Register(ModuleName, errCodeAttributeMismatch, "attribute mismatch")
	// ErrUnknownBid is the error for unknown bid
	ErrUnknownBid = sdkerrors.Register(ModuleName, errCodeUnknownBid, "unknown bid")
	// ErrUnknownLease is the error for unknown bid
	ErrUnknownLease = sdkerrors.Register(ModuleName, errCodeUnknownLease, "unknown lease")
	// ErrUnknownLeaseForBid is the error when lease is unknown for bid
	ErrUnknownLeaseForBid = sdkerrors.Register(ModuleName, errCodeUnknownLeaseForOrder, "unknown lease for bid")
	// ErrUnknownOrderForBid is the error when order is unknown for bid
	ErrUnknownOrderForBid = sdkerrors.Register(ModuleName, errCodeUnknownOrderForBid, "unknown order for bid")
	// ErrLeaseNotActive is the error when lease is not active
	ErrLeaseNotActive = sdkerrors.Register(ModuleName, errCodeLeaseNotActive, "lease not active")
	// ErrBidNotActive is the error when bid is not matched
	ErrBidNotActive = sdkerrors.Register(ModuleName, errCodeBidNotActive, "bid not active")
	// ErrUnknownOrder is the error when order is unknown
	ErrUnknownOrder = sdkerrors.Register(ModuleName, errCodeUnknownOrder, "unknown order")
	// ErrNoLeaseForOrder is the error when there is no lease for order
	ErrNoLeaseForOrder = sdkerrors.Register(ModuleName, errCodeNoLeaseForOrder, "no lease for order")
	// ErrOrderNotFound order not found
	ErrOrderNotFound = sdkerrors.Register(ModuleName, errCodeOrderNotFound, "invalid order: order not found")
	// ErrBidNotFound bid not found
	ErrBidNotFound = sdkerrors.Register(ModuleName, errCodeBidNotFound, "invalid bid: bid not found")
	// ErrBidZeroPrice zero price
	ErrBidZeroPrice = sdkerrors.Register(ModuleName, errCodeBidZeroPrice, "invalid bid: zero price")
	// ErrLeaseNotFound lease not found
	ErrLeaseNotFound = sdkerrors.Register(ModuleName, errCodeLeaseNotFound, "invalid lease: lease not found")
	// ErrBidExists bid exists
	ErrBidExists = sdkerrors.Register(ModuleName, errCodeBidExists, "invalid bid: bid exists from provider")
	// ErrBidInvalidPrice bid invalid price
	ErrBidInvalidPrice = sdkerrors.Register(ModuleName, errCodeInvalidPrice, "bid price is invalid")
	// ErrOrderMatched order matched
	ErrOrderActive = sdkerrors.New(ModuleName, errCodeOrderActive, "order active")
	// ErrOrderClosed order closed
	ErrOrderClosed = sdkerrors.New(ModuleName, errCodeOrderClosed, "order closed")
	// ErrOrderExists indicates a new order was proposed overwrite the existing store key
	ErrOrderExists = sdkerrors.New(ModuleName, errCodeOrderExists, "order already exists in store")
	// ErrOrderTooEarly to match bid
	ErrOrderTooEarly = sdkerrors.New(ModuleName, errCodeOrderTooEarly, "order: chain height to low for bidding")
	// ErrOrderDurationExceeded order should be closed
	ErrOrderDurationExceeded = sdkerrors.New(ModuleName, errCodeOrderDurationExceeded, "order duration has exceeded the bidding duration")
	// ErrInvalidDeposit indicates an invalid deposit
	ErrInvalidDeposit = sdkerrors.Register(ModuleName, errInvalidDeposit, "Deposit invalid")
	// ErrInvalidParam indicates an invalid chain parameter
	ErrInvalidParam = sdkerrors.Register(ModuleName, errInvalidParam, "parameter invalid")
)

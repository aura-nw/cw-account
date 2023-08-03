package utils

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	AuraExponent = 6
	BaseCoinUnit = "uaura"
)

var (
	// DevnetChainID defines the Aura chain ID for devnet
	DevnetChainID = "testnet"
)

// IsDevnet returns true if the chain-id has the Aura devnet chain prefix.
func IsDevnet(chainID string) bool {
	return strings.HasPrefix(chainID, DevnetChainID)
}

// RegisterDenoms registers token denoms.
func RegisterDenoms() {
	err := sdk.RegisterDenom(BaseCoinUnit, sdk.NewDecWithPrec(1, AuraExponent))
	if err != nil {
		panic(err)
	}
}

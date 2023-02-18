package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	yaml "gopkg.in/yaml.v2"
)

// OrderBook params default values
const (
	// Default maximum book participants.
	DefaultMaxBookParticipants uint64 = 100

	// Default batch settlement count.
	DefaultBatchSettlementCount uint64 = 100
)

var (
	KeyMaxBookParticipants = []byte("MaxBookParticipants")

	KeyBatchSettlementCount = []byte("BatchSettlementCount")
)

// ParamTable for orderbook module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(maxBookParticipants, batchSettlementCount uint64) Params {
	return Params{
		MaxBookParticipants:  maxBookParticipants,
		BatchSettlementCount: batchSettlementCount,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxBookParticipants, &p.MaxBookParticipants, validateMaxBookParticipants),
		paramtypes.NewParamSetPair(KeyBatchSettlementCount, &p.BatchSettlementCount, validateBatchSettlementCount),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultMaxBookParticipants,
		DefaultBatchSettlementCount,
	)
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validate a set of params
func (p Params) Validate() error {
	if err := validateMaxBookParticipants(p.MaxBookParticipants); err != nil {
		return err
	}

	if err := validateBatchSettlementCount(p.BatchSettlementCount); err != nil {
		return err
	}

	return nil
}

func validateMaxBookParticipants(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("maximum book participants must be positive: %d", v)
	}

	return nil
}

func validateBatchSettlementCount(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("batch settlement count must be positive: %d", v)
	}

	return nil
}

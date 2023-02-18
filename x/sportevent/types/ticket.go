package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sge-network/sge/utils"
)

// Validate validates sport-event add ticket payload.
func (payload *SportEventAddTicketPayload) Validate(ctx sdk.Context, p *Params) error {
	if err := validateEventTS(ctx, payload.StartTS, payload.EndTS); err != nil {
		return err
	}

	if !utils.IsValidUID(payload.UID) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid uid for the sport-event")
	}

	if len(payload.Odds) < 2 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "not provided enough odds for the event")
	}

	if strings.TrimSpace(payload.Meta) == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "meta is mandatory for the sport-event")
	}

	if len(payload.Meta) > MaxAllowedCharactersForMeta {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "meta length should be less than %d characters", MaxAllowedCharactersForMeta)
	}

	if payload.SrContributionForHouse.IsNil() || payload.SrContributionForHouse.LT(sdk.OneInt()) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "sr contribution cannot be nil or less than 1")
	}

	oddsSet := make(map[string]Odds, len(payload.Odds))
	for _, o := range payload.Odds {
		if o.Meta == "" {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "meta is mandatory for odds with uuid %s", o.UID)
		}
		if len(o.Meta) > MaxAllowedCharactersForMeta {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "meta length should be less than %d characters", MaxAllowedCharactersForMeta)
		}
		if !utils.IsValidUID(o.UID) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "odds-uid passed is invalid")
		}
		if _, exist := oddsSet[o.UID]; exist {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "duplicate odds-uid in request")
		}
		oddsSet[o.UID] = Odds{}
	}

	betConstraints := payload.GetBetConstraints()
	if betConstraints != nil {
		if err := betConstraints.validate(p); err != nil {
			return err
		}
	}

	if payload.SrContributionForHouse.GT(p.EventMaxSrContribution) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "sr contribution cannot be more than %d", p.EventMaxSrContribution.Int64())
	}

	return nil
}

// Validate validates sport-event update ticket payload.
func (payload *SportEventUpdateTicketPayload) Validate(ctx sdk.Context, p *Params) error {
	if err := validateEventTS(ctx, payload.StartTS, payload.EndTS); err != nil {
		return err
	}

	betConstraints := payload.GetBetConstraints()
	if betConstraints != nil {
		if err := betConstraints.validate(p); err != nil {
			return err
		}
	}

	return nil
}

// Validate validates sport-event resolution ticket payload.
func (payload *SportEventResolutionTicketPayload) Validate() error {
	if payload.Status == SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED ||
		payload.Status == SportEventStatus_SPORT_EVENT_STATUS_INVALID {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "resolution status passed for the sports event is invalid")
	}

	if payload.ResolutionTS == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid resolution timestamp for the sport-event")
	}

	if !utils.IsValidUID(payload.UID) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid uid for the sport-event")
	}

	if payload.Status == SportEventStatus_SPORT_EVENT_STATUS_RESULT_DECLARED && len(payload.WinnerOddsUIDs) < 1 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "not provided enough winner odds for the sports event")
	}

	for _, wid := range payload.WinnerOddsUIDs {
		if !utils.IsValidUID(wid) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "odds-uid passed is invalid")
		}
	}

	if payload.Status > SportEventStatus_SPORT_EVENT_STATUS_RESULT_DECLARED {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid event resolution status ")
	}

	return nil
}

// validateEventTS validates start and end timestamp of a sport-event.
func validateEventTS(ctx sdk.Context, startTS, endTS uint64) error {
	if endTS <= uint64(ctx.BlockTime().Unix()) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid end timestamp for the sport-event")
	}

	if startTS >= endTS || startTS == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid start timestamp for the sport-event, cannot be (greater than eql to EndTs) or 0")
	}

	return nil
}

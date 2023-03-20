package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewSportEvent(
	uid, creator string,
	startTS, endTS uint64,
	odds []*Odds,
	betConstraints *EventBetConstraints,
	meta string,
	bookUID string,
	srContributionForHouse sdk.Int,
	status SportEventStatus,
) SportEvent {
	return SportEvent{
		UID:                    uid,
		Creator:                creator,
		StartTS:                startTS,
		EndTS:                  endTS,
		Odds:                   odds,
		BetConstraints:         betConstraints,
		Meta:                   meta,
		BookUID:                bookUID,
		SrContributionForHouse: srContributionForHouse,
		Status:                 status,
	}
}

package db

import (
	"context"

	"github.com/scalarorg/xchains-api/internal/db/model"
	"github.com/scalarorg/xchains-api/internal/types"
)

type DBClient interface {
	Ping(ctx context.Context) error
	SaveActiveStakingDelegation(
		ctx context.Context, stakingTxHashHex, stakerPkHex, fpPkHex string,
		stakingTxHex string, amount, startHeight, timelock, outputIndex uint64,
		startTimestamp int64, isOverflow bool, stakerTaprootAddress string,
	) error
	FindDelegationsByStakerPk(
		ctx context.Context, stakerPk string, paginationToken string,
	) (*DbResultMap[model.DelegationDocument], error)
	SaveUnbondingTx(
		ctx context.Context, stakingTxHashHex, unbondingTxHashHex, txHex, signatureHex string,
	) error
	FindDelegationByTxHashHex(ctx context.Context, txHashHex string) (*model.DelegationDocument, error)
	SaveTimeLockExpireCheck(ctx context.Context, stakingTxHashHex string, expireHeight uint64, txType string) error
	SaveUnprocessableMessage(ctx context.Context, messageBody, receipt string) error
	FindUnprocessableMessages(ctx context.Context) ([]model.UnprocessableMessageDocument, error)
	DeleteUnprocessableMessage(ctx context.Context, Receipt interface{}) error
	TransitionToUnbondedState(
		ctx context.Context, stakingTxHashHex string, eligiblePreviousState []types.DelegationState,
	) error
	TransitionToUnbondingState(
		ctx context.Context, txHashHex string, startHeight, timelock, outputIndex uint64, txHex string, startTimestamp int64,
	) error
	TransitionToWithdrawnState(ctx context.Context, txHashHex string) error
	GetOrCreateStatsLock(
		ctx context.Context, stakingTxHashHex string, state string,
	) (*model.StatsLockDocument, error)
	SubtractOverallStats(
		ctx context.Context, stakingTxHashHex, stakerPkHex string, amount uint64,
	) error
	IncrementOverallStats(
		ctx context.Context, stakingTxHashHex, stakerPkHex string, amount uint64,
	) error
	GetOverallStats(ctx context.Context) (*model.OverallStatsDocument, error)
	IncrementFinalityProviderStats(
		ctx context.Context, stakingTxHashHex, fpPkHex string, amount uint64,
	) error
	SubtractFinalityProviderStats(
		ctx context.Context, stakingTxHashHex, fpPkHex string, amount uint64,
	) error
	FindFinalityProviderStats(ctx context.Context, paginationToken string) (*DbResultMap[*model.FinalityProviderStatsDocument], error)
	FindFinalityProviderStatsByFinalityProviderPkHex(
		ctx context.Context, finalityProviderPkHex []string,
	) ([]*model.FinalityProviderStatsDocument, error)
	IncrementStakerStats(
		ctx context.Context, stakingTxHashHex, stakerPkHex string, amount uint64,
	) error
	SubtractStakerStats(
		ctx context.Context, stakingTxHashHex, stakerPkHex string, amount uint64,
	) error
	FindTopStakersByTvl(ctx context.Context, paginationToken string) (*DbResultMap[*model.StakerStatsDocument], error)
	UpsertLatestBtcInfo(
		ctx context.Context, height uint64, confirmedTvl uint64, unconfirmedTvl uint64,
	) error
	GetLatestBtcInfo(ctx context.Context) (*model.BtcInfo, error)
	CheckDelegationExistByStakerTaprootAddress(
		ctx context.Context, address string, extraFilter *DelegationFilter,
	) (bool, error)

	SaveDApp(ctx context.Context, dApp *model.DAppDocument) error
	GetDApp(ctx context.Context) ([]*model.DAppDocument, error)
	UpdateDApp(ctx context.Context, dApp *model.DAppDocument) error
	// TODO: refactor the ID -> Should use ObjectID
	ToggleDApp(ctx context.Context, ID string) error
	DeleteDApp(ctx context.Context, ID string) error

	GetGMPs(ctx context.Context) ([]*model.GMPDocument, error)
}

type DelegationFilter struct {
	AfterTimestamp int64
	States         []types.DelegationState
}

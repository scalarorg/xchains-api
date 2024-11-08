package api

import (
	"github.com/go-chi/chi"
	_ "github.com/scalarorg/xchains-api/docs"
	"github.com/scalarorg/xchains-api/internal/api/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *Server) SetupRoutes(r *chi.Mux) {
	handlers := a.handlers
	r.Get("/healthcheck", registerHandler(handlers.HealthCheck))

	r.Get("/v1/staker/delegations", registerHandler(handlers.GetStakerDelegations))
	r.Get("/v1/staker/delegation/check", registerHandler(handlers.CheckStakerDelegationExist))
	r.Post("/v1/staker/staking-psbt/create", registerHandler(handlers.CreateStakingPsbt))
	r.Post("/v1/staker/unstaking-psbt/create", registerHandler(handlers.CreateUnstakingPsbt))
	r.Post("/v1/unbonding", registerHandler(handlers.UnbondDelegation))
	r.Get("/v1/unbonding/eligibility", registerHandler(handlers.GetUnbondingEligibility))
	r.Get("/v1/global-params", registerHandler(handlers.GetBabylonGlobalParams))
	r.Get("/v1/finality-providers", registerHandler(handlers.GetFinalityProviders))
	r.Get("/v1/stats", registerHandler(handlers.GetOverallStats))
	r.Get("/v1/stats/staker", registerHandler(handlers.GetTopStakerStats))

	r.Get("/v1/delegation", registerHandler(handlers.GetDelegationByTxHash))

	registerParamsHandler(r, handlers)
	registerDAppHandler(r, handlers)
	registerGmpHandler(r, handlers)
	registerVaultHandler(r, handlers)

	r.Get("/swagger/*", httpSwagger.WrapHandler)
}
func registerParamsHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Get("/v1/params/covenant", registerHandler(handlers.GetCovenantParams))
}
func registerDAppHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Get("/v1/dApp", registerHandler(handlers.GetDApp))
	r.Post("/v1/dApp", registerHandler(handlers.CreateDApp))
	r.Put("/v1/dApp", registerHandler(handlers.UpdateDApp))
	r.Patch("/v1/dApp", registerHandler(handlers.ToggleDApp))
	r.Delete("/v1/dApp", registerHandler(handlers.DeleteDApp))
}
func registerGmpHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Post("/v1/gmp/GMPStats", registerHandler(handlers.GMPStats))
	r.Post("/v1/gmp/GMPStatsAVGTimes", registerHandler(handlers.GMPStatsAVGTimes))
	r.Post("/v1/gmp/GMPChart", registerHandler(handlers.GMPChart))
	r.Post("/v1/gmp/GMPCumulativeVolume", registerHandler(handlers.GMPCumulativeVolume))
	r.Post("/v1/gmp/GMPTotalVolume", registerHandler(handlers.GMPTotalVolume))
	r.Post("/v1/gmp/GMPTotalFee", registerHandler(handlers.GMPTotalFee))
	r.Post("/v1/gmp/GMPTotalActiveUsers", registerHandler(handlers.GMPTotalActiveUsers))
	r.Post("/v1/gmp/GMPTopUsers", registerHandler(handlers.GMPTopUsers))
	r.Post("/v1/gmp/GMPTopITSAssets", registerHandler(handlers.GMPTopITSAssets))
	r.Post("/v1/gmp/searchGMP", registerHandler(handlers.GMPSearch))
	r.Post("/v1/gmp/getContracts", registerHandler(handlers.GmpGetContracts))
	r.Post("/v1/gmp/getConfigurations", registerHandler(handlers.GmpGetConfigurations))
	r.Post("/v1/gmp/getDataMapping", registerHandler(handlers.GetGMPDataMapping))
	r.Post("/v1/gmp/estimateTimeSpent", registerHandler(handlers.EstimateTimeSpent))

}

func registerVaultHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Post("/v1/vault/searchVault", registerHandler(handlers.SearchVault))
}

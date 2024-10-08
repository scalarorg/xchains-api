package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/gmp"
	"github.com/scalarorg/xchains-api/internal/db/model"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) GMPSearch(ctx context.Context, payload *types.GmpPayload) ([]*gmp.GMPDocument, *types.Error) {
	gmps, err := s.GmpClient.GMPSearch(ctx, payload)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	if gmps == nil {
		gmps = []*gmp.GMPDocument{}
	}
	return gmps, nil
}

func (s *Services) GmpGetContracts(ctx context.Context) ([]*model.GMPDocument, *types.Error) {
	gmps, err := s.DbClient.GetGMPs(ctx)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return gmps, nil
}

func (s *Services) GmpGetConfigurations(ctx context.Context) ([]*model.GMPDocument, *types.Error) {
	gmps, err := s.DbClient.GetGMPs(ctx)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return gmps, nil
}

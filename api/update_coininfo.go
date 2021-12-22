package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	ccoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCoinInfo(ctx context.Context, in *npool.UpdateCoinInfoRequest) (*npool.UpdateCoinInfoResponse, error) {
	if in.GetID() == "" {
		logger.Sugar().Errorf("UpdateCoinInfo check ID is empty")
		return nil, status.Error(codes.InvalidArgument, "ID empty")
	}

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("UpdateCoinInfo check ID not a valid uuid")
		return nil, status.Error(codes.InvalidArgument, "ID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	_, err = coininfo.GetCoinInfoByID(ctx, id)
	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("UpdateCoinInfo call GetCoinInfoByID ID: %v not found", in.GetID())
		return nil, status.Errorf(codes.NotFound, "ID: %v not found", in.GetID())
	}

	if err != nil {
		logger.Sugar().Errorf("UpdateCoinInfo call GetCoinInfoByID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	coinInfo, err := coininfo.UpdateCoinInfoByID(ctx, in.GetPreSale(), in.GetLogo(), in.GetID(), in.GetReservedAmount())
	if err != nil {
		logger.Sugar().Errorf("UpdateCoinInfo call UpdateCoinInfoByID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.UpdateCoinInfoResponse{
		Info: &npool.CoinInfo{
			ID:             coinInfo.ID.String(),
			PreSale:        coinInfo.PreSale,
			Name:           coinInfo.Name,
			Unit:           coinInfo.Unit,
			Logo:           coinInfo.Logo,
			ReservedAmount: price.DBPriceToVisualPrice(coinInfo.ReservedAmount),
			CreatedAt:      coinInfo.CreatedAt,
			UpdatedAt:      coinInfo.UpdatedAt,
		},
	}, nil
}

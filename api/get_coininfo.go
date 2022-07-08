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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCoinInfo get coininfo by id or name, id high prio
func (s *Server) GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (*npool.GetCoinInfoResponse, error) {
	_, span := otel.Tracer(ccoin.ServiceName).Start(ctx, "GetCoinInfo")
	defer span.End()

	span.SetAttributes(
		attribute.String("ID", in.GetID()),
		attribute.String("Name", in.GetName()),
	)

	if in.GetID() == "" && in.GetName() == "" {
		logger.Sugar().Errorf("GetCoinInfo check ID or Name is empty")
		return nil, status.Error(codes.InvalidArgument, "not allow ID or Name both empty")
	}

	var (
		err      error
		coinInfo *ent.CoinInfo
	)

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	if in.GetID() != "" {
		id, err := uuid.Parse(in.GetID())
		if err != nil {
			logger.Sugar().Errorf("GetCoinInfo check ID: %v not a valid uuid", in.GetID())
			return nil, status.Errorf(codes.InvalidArgument, "ID: %v invalid", in.GetID())
		}

		span.AddEvent("call db GetCoinInfoByID",
			trace.WithAttributes(
				attribute.String("ID", in.GetID()),
			),
		)
		coinInfo, err = coininfo.GetCoinInfoByID(ctx, id)
		span.AddEvent("call db GetCoinInfoByID done")

		if ent.IsNotFound(err) {
			logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByID ID: %v not found", in.GetID())
			return nil, status.Errorf(codes.NotFound, "ID: %v not found", in.GetID())
		}

		if err != nil {
			logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByID error %v", err)
			return nil, status.Error(codes.Internal, "internal server error")
		}
	} else if in.GetName() != "" {
		span.AddEvent("call db GetCoinInfoByName",
			trace.WithAttributes(
				attribute.String("Name", in.GetName()),
			),
		)
		coinInfo, err = coininfo.GetCoinInfoByName(ctx, in.GetName())
		span.AddEvent("call db GetCoinInfoByName done")

		if ent.IsNotFound(err) {
			logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByName Name: %v not found", in.GetName())
			return nil, status.Errorf(codes.NotFound, "Name: %v not found", in.GetName())
		}
		if err != nil {
			logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByName error %v", err)
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &npool.GetCoinInfoResponse{
		Info: &npool.CoinInfo{
			ID:             coinInfo.ID.String(),
			PreSale:        coinInfo.PreSale,
			ForPay:         coinInfo.ForPay,
			Name:           coinInfo.Name,
			ENV:            coinInfo.Env,
			Unit:           coinInfo.Unit,
			ReservedAmount: price.DBPriceToVisualPrice(coinInfo.ReservedAmount),
			Logo:           coinInfo.Logo,
			HomePage:       coinInfo.HomePage,
			Specs:          coinInfo.Specs,
			CreatedAt:      coinInfo.CreatedAt,
			UpdatedAt:      coinInfo.UpdatedAt,
		},
	}, nil
}

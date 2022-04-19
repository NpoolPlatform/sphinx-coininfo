package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/description"
	ccoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCoinDescription ..
func (s *Server) GetCoinDescription(ctx context.Context, in *npool.GetCoinDescriptionRequest) (*npool.GetCoinDescriptionResponse, error) {
	if in.GetCoinTypeID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check CoinTypeID is empty")
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID is empty")
	}

	coinID, err := uuid.Parse(in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse CoinTypeID: %s invalid", in.GetCoinTypeID())
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	coinDescs, total, err := description.GetCoinDescriptionByCoinID(ctx, coinID, in.GetLimit(), in.GetOffset())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription call GetCoinDescriptionByCoinID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	descInfos := make([]*npool.CoinDescriptionInfo, total)
	for _, info := range coinDescs {
		descInfos = append(descInfos, &npool.CoinDescriptionInfo{
			ID:         info.ID.String(),
			CoinTypeID: info.CoinTypeID.String(),
			Title:      info.Title,
			Message:    info.Message,
			UsedFor:    info.UsedFor,
			CreatedAt:  info.CreatedAt,
			UpdatedAt:  info.UpdatedAt,
		})
	}

	return &npool.GetCoinDescriptionResponse{
		Infos: descInfos,
		Total: int32(total),
	}, nil
}

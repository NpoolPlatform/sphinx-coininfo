package coininfo

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	constant "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/google/uuid"
)

func GetCoinInfoByID(ctx context.Context, id uuid.UUID) (coinInfo *ent.CoinInfo, err error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}
	coinInfo, err = client.CoinInfo.Query().
		Where(coininfo.ID(id)).
		Only(ctx)
	return
}

func GetCoinInfoByName(ctx context.Context, coinName string) (coinInfo *ent.CoinInfo, err error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}
	coinInfo, err = client.CoinInfo.Query().
		Where(coininfo.Name(coinName)).
		Only(ctx)
	return
}

func ExistCoinInfoByID(ctx context.Context, coinID uuid.UUID) (bool, error) {
	client, err := db.Client()
	if err != nil {
		return false, err
	}
	return client.CoinInfo.Query().
		Where(coininfo.IDEQ(coinID)).
		Exist(ctx)
}

type GetAllCoinInfosParams struct {
	PreSale       bool
	Name          string
	Offset, Limit int
}

func GetAllCoinInfos(ctx context.Context, params GetAllCoinInfosParams) ([]*ent.CoinInfo, int, error) {
	if params.Limit == 0 {
		params.Limit = int(constant.DefaultPageSize)
	}

	client, err := db.Client()
	if err != nil {
		return nil, 0, err
	}

	stm := client.
		CoinInfo.
		Query()

	if params.Name != "" {
		stm.Where(coininfo.NameEQ(params.Name))
	}

	if params.PreSale {
		stm.Where(coininfo.PreSaleEQ(params.PreSale))
	}

	// total
	total, err := stm.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// infos
	coinInfos, err := stm.
		Order(ent.Desc(coininfo.FieldUpdatedAt)).
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)

	return coinInfos, total, err
}

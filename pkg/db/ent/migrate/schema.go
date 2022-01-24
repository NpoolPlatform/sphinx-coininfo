// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoinInfosColumns holds the columns for the "coin_infos" table.
	CoinInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "unit", Type: field.TypeString, Default: ""},
		{Name: "reserved_amount", Type: field.TypeUint64, Default: 0},
		{Name: "pre_sale", Type: field.TypeBool, Default: false},
		{Name: "logo", Type: field.TypeString, Default: ""},
		{Name: "env", Type: field.TypeString, Default: "main"},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
	}
	// CoinInfosTable holds the schema information for the "coin_infos" table.
	CoinInfosTable = &schema.Table{
		Name:       "coin_infos",
		Columns:    CoinInfosColumns,
		PrimaryKey: []*schema.Column{CoinInfosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "coininfo_name",
				Unique:  true,
				Columns: []*schema.Column{CoinInfosColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoinInfosTable,
	}
)

func init() {
}

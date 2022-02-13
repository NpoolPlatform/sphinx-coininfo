package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type CoinInfo struct {
	ent.Schema
}

func (CoinInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").NotEmpty().Unique(),
		field.String("unit").NotEmpty().Default(""),
		field.Uint64("reserved_amount").
			Default(0),
		field.Bool("pre_sale").Default(false),
		field.String("logo").Default(""),
		field.String("env").Default("main"), // main or test
		field.Bool("for_pay"),
		field.Uint32("created_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("updated_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("deleted_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

func (CoinInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
	}
}

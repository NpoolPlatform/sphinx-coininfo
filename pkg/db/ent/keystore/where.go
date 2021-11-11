// Code generated by entc, DO NOT EDIT.

package keystore

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// PrivateKey applies equality check predicate on the "private_key" field. It's identical to PrivateKeyEQ.
func PrivateKey(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrivateKey), v))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.KeyStore {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KeyStore(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.KeyStore {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KeyStore(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// PrivateKeyEQ applies the EQ predicate on the "private_key" field.
func PrivateKeyEQ(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrivateKey), v))
	})
}

// PrivateKeyNEQ applies the NEQ predicate on the "private_key" field.
func PrivateKeyNEQ(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrivateKey), v))
	})
}

// PrivateKeyIn applies the In predicate on the "private_key" field.
func PrivateKeyIn(vs ...string) predicate.KeyStore {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KeyStore(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrivateKey), v...))
	})
}

// PrivateKeyNotIn applies the NotIn predicate on the "private_key" field.
func PrivateKeyNotIn(vs ...string) predicate.KeyStore {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KeyStore(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrivateKey), v...))
	})
}

// PrivateKeyGT applies the GT predicate on the "private_key" field.
func PrivateKeyGT(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrivateKey), v))
	})
}

// PrivateKeyGTE applies the GTE predicate on the "private_key" field.
func PrivateKeyGTE(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrivateKey), v))
	})
}

// PrivateKeyLT applies the LT predicate on the "private_key" field.
func PrivateKeyLT(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrivateKey), v))
	})
}

// PrivateKeyLTE applies the LTE predicate on the "private_key" field.
func PrivateKeyLTE(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrivateKey), v))
	})
}

// PrivateKeyContains applies the Contains predicate on the "private_key" field.
func PrivateKeyContains(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrivateKey), v))
	})
}

// PrivateKeyHasPrefix applies the HasPrefix predicate on the "private_key" field.
func PrivateKeyHasPrefix(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrivateKey), v))
	})
}

// PrivateKeyHasSuffix applies the HasSuffix predicate on the "private_key" field.
func PrivateKeyHasSuffix(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrivateKey), v))
	})
}

// PrivateKeyEqualFold applies the EqualFold predicate on the "private_key" field.
func PrivateKeyEqualFold(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrivateKey), v))
	})
}

// PrivateKeyContainsFold applies the ContainsFold predicate on the "private_key" field.
func PrivateKeyContainsFold(v string) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrivateKey), v))
	})
}

// HasCoin applies the HasEdge predicate on the "coin" edge.
func HasCoin() predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CoinTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CoinTable, CoinColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCoinWith applies the HasEdge predicate on the "coin" edge with a given conditions (other predicates).
func HasCoinWith(preds ...predicate.CoinInfo) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CoinInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CoinTable, CoinColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.KeyStore) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.KeyStore) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.KeyStore) predicate.KeyStore {
	return predicate.KeyStore(func(s *sql.Selector) {
		p(s.Not())
	})
}

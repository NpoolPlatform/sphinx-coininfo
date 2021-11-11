// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/keystore"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/review"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/walletnode"
)

// CoinInfoUpdate is the builder for updating CoinInfo entities.
type CoinInfoUpdate struct {
	config
	hooks    []Hook
	mutation *CoinInfoMutation
}

// Where appends a list predicates to the CoinInfoUpdate builder.
func (ciu *CoinInfoUpdate) Where(ps ...predicate.CoinInfo) *CoinInfoUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetName sets the "name" field.
func (ciu *CoinInfoUpdate) SetName(s string) *CoinInfoUpdate {
	ciu.mutation.SetName(s)
	return ciu
}

// SetUnit sets the "unit" field.
func (ciu *CoinInfoUpdate) SetUnit(s string) *CoinInfoUpdate {
	ciu.mutation.SetUnit(s)
	return ciu
}

// SetNeedSigninfo sets the "need_signinfo" field.
func (ciu *CoinInfoUpdate) SetNeedSigninfo(b bool) *CoinInfoUpdate {
	ciu.mutation.SetNeedSigninfo(b)
	return ciu
}

// SetNillableNeedSigninfo sets the "need_signinfo" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableNeedSigninfo(b *bool) *CoinInfoUpdate {
	if b != nil {
		ciu.SetNeedSigninfo(*b)
	}
	return ciu
}

// AddKeyIDs adds the "keys" edge to the KeyStore entity by IDs.
func (ciu *CoinInfoUpdate) AddKeyIDs(ids ...int32) *CoinInfoUpdate {
	ciu.mutation.AddKeyIDs(ids...)
	return ciu
}

// AddKeys adds the "keys" edges to the KeyStore entity.
func (ciu *CoinInfoUpdate) AddKeys(k ...*KeyStore) *CoinInfoUpdate {
	ids := make([]int32, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return ciu.AddKeyIDs(ids...)
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (ciu *CoinInfoUpdate) AddTransactionIDs(ids ...int32) *CoinInfoUpdate {
	ciu.mutation.AddTransactionIDs(ids...)
	return ciu
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (ciu *CoinInfoUpdate) AddTransactions(t ...*Transaction) *CoinInfoUpdate {
	ids := make([]int32, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ciu.AddTransactionIDs(ids...)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (ciu *CoinInfoUpdate) AddReviewIDs(ids ...int32) *CoinInfoUpdate {
	ciu.mutation.AddReviewIDs(ids...)
	return ciu
}

// AddReviews adds the "reviews" edges to the Review entity.
func (ciu *CoinInfoUpdate) AddReviews(r ...*Review) *CoinInfoUpdate {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ciu.AddReviewIDs(ids...)
}

// AddWalletNodeIDs adds the "wallet_nodes" edge to the WalletNode entity by IDs.
func (ciu *CoinInfoUpdate) AddWalletNodeIDs(ids ...int32) *CoinInfoUpdate {
	ciu.mutation.AddWalletNodeIDs(ids...)
	return ciu
}

// AddWalletNodes adds the "wallet_nodes" edges to the WalletNode entity.
func (ciu *CoinInfoUpdate) AddWalletNodes(w ...*WalletNode) *CoinInfoUpdate {
	ids := make([]int32, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ciu.AddWalletNodeIDs(ids...)
}

// Mutation returns the CoinInfoMutation object of the builder.
func (ciu *CoinInfoUpdate) Mutation() *CoinInfoMutation {
	return ciu.mutation
}

// ClearKeys clears all "keys" edges to the KeyStore entity.
func (ciu *CoinInfoUpdate) ClearKeys() *CoinInfoUpdate {
	ciu.mutation.ClearKeys()
	return ciu
}

// RemoveKeyIDs removes the "keys" edge to KeyStore entities by IDs.
func (ciu *CoinInfoUpdate) RemoveKeyIDs(ids ...int32) *CoinInfoUpdate {
	ciu.mutation.RemoveKeyIDs(ids...)
	return ciu
}

// RemoveKeys removes "keys" edges to KeyStore entities.
func (ciu *CoinInfoUpdate) RemoveKeys(k ...*KeyStore) *CoinInfoUpdate {
	ids := make([]int32, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return ciu.RemoveKeyIDs(ids...)
}

// ClearTransactions clears all "transactions" edges to the Transaction entity.
func (ciu *CoinInfoUpdate) ClearTransactions() *CoinInfoUpdate {
	ciu.mutation.ClearTransactions()
	return ciu
}

// RemoveTransactionIDs removes the "transactions" edge to Transaction entities by IDs.
func (ciu *CoinInfoUpdate) RemoveTransactionIDs(ids ...int32) *CoinInfoUpdate {
	ciu.mutation.RemoveTransactionIDs(ids...)
	return ciu
}

// RemoveTransactions removes "transactions" edges to Transaction entities.
func (ciu *CoinInfoUpdate) RemoveTransactions(t ...*Transaction) *CoinInfoUpdate {
	ids := make([]int32, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ciu.RemoveTransactionIDs(ids...)
}

// ClearReviews clears all "reviews" edges to the Review entity.
func (ciu *CoinInfoUpdate) ClearReviews() *CoinInfoUpdate {
	ciu.mutation.ClearReviews()
	return ciu
}

// RemoveReviewIDs removes the "reviews" edge to Review entities by IDs.
func (ciu *CoinInfoUpdate) RemoveReviewIDs(ids ...int32) *CoinInfoUpdate {
	ciu.mutation.RemoveReviewIDs(ids...)
	return ciu
}

// RemoveReviews removes "reviews" edges to Review entities.
func (ciu *CoinInfoUpdate) RemoveReviews(r ...*Review) *CoinInfoUpdate {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ciu.RemoveReviewIDs(ids...)
}

// ClearWalletNodes clears all "wallet_nodes" edges to the WalletNode entity.
func (ciu *CoinInfoUpdate) ClearWalletNodes() *CoinInfoUpdate {
	ciu.mutation.ClearWalletNodes()
	return ciu
}

// RemoveWalletNodeIDs removes the "wallet_nodes" edge to WalletNode entities by IDs.
func (ciu *CoinInfoUpdate) RemoveWalletNodeIDs(ids ...int32) *CoinInfoUpdate {
	ciu.mutation.RemoveWalletNodeIDs(ids...)
	return ciu
}

// RemoveWalletNodes removes "wallet_nodes" edges to WalletNode entities.
func (ciu *CoinInfoUpdate) RemoveWalletNodes(w ...*WalletNode) *CoinInfoUpdate {
	ids := make([]int32, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ciu.RemoveWalletNodeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *CoinInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ciu.hooks) == 0 {
		if err = ciu.check(); err != nil {
			return 0, err
		}
		affected, err = ciu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ciu.check(); err != nil {
				return 0, err
			}
			ciu.mutation = mutation
			affected, err = ciu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ciu.hooks) - 1; i >= 0; i-- {
			if ciu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ciu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ciu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *CoinInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *CoinInfoUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *CoinInfoUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciu *CoinInfoUpdate) check() error {
	if v, ok := ciu.mutation.Name(); ok {
		if err := coininfo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ciu.mutation.Unit(); ok {
		if err := coininfo.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf("ent: validator failed for field \"unit\": %w", err)}
		}
	}
	return nil
}

func (ciu *CoinInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coininfo.Table,
			Columns: coininfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: coininfo.FieldID,
			},
		},
	}
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldName,
		})
	}
	if value, ok := ciu.mutation.Unit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldUnit,
		})
	}
	if value, ok := ciu.mutation.NeedSigninfo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coininfo.FieldNeedSigninfo,
		})
	}
	if ciu.mutation.KeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.KeysTable,
			Columns: []string{coininfo.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: keystore.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.RemovedKeysIDs(); len(nodes) > 0 && !ciu.mutation.KeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.KeysTable,
			Columns: []string{coininfo.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: keystore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.KeysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.KeysTable,
			Columns: []string{coininfo.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: keystore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciu.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.TransactionsTable,
			Columns: []string{coininfo.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !ciu.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.TransactionsTable,
			Columns: []string{coininfo.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.TransactionsTable,
			Columns: []string{coininfo.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciu.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.ReviewsTable,
			Columns: []string{coininfo.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !ciu.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.ReviewsTable,
			Columns: []string{coininfo.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.ReviewsTable,
			Columns: []string{coininfo.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciu.mutation.WalletNodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.WalletNodesTable,
			Columns: []string{coininfo.WalletNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: walletnode.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.RemovedWalletNodesIDs(); len(nodes) > 0 && !ciu.mutation.WalletNodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.WalletNodesTable,
			Columns: []string{coininfo.WalletNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: walletnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.WalletNodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.WalletNodesTable,
			Columns: []string{coininfo.WalletNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: walletnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coininfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CoinInfoUpdateOne is the builder for updating a single CoinInfo entity.
type CoinInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CoinInfoMutation
}

// SetName sets the "name" field.
func (ciuo *CoinInfoUpdateOne) SetName(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetName(s)
	return ciuo
}

// SetUnit sets the "unit" field.
func (ciuo *CoinInfoUpdateOne) SetUnit(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetUnit(s)
	return ciuo
}

// SetNeedSigninfo sets the "need_signinfo" field.
func (ciuo *CoinInfoUpdateOne) SetNeedSigninfo(b bool) *CoinInfoUpdateOne {
	ciuo.mutation.SetNeedSigninfo(b)
	return ciuo
}

// SetNillableNeedSigninfo sets the "need_signinfo" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableNeedSigninfo(b *bool) *CoinInfoUpdateOne {
	if b != nil {
		ciuo.SetNeedSigninfo(*b)
	}
	return ciuo
}

// AddKeyIDs adds the "keys" edge to the KeyStore entity by IDs.
func (ciuo *CoinInfoUpdateOne) AddKeyIDs(ids ...int32) *CoinInfoUpdateOne {
	ciuo.mutation.AddKeyIDs(ids...)
	return ciuo
}

// AddKeys adds the "keys" edges to the KeyStore entity.
func (ciuo *CoinInfoUpdateOne) AddKeys(k ...*KeyStore) *CoinInfoUpdateOne {
	ids := make([]int32, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return ciuo.AddKeyIDs(ids...)
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (ciuo *CoinInfoUpdateOne) AddTransactionIDs(ids ...int32) *CoinInfoUpdateOne {
	ciuo.mutation.AddTransactionIDs(ids...)
	return ciuo
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (ciuo *CoinInfoUpdateOne) AddTransactions(t ...*Transaction) *CoinInfoUpdateOne {
	ids := make([]int32, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ciuo.AddTransactionIDs(ids...)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (ciuo *CoinInfoUpdateOne) AddReviewIDs(ids ...int32) *CoinInfoUpdateOne {
	ciuo.mutation.AddReviewIDs(ids...)
	return ciuo
}

// AddReviews adds the "reviews" edges to the Review entity.
func (ciuo *CoinInfoUpdateOne) AddReviews(r ...*Review) *CoinInfoUpdateOne {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ciuo.AddReviewIDs(ids...)
}

// AddWalletNodeIDs adds the "wallet_nodes" edge to the WalletNode entity by IDs.
func (ciuo *CoinInfoUpdateOne) AddWalletNodeIDs(ids ...int32) *CoinInfoUpdateOne {
	ciuo.mutation.AddWalletNodeIDs(ids...)
	return ciuo
}

// AddWalletNodes adds the "wallet_nodes" edges to the WalletNode entity.
func (ciuo *CoinInfoUpdateOne) AddWalletNodes(w ...*WalletNode) *CoinInfoUpdateOne {
	ids := make([]int32, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ciuo.AddWalletNodeIDs(ids...)
}

// Mutation returns the CoinInfoMutation object of the builder.
func (ciuo *CoinInfoUpdateOne) Mutation() *CoinInfoMutation {
	return ciuo.mutation
}

// ClearKeys clears all "keys" edges to the KeyStore entity.
func (ciuo *CoinInfoUpdateOne) ClearKeys() *CoinInfoUpdateOne {
	ciuo.mutation.ClearKeys()
	return ciuo
}

// RemoveKeyIDs removes the "keys" edge to KeyStore entities by IDs.
func (ciuo *CoinInfoUpdateOne) RemoveKeyIDs(ids ...int32) *CoinInfoUpdateOne {
	ciuo.mutation.RemoveKeyIDs(ids...)
	return ciuo
}

// RemoveKeys removes "keys" edges to KeyStore entities.
func (ciuo *CoinInfoUpdateOne) RemoveKeys(k ...*KeyStore) *CoinInfoUpdateOne {
	ids := make([]int32, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return ciuo.RemoveKeyIDs(ids...)
}

// ClearTransactions clears all "transactions" edges to the Transaction entity.
func (ciuo *CoinInfoUpdateOne) ClearTransactions() *CoinInfoUpdateOne {
	ciuo.mutation.ClearTransactions()
	return ciuo
}

// RemoveTransactionIDs removes the "transactions" edge to Transaction entities by IDs.
func (ciuo *CoinInfoUpdateOne) RemoveTransactionIDs(ids ...int32) *CoinInfoUpdateOne {
	ciuo.mutation.RemoveTransactionIDs(ids...)
	return ciuo
}

// RemoveTransactions removes "transactions" edges to Transaction entities.
func (ciuo *CoinInfoUpdateOne) RemoveTransactions(t ...*Transaction) *CoinInfoUpdateOne {
	ids := make([]int32, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ciuo.RemoveTransactionIDs(ids...)
}

// ClearReviews clears all "reviews" edges to the Review entity.
func (ciuo *CoinInfoUpdateOne) ClearReviews() *CoinInfoUpdateOne {
	ciuo.mutation.ClearReviews()
	return ciuo
}

// RemoveReviewIDs removes the "reviews" edge to Review entities by IDs.
func (ciuo *CoinInfoUpdateOne) RemoveReviewIDs(ids ...int32) *CoinInfoUpdateOne {
	ciuo.mutation.RemoveReviewIDs(ids...)
	return ciuo
}

// RemoveReviews removes "reviews" edges to Review entities.
func (ciuo *CoinInfoUpdateOne) RemoveReviews(r ...*Review) *CoinInfoUpdateOne {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ciuo.RemoveReviewIDs(ids...)
}

// ClearWalletNodes clears all "wallet_nodes" edges to the WalletNode entity.
func (ciuo *CoinInfoUpdateOne) ClearWalletNodes() *CoinInfoUpdateOne {
	ciuo.mutation.ClearWalletNodes()
	return ciuo
}

// RemoveWalletNodeIDs removes the "wallet_nodes" edge to WalletNode entities by IDs.
func (ciuo *CoinInfoUpdateOne) RemoveWalletNodeIDs(ids ...int32) *CoinInfoUpdateOne {
	ciuo.mutation.RemoveWalletNodeIDs(ids...)
	return ciuo
}

// RemoveWalletNodes removes "wallet_nodes" edges to WalletNode entities.
func (ciuo *CoinInfoUpdateOne) RemoveWalletNodes(w ...*WalletNode) *CoinInfoUpdateOne {
	ids := make([]int32, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ciuo.RemoveWalletNodeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *CoinInfoUpdateOne) Select(field string, fields ...string) *CoinInfoUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated CoinInfo entity.
func (ciuo *CoinInfoUpdateOne) Save(ctx context.Context) (*CoinInfo, error) {
	var (
		err  error
		node *CoinInfo
	)
	if len(ciuo.hooks) == 0 {
		if err = ciuo.check(); err != nil {
			return nil, err
		}
		node, err = ciuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ciuo.check(); err != nil {
				return nil, err
			}
			ciuo.mutation = mutation
			node, err = ciuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ciuo.hooks) - 1; i >= 0; i-- {
			if ciuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ciuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ciuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *CoinInfoUpdateOne) SaveX(ctx context.Context) *CoinInfo {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *CoinInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *CoinInfoUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciuo *CoinInfoUpdateOne) check() error {
	if v, ok := ciuo.mutation.Name(); ok {
		if err := coininfo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ciuo.mutation.Unit(); ok {
		if err := coininfo.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf("ent: validator failed for field \"unit\": %w", err)}
		}
	}
	return nil
}

func (ciuo *CoinInfoUpdateOne) sqlSave(ctx context.Context) (_node *CoinInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coininfo.Table,
			Columns: coininfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: coininfo.FieldID,
			},
		},
	}
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CoinInfo.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coininfo.FieldID)
		for _, f := range fields {
			if !coininfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coininfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldName,
		})
	}
	if value, ok := ciuo.mutation.Unit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldUnit,
		})
	}
	if value, ok := ciuo.mutation.NeedSigninfo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coininfo.FieldNeedSigninfo,
		})
	}
	if ciuo.mutation.KeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.KeysTable,
			Columns: []string{coininfo.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: keystore.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.RemovedKeysIDs(); len(nodes) > 0 && !ciuo.mutation.KeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.KeysTable,
			Columns: []string{coininfo.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: keystore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.KeysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.KeysTable,
			Columns: []string{coininfo.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: keystore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciuo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.TransactionsTable,
			Columns: []string{coininfo.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !ciuo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.TransactionsTable,
			Columns: []string{coininfo.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.TransactionsTable,
			Columns: []string{coininfo.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciuo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.ReviewsTable,
			Columns: []string{coininfo.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !ciuo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.ReviewsTable,
			Columns: []string{coininfo.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.ReviewsTable,
			Columns: []string{coininfo.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciuo.mutation.WalletNodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.WalletNodesTable,
			Columns: []string{coininfo.WalletNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: walletnode.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.RemovedWalletNodesIDs(); len(nodes) > 0 && !ciuo.mutation.WalletNodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.WalletNodesTable,
			Columns: []string{coininfo.WalletNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: walletnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.WalletNodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.WalletNodesTable,
			Columns: []string{coininfo.WalletNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: walletnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CoinInfo{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coininfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

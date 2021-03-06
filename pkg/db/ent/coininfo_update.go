// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/predicate"
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

// SetCreatedAt sets the "created_at" field.
func (ciu *CoinInfoUpdate) SetCreatedAt(u uint32) *CoinInfoUpdate {
	ciu.mutation.ResetCreatedAt()
	ciu.mutation.SetCreatedAt(u)
	return ciu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableCreatedAt(u *uint32) *CoinInfoUpdate {
	if u != nil {
		ciu.SetCreatedAt(*u)
	}
	return ciu
}

// AddCreatedAt adds u to the "created_at" field.
func (ciu *CoinInfoUpdate) AddCreatedAt(u int32) *CoinInfoUpdate {
	ciu.mutation.AddCreatedAt(u)
	return ciu
}

// SetUpdatedAt sets the "updated_at" field.
func (ciu *CoinInfoUpdate) SetUpdatedAt(u uint32) *CoinInfoUpdate {
	ciu.mutation.ResetUpdatedAt()
	ciu.mutation.SetUpdatedAt(u)
	return ciu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ciu *CoinInfoUpdate) AddUpdatedAt(u int32) *CoinInfoUpdate {
	ciu.mutation.AddUpdatedAt(u)
	return ciu
}

// SetDeletedAt sets the "deleted_at" field.
func (ciu *CoinInfoUpdate) SetDeletedAt(u uint32) *CoinInfoUpdate {
	ciu.mutation.ResetDeletedAt()
	ciu.mutation.SetDeletedAt(u)
	return ciu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableDeletedAt(u *uint32) *CoinInfoUpdate {
	if u != nil {
		ciu.SetDeletedAt(*u)
	}
	return ciu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ciu *CoinInfoUpdate) AddDeletedAt(u int32) *CoinInfoUpdate {
	ciu.mutation.AddDeletedAt(u)
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

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableUnit(s *string) *CoinInfoUpdate {
	if s != nil {
		ciu.SetUnit(*s)
	}
	return ciu
}

// SetReservedAmount sets the "reserved_amount" field.
func (ciu *CoinInfoUpdate) SetReservedAmount(u uint64) *CoinInfoUpdate {
	ciu.mutation.ResetReservedAmount()
	ciu.mutation.SetReservedAmount(u)
	return ciu
}

// SetNillableReservedAmount sets the "reserved_amount" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableReservedAmount(u *uint64) *CoinInfoUpdate {
	if u != nil {
		ciu.SetReservedAmount(*u)
	}
	return ciu
}

// AddReservedAmount adds u to the "reserved_amount" field.
func (ciu *CoinInfoUpdate) AddReservedAmount(u int64) *CoinInfoUpdate {
	ciu.mutation.AddReservedAmount(u)
	return ciu
}

// SetPreSale sets the "pre_sale" field.
func (ciu *CoinInfoUpdate) SetPreSale(b bool) *CoinInfoUpdate {
	ciu.mutation.SetPreSale(b)
	return ciu
}

// SetNillablePreSale sets the "pre_sale" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillablePreSale(b *bool) *CoinInfoUpdate {
	if b != nil {
		ciu.SetPreSale(*b)
	}
	return ciu
}

// SetLogo sets the "logo" field.
func (ciu *CoinInfoUpdate) SetLogo(s string) *CoinInfoUpdate {
	ciu.mutation.SetLogo(s)
	return ciu
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableLogo(s *string) *CoinInfoUpdate {
	if s != nil {
		ciu.SetLogo(*s)
	}
	return ciu
}

// SetEnv sets the "env" field.
func (ciu *CoinInfoUpdate) SetEnv(s string) *CoinInfoUpdate {
	ciu.mutation.SetEnv(s)
	return ciu
}

// SetNillableEnv sets the "env" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableEnv(s *string) *CoinInfoUpdate {
	if s != nil {
		ciu.SetEnv(*s)
	}
	return ciu
}

// SetForPay sets the "for_pay" field.
func (ciu *CoinInfoUpdate) SetForPay(b bool) *CoinInfoUpdate {
	ciu.mutation.SetForPay(b)
	return ciu
}

// SetNillableForPay sets the "for_pay" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableForPay(b *bool) *CoinInfoUpdate {
	if b != nil {
		ciu.SetForPay(*b)
	}
	return ciu
}

// SetHomePage sets the "home_page" field.
func (ciu *CoinInfoUpdate) SetHomePage(s string) *CoinInfoUpdate {
	ciu.mutation.SetHomePage(s)
	return ciu
}

// SetNillableHomePage sets the "home_page" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableHomePage(s *string) *CoinInfoUpdate {
	if s != nil {
		ciu.SetHomePage(*s)
	}
	return ciu
}

// SetSpecs sets the "specs" field.
func (ciu *CoinInfoUpdate) SetSpecs(s string) *CoinInfoUpdate {
	ciu.mutation.SetSpecs(s)
	return ciu
}

// SetNillableSpecs sets the "specs" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableSpecs(s *string) *CoinInfoUpdate {
	if s != nil {
		ciu.SetSpecs(*s)
	}
	return ciu
}

// Mutation returns the CoinInfoMutation object of the builder.
func (ciu *CoinInfoUpdate) Mutation() *CoinInfoMutation {
	return ciu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *CoinInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ciu.defaults(); err != nil {
		return 0, err
	}
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

// defaults sets the default values of the builder before save.
func (ciu *CoinInfoUpdate) defaults() error {
	if _, ok := ciu.mutation.UpdatedAt(); !ok {
		if coininfo.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coininfo.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coininfo.UpdateDefaultUpdatedAt()
		ciu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ciu *CoinInfoUpdate) check() error {
	if v, ok := ciu.mutation.Name(); ok {
		if err := coininfo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CoinInfo.name": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.Unit(); ok {
		if err := coininfo.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf(`ent: validator failed for field "CoinInfo.unit": %w`, err)}
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
				Type:   field.TypeUUID,
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
	if value, ok := ciu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldCreatedAt,
		})
	}
	if value, ok := ciu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldCreatedAt,
		})
	}
	if value, ok := ciu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldDeletedAt,
		})
	}
	if value, ok := ciu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldDeletedAt,
		})
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
	if value, ok := ciu.mutation.ReservedAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coininfo.FieldReservedAmount,
		})
	}
	if value, ok := ciu.mutation.AddedReservedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coininfo.FieldReservedAmount,
		})
	}
	if value, ok := ciu.mutation.PreSale(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coininfo.FieldPreSale,
		})
	}
	if value, ok := ciu.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldLogo,
		})
	}
	if value, ok := ciu.mutation.Env(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldEnv,
		})
	}
	if value, ok := ciu.mutation.ForPay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coininfo.FieldForPay,
		})
	}
	if value, ok := ciu.mutation.HomePage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldHomePage,
		})
	}
	if value, ok := ciu.mutation.Specs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldSpecs,
		})
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

// SetCreatedAt sets the "created_at" field.
func (ciuo *CoinInfoUpdateOne) SetCreatedAt(u uint32) *CoinInfoUpdateOne {
	ciuo.mutation.ResetCreatedAt()
	ciuo.mutation.SetCreatedAt(u)
	return ciuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableCreatedAt(u *uint32) *CoinInfoUpdateOne {
	if u != nil {
		ciuo.SetCreatedAt(*u)
	}
	return ciuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ciuo *CoinInfoUpdateOne) AddCreatedAt(u int32) *CoinInfoUpdateOne {
	ciuo.mutation.AddCreatedAt(u)
	return ciuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ciuo *CoinInfoUpdateOne) SetUpdatedAt(u uint32) *CoinInfoUpdateOne {
	ciuo.mutation.ResetUpdatedAt()
	ciuo.mutation.SetUpdatedAt(u)
	return ciuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ciuo *CoinInfoUpdateOne) AddUpdatedAt(u int32) *CoinInfoUpdateOne {
	ciuo.mutation.AddUpdatedAt(u)
	return ciuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ciuo *CoinInfoUpdateOne) SetDeletedAt(u uint32) *CoinInfoUpdateOne {
	ciuo.mutation.ResetDeletedAt()
	ciuo.mutation.SetDeletedAt(u)
	return ciuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableDeletedAt(u *uint32) *CoinInfoUpdateOne {
	if u != nil {
		ciuo.SetDeletedAt(*u)
	}
	return ciuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ciuo *CoinInfoUpdateOne) AddDeletedAt(u int32) *CoinInfoUpdateOne {
	ciuo.mutation.AddDeletedAt(u)
	return ciuo
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

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableUnit(s *string) *CoinInfoUpdateOne {
	if s != nil {
		ciuo.SetUnit(*s)
	}
	return ciuo
}

// SetReservedAmount sets the "reserved_amount" field.
func (ciuo *CoinInfoUpdateOne) SetReservedAmount(u uint64) *CoinInfoUpdateOne {
	ciuo.mutation.ResetReservedAmount()
	ciuo.mutation.SetReservedAmount(u)
	return ciuo
}

// SetNillableReservedAmount sets the "reserved_amount" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableReservedAmount(u *uint64) *CoinInfoUpdateOne {
	if u != nil {
		ciuo.SetReservedAmount(*u)
	}
	return ciuo
}

// AddReservedAmount adds u to the "reserved_amount" field.
func (ciuo *CoinInfoUpdateOne) AddReservedAmount(u int64) *CoinInfoUpdateOne {
	ciuo.mutation.AddReservedAmount(u)
	return ciuo
}

// SetPreSale sets the "pre_sale" field.
func (ciuo *CoinInfoUpdateOne) SetPreSale(b bool) *CoinInfoUpdateOne {
	ciuo.mutation.SetPreSale(b)
	return ciuo
}

// SetNillablePreSale sets the "pre_sale" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillablePreSale(b *bool) *CoinInfoUpdateOne {
	if b != nil {
		ciuo.SetPreSale(*b)
	}
	return ciuo
}

// SetLogo sets the "logo" field.
func (ciuo *CoinInfoUpdateOne) SetLogo(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetLogo(s)
	return ciuo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableLogo(s *string) *CoinInfoUpdateOne {
	if s != nil {
		ciuo.SetLogo(*s)
	}
	return ciuo
}

// SetEnv sets the "env" field.
func (ciuo *CoinInfoUpdateOne) SetEnv(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetEnv(s)
	return ciuo
}

// SetNillableEnv sets the "env" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableEnv(s *string) *CoinInfoUpdateOne {
	if s != nil {
		ciuo.SetEnv(*s)
	}
	return ciuo
}

// SetForPay sets the "for_pay" field.
func (ciuo *CoinInfoUpdateOne) SetForPay(b bool) *CoinInfoUpdateOne {
	ciuo.mutation.SetForPay(b)
	return ciuo
}

// SetNillableForPay sets the "for_pay" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableForPay(b *bool) *CoinInfoUpdateOne {
	if b != nil {
		ciuo.SetForPay(*b)
	}
	return ciuo
}

// SetHomePage sets the "home_page" field.
func (ciuo *CoinInfoUpdateOne) SetHomePage(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetHomePage(s)
	return ciuo
}

// SetNillableHomePage sets the "home_page" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableHomePage(s *string) *CoinInfoUpdateOne {
	if s != nil {
		ciuo.SetHomePage(*s)
	}
	return ciuo
}

// SetSpecs sets the "specs" field.
func (ciuo *CoinInfoUpdateOne) SetSpecs(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetSpecs(s)
	return ciuo
}

// SetNillableSpecs sets the "specs" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableSpecs(s *string) *CoinInfoUpdateOne {
	if s != nil {
		ciuo.SetSpecs(*s)
	}
	return ciuo
}

// Mutation returns the CoinInfoMutation object of the builder.
func (ciuo *CoinInfoUpdateOne) Mutation() *CoinInfoMutation {
	return ciuo.mutation
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
	if err := ciuo.defaults(); err != nil {
		return nil, err
	}
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

// defaults sets the default values of the builder before save.
func (ciuo *CoinInfoUpdateOne) defaults() error {
	if _, ok := ciuo.mutation.UpdatedAt(); !ok {
		if coininfo.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coininfo.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coininfo.UpdateDefaultUpdatedAt()
		ciuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ciuo *CoinInfoUpdateOne) check() error {
	if v, ok := ciuo.mutation.Name(); ok {
		if err := coininfo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CoinInfo.name": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.Unit(); ok {
		if err := coininfo.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf(`ent: validator failed for field "CoinInfo.unit": %w`, err)}
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
				Type:   field.TypeUUID,
				Column: coininfo.FieldID,
			},
		},
	}
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinInfo.id" for update`)}
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
	if value, ok := ciuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldCreatedAt,
		})
	}
	if value, ok := ciuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldCreatedAt,
		})
	}
	if value, ok := ciuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldDeletedAt,
		})
	}
	if value, ok := ciuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldDeletedAt,
		})
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
	if value, ok := ciuo.mutation.ReservedAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coininfo.FieldReservedAmount,
		})
	}
	if value, ok := ciuo.mutation.AddedReservedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coininfo.FieldReservedAmount,
		})
	}
	if value, ok := ciuo.mutation.PreSale(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coininfo.FieldPreSale,
		})
	}
	if value, ok := ciuo.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldLogo,
		})
	}
	if value, ok := ciuo.mutation.Env(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldEnv,
		})
	}
	if value, ok := ciuo.mutation.ForPay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coininfo.FieldForPay,
		})
	}
	if value, ok := ciuo.mutation.HomePage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldHomePage,
		})
	}
	if value, ok := ciuo.mutation.Specs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldSpecs,
		})
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

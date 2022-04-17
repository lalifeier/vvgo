// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/data/ent/api"
)

// APICreate is the builder for creating a Api entity.
type APICreate struct {
	config
	mutation *APIMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *APICreate) SetCreatedAt(t time.Time) *APICreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *APICreate) SetNillableCreatedAt(t *time.Time) *APICreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetCreatedBy sets the "created_by" field.
func (ac *APICreate) SetCreatedBy(i int64) *APICreate {
	ac.mutation.SetCreatedBy(i)
	return ac
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ac *APICreate) SetNillableCreatedBy(i *int64) *APICreate {
	if i != nil {
		ac.SetCreatedBy(*i)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *APICreate) SetUpdatedAt(t time.Time) *APICreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *APICreate) SetNillableUpdatedAt(t *time.Time) *APICreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetUpdatedBy sets the "updated_by" field.
func (ac *APICreate) SetUpdatedBy(i int64) *APICreate {
	ac.mutation.SetUpdatedBy(i)
	return ac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ac *APICreate) SetNillableUpdatedBy(i *int64) *APICreate {
	if i != nil {
		ac.SetUpdatedBy(*i)
	}
	return ac
}

// SetGroup sets the "group" field.
func (ac *APICreate) SetGroup(s string) *APICreate {
	ac.mutation.SetGroup(s)
	return ac
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (ac *APICreate) SetNillableGroup(s *string) *APICreate {
	if s != nil {
		ac.SetGroup(*s)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *APICreate) SetName(s string) *APICreate {
	ac.mutation.SetName(s)
	return ac
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ac *APICreate) SetNillableName(s *string) *APICreate {
	if s != nil {
		ac.SetName(*s)
	}
	return ac
}

// SetPath sets the "path" field.
func (ac *APICreate) SetPath(s string) *APICreate {
	ac.mutation.SetPath(s)
	return ac
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (ac *APICreate) SetNillablePath(s *string) *APICreate {
	if s != nil {
		ac.SetPath(*s)
	}
	return ac
}

// SetMethod sets the "method" field.
func (ac *APICreate) SetMethod(s string) *APICreate {
	ac.mutation.SetMethod(s)
	return ac
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (ac *APICreate) SetNillableMethod(s *string) *APICreate {
	if s != nil {
		ac.SetMethod(*s)
	}
	return ac
}

// SetDesc sets the "desc" field.
func (ac *APICreate) SetDesc(s string) *APICreate {
	ac.mutation.SetDesc(s)
	return ac
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (ac *APICreate) SetNillableDesc(s *string) *APICreate {
	if s != nil {
		ac.SetDesc(*s)
	}
	return ac
}

// SetPermission sets the "permission" field.
func (ac *APICreate) SetPermission(s string) *APICreate {
	ac.mutation.SetPermission(s)
	return ac
}

// SetNillablePermission sets the "permission" field if the given value is not nil.
func (ac *APICreate) SetNillablePermission(s *string) *APICreate {
	if s != nil {
		ac.SetPermission(*s)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *APICreate) SetStatus(u uint8) *APICreate {
	ac.mutation.SetStatus(u)
	return ac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ac *APICreate) SetNillableStatus(u *uint8) *APICreate {
	if u != nil {
		ac.SetStatus(*u)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *APICreate) SetID(i int64) *APICreate {
	ac.mutation.SetID(i)
	return ac
}

// Mutation returns the APIMutation object of the builder.
func (ac *APICreate) Mutation() *APIMutation {
	return ac.mutation
}

// Save creates the Api in the database.
func (ac *APICreate) Save(ctx context.Context) (*Api, error) {
	var (
		err  error
		node *Api
	)
	if err := ac.defaults(); err != nil {
		return nil, err
	}
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*APIMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *APICreate) SaveX(ctx context.Context) *Api {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *APICreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *APICreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *APICreate) defaults() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		if api.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized api.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := api.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.CreatedBy(); !ok {
		v := api.DefaultCreatedBy
		ac.mutation.SetCreatedBy(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		if api.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized api.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := api.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedBy(); !ok {
		v := api.DefaultUpdatedBy
		ac.mutation.SetUpdatedBy(v)
	}
	if _, ok := ac.mutation.Group(); !ok {
		v := api.DefaultGroup
		ac.mutation.SetGroup(v)
	}
	if _, ok := ac.mutation.Name(); !ok {
		v := api.DefaultName
		ac.mutation.SetName(v)
	}
	if _, ok := ac.mutation.Path(); !ok {
		v := api.DefaultPath
		ac.mutation.SetPath(v)
	}
	if _, ok := ac.mutation.Method(); !ok {
		v := api.DefaultMethod
		ac.mutation.SetMethod(v)
	}
	if _, ok := ac.mutation.Desc(); !ok {
		v := api.DefaultDesc
		ac.mutation.SetDesc(v)
	}
	if _, ok := ac.mutation.Permission(); !ok {
		v := api.DefaultPermission
		ac.mutation.SetPermission(v)
	}
	if _, ok := ac.mutation.Status(); !ok {
		v := api.DefaultStatus
		ac.mutation.SetStatus(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ac *APICreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Api.created_at"`)}
	}
	if _, ok := ac.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Api.created_by"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Api.updated_at"`)}
	}
	if _, ok := ac.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Api.updated_by"`)}
	}
	if _, ok := ac.mutation.Group(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required field "Api.group"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Api.name"`)}
	}
	if _, ok := ac.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Api.path"`)}
	}
	if _, ok := ac.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "Api.method"`)}
	}
	if _, ok := ac.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "Api.desc"`)}
	}
	if _, ok := ac.mutation.Permission(); !ok {
		return &ValidationError{Name: "permission", err: errors.New(`ent: missing required field "Api.permission"`)}
	}
	if _, ok := ac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Api.status"`)}
	}
	return nil
}

func (ac *APICreate) sqlSave(ctx context.Context) (*Api, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (ac *APICreate) createSpec() (*Api, *sqlgraph.CreateSpec) {
	var (
		_node = &Api{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: api.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: api.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: api.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: api.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: api.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: api.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := ac.mutation.Group(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: api.FieldGroup,
		})
		_node.Group = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: api.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: api.FieldPath,
		})
		_node.Path = value
	}
	if value, ok := ac.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: api.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := ac.mutation.Desc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: api.FieldDesc,
		})
		_node.Desc = value
	}
	if value, ok := ac.mutation.Permission(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: api.FieldPermission,
		})
		_node.Permission = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: api.FieldStatus,
		})
		_node.Status = value
	}
	return _node, _spec
}

// APICreateBulk is the builder for creating many Api entities in bulk.
type APICreateBulk struct {
	config
	builders []*APICreate
}

// Save creates the Api entities in the database.
func (acb *APICreateBulk) Save(ctx context.Context) ([]*Api, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Api, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*APIMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *APICreateBulk) SaveX(ctx context.Context) []*Api {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *APICreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *APICreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
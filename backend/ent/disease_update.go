// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/ADMIN/app/ent/diagnose"
	"github.com/ADMIN/app/ent/disease"
	"github.com/ADMIN/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DiseaseUpdate is the builder for updating Disease entities.
type DiseaseUpdate struct {
	config
	hooks      []Hook
	mutation   *DiseaseMutation
	predicates []predicate.Disease
}

// Where adds a new predicate for the builder.
func (du *DiseaseUpdate) Where(ps ...predicate.Disease) *DiseaseUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDiseaseName sets the Disease_Name field.
func (du *DiseaseUpdate) SetDiseaseName(s string) *DiseaseUpdate {
	du.mutation.SetDiseaseName(s)
	return du
}

// AddDiseaseDiagnoseIDs adds the disease_diagnose edge to Diagnose by ids.
func (du *DiseaseUpdate) AddDiseaseDiagnoseIDs(ids ...int) *DiseaseUpdate {
	du.mutation.AddDiseaseDiagnoseIDs(ids...)
	return du
}

// AddDiseaseDiagnose adds the disease_diagnose edges to Diagnose.
func (du *DiseaseUpdate) AddDiseaseDiagnose(d ...*Diagnose) *DiseaseUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDiseaseDiagnoseIDs(ids...)
}

// Mutation returns the DiseaseMutation object of the builder.
func (du *DiseaseUpdate) Mutation() *DiseaseMutation {
	return du.mutation
}

// RemoveDiseaseDiagnoseIDs removes the disease_diagnose edge to Diagnose by ids.
func (du *DiseaseUpdate) RemoveDiseaseDiagnoseIDs(ids ...int) *DiseaseUpdate {
	du.mutation.RemoveDiseaseDiagnoseIDs(ids...)
	return du
}

// RemoveDiseaseDiagnose removes disease_diagnose edges to Diagnose.
func (du *DiseaseUpdate) RemoveDiseaseDiagnose(d ...*Diagnose) *DiseaseUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDiseaseDiagnoseIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DiseaseUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.DiseaseName(); ok {
		if err := disease.DiseaseNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Disease_Name", err: fmt.Errorf("ent: validator failed for field \"Disease_Name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiseaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DiseaseUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DiseaseUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DiseaseUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DiseaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   disease.Table,
			Columns: disease.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: disease.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.DiseaseName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: disease.FieldDiseaseName,
		})
	}
	if nodes := du.mutation.RemovedDiseaseDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   disease.DiseaseDiagnoseTable,
			Columns: []string{disease.DiseaseDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DiseaseDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   disease.DiseaseDiagnoseTable,
			Columns: []string{disease.DiseaseDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{disease.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DiseaseUpdateOne is the builder for updating a single Disease entity.
type DiseaseUpdateOne struct {
	config
	hooks    []Hook
	mutation *DiseaseMutation
}

// SetDiseaseName sets the Disease_Name field.
func (duo *DiseaseUpdateOne) SetDiseaseName(s string) *DiseaseUpdateOne {
	duo.mutation.SetDiseaseName(s)
	return duo
}

// AddDiseaseDiagnoseIDs adds the disease_diagnose edge to Diagnose by ids.
func (duo *DiseaseUpdateOne) AddDiseaseDiagnoseIDs(ids ...int) *DiseaseUpdateOne {
	duo.mutation.AddDiseaseDiagnoseIDs(ids...)
	return duo
}

// AddDiseaseDiagnose adds the disease_diagnose edges to Diagnose.
func (duo *DiseaseUpdateOne) AddDiseaseDiagnose(d ...*Diagnose) *DiseaseUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDiseaseDiagnoseIDs(ids...)
}

// Mutation returns the DiseaseMutation object of the builder.
func (duo *DiseaseUpdateOne) Mutation() *DiseaseMutation {
	return duo.mutation
}

// RemoveDiseaseDiagnoseIDs removes the disease_diagnose edge to Diagnose by ids.
func (duo *DiseaseUpdateOne) RemoveDiseaseDiagnoseIDs(ids ...int) *DiseaseUpdateOne {
	duo.mutation.RemoveDiseaseDiagnoseIDs(ids...)
	return duo
}

// RemoveDiseaseDiagnose removes disease_diagnose edges to Diagnose.
func (duo *DiseaseUpdateOne) RemoveDiseaseDiagnose(d ...*Diagnose) *DiseaseUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDiseaseDiagnoseIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DiseaseUpdateOne) Save(ctx context.Context) (*Disease, error) {
	if v, ok := duo.mutation.DiseaseName(); ok {
		if err := disease.DiseaseNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Disease_Name", err: fmt.Errorf("ent: validator failed for field \"Disease_Name\": %w", err)}
		}
	}

	var (
		err  error
		node *Disease
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiseaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DiseaseUpdateOne) SaveX(ctx context.Context) *Disease {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DiseaseUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DiseaseUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DiseaseUpdateOne) sqlSave(ctx context.Context) (d *Disease, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   disease.Table,
			Columns: disease.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: disease.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Disease.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.DiseaseName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: disease.FieldDiseaseName,
		})
	}
	if nodes := duo.mutation.RemovedDiseaseDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   disease.DiseaseDiagnoseTable,
			Columns: []string{disease.DiseaseDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DiseaseDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   disease.DiseaseDiagnoseTable,
			Columns: []string{disease.DiseaseDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Disease{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{disease.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/ADMIN/app/ent/diagnose"
	"github.com/ADMIN/app/ent/doctor"
	"github.com/ADMIN/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DoctorUpdate is the builder for updating Doctor entities.
type DoctorUpdate struct {
	config
	hooks      []Hook
	mutation   *DoctorMutation
	predicates []predicate.Doctor
}

// Where adds a new predicate for the builder.
func (du *DoctorUpdate) Where(ps ...predicate.Doctor) *DoctorUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDoctorName sets the Doctor_Name field.
func (du *DoctorUpdate) SetDoctorName(s string) *DoctorUpdate {
	du.mutation.SetDoctorName(s)
	return du
}

// SetDoctorPassword sets the Doctor_Password field.
func (du *DoctorUpdate) SetDoctorPassword(s string) *DoctorUpdate {
	du.mutation.SetDoctorPassword(s)
	return du
}

// SetDoctorEmail sets the Doctor_Email field.
func (du *DoctorUpdate) SetDoctorEmail(s string) *DoctorUpdate {
	du.mutation.SetDoctorEmail(s)
	return du
}

// SetDoctorTel sets the Doctor_tel field.
func (du *DoctorUpdate) SetDoctorTel(s string) *DoctorUpdate {
	du.mutation.SetDoctorTel(s)
	return du
}

// AddDoctorDiagnoseIDs adds the doctor_diagnose edge to Diagnose by ids.
func (du *DoctorUpdate) AddDoctorDiagnoseIDs(ids ...int) *DoctorUpdate {
	du.mutation.AddDoctorDiagnoseIDs(ids...)
	return du
}

// AddDoctorDiagnose adds the doctor_diagnose edges to Diagnose.
func (du *DoctorUpdate) AddDoctorDiagnose(d ...*Diagnose) *DoctorUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDoctorDiagnoseIDs(ids...)
}

// Mutation returns the DoctorMutation object of the builder.
func (du *DoctorUpdate) Mutation() *DoctorMutation {
	return du.mutation
}

// RemoveDoctorDiagnoseIDs removes the doctor_diagnose edge to Diagnose by ids.
func (du *DoctorUpdate) RemoveDoctorDiagnoseIDs(ids ...int) *DoctorUpdate {
	du.mutation.RemoveDoctorDiagnoseIDs(ids...)
	return du
}

// RemoveDoctorDiagnose removes doctor_diagnose edges to Diagnose.
func (du *DoctorUpdate) RemoveDoctorDiagnose(d ...*Diagnose) *DoctorUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDoctorDiagnoseIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DoctorUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.DoctorName(); ok {
		if err := doctor.DoctorNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Doctor_Name", err: fmt.Errorf("ent: validator failed for field \"Doctor_Name\": %w", err)}
		}
	}
	if v, ok := du.mutation.DoctorPassword(); ok {
		if err := doctor.DoctorPasswordValidator(v); err != nil {
			return 0, &ValidationError{Name: "Doctor_Password", err: fmt.Errorf("ent: validator failed for field \"Doctor_Password\": %w", err)}
		}
	}
	if v, ok := du.mutation.DoctorEmail(); ok {
		if err := doctor.DoctorEmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "Doctor_Email", err: fmt.Errorf("ent: validator failed for field \"Doctor_Email\": %w", err)}
		}
	}
	if v, ok := du.mutation.DoctorTel(); ok {
		if err := doctor.DoctorTelValidator(v); err != nil {
			return 0, &ValidationError{Name: "Doctor_tel", err: fmt.Errorf("ent: validator failed for field \"Doctor_tel\": %w", err)}
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
			mutation, ok := m.(*DoctorMutation)
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
func (du *DoctorUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DoctorUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DoctorUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DoctorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
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
	if value, ok := du.mutation.DoctorName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorName,
		})
	}
	if value, ok := du.mutation.DoctorPassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorPassword,
		})
	}
	if value, ok := du.mutation.DoctorEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorEmail,
		})
	}
	if value, ok := du.mutation.DoctorTel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorTel,
		})
	}
	if nodes := du.mutation.RemovedDoctorDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorDiagnoseTable,
			Columns: []string{doctor.DoctorDiagnoseColumn},
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
	if nodes := du.mutation.DoctorDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorDiagnoseTable,
			Columns: []string{doctor.DoctorDiagnoseColumn},
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
			err = &NotFoundError{doctor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DoctorUpdateOne is the builder for updating a single Doctor entity.
type DoctorUpdateOne struct {
	config
	hooks    []Hook
	mutation *DoctorMutation
}

// SetDoctorName sets the Doctor_Name field.
func (duo *DoctorUpdateOne) SetDoctorName(s string) *DoctorUpdateOne {
	duo.mutation.SetDoctorName(s)
	return duo
}

// SetDoctorPassword sets the Doctor_Password field.
func (duo *DoctorUpdateOne) SetDoctorPassword(s string) *DoctorUpdateOne {
	duo.mutation.SetDoctorPassword(s)
	return duo
}

// SetDoctorEmail sets the Doctor_Email field.
func (duo *DoctorUpdateOne) SetDoctorEmail(s string) *DoctorUpdateOne {
	duo.mutation.SetDoctorEmail(s)
	return duo
}

// SetDoctorTel sets the Doctor_tel field.
func (duo *DoctorUpdateOne) SetDoctorTel(s string) *DoctorUpdateOne {
	duo.mutation.SetDoctorTel(s)
	return duo
}

// AddDoctorDiagnoseIDs adds the doctor_diagnose edge to Diagnose by ids.
func (duo *DoctorUpdateOne) AddDoctorDiagnoseIDs(ids ...int) *DoctorUpdateOne {
	duo.mutation.AddDoctorDiagnoseIDs(ids...)
	return duo
}

// AddDoctorDiagnose adds the doctor_diagnose edges to Diagnose.
func (duo *DoctorUpdateOne) AddDoctorDiagnose(d ...*Diagnose) *DoctorUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDoctorDiagnoseIDs(ids...)
}

// Mutation returns the DoctorMutation object of the builder.
func (duo *DoctorUpdateOne) Mutation() *DoctorMutation {
	return duo.mutation
}

// RemoveDoctorDiagnoseIDs removes the doctor_diagnose edge to Diagnose by ids.
func (duo *DoctorUpdateOne) RemoveDoctorDiagnoseIDs(ids ...int) *DoctorUpdateOne {
	duo.mutation.RemoveDoctorDiagnoseIDs(ids...)
	return duo
}

// RemoveDoctorDiagnose removes doctor_diagnose edges to Diagnose.
func (duo *DoctorUpdateOne) RemoveDoctorDiagnose(d ...*Diagnose) *DoctorUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDoctorDiagnoseIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DoctorUpdateOne) Save(ctx context.Context) (*Doctor, error) {
	if v, ok := duo.mutation.DoctorName(); ok {
		if err := doctor.DoctorNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Doctor_Name", err: fmt.Errorf("ent: validator failed for field \"Doctor_Name\": %w", err)}
		}
	}
	if v, ok := duo.mutation.DoctorPassword(); ok {
		if err := doctor.DoctorPasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "Doctor_Password", err: fmt.Errorf("ent: validator failed for field \"Doctor_Password\": %w", err)}
		}
	}
	if v, ok := duo.mutation.DoctorEmail(); ok {
		if err := doctor.DoctorEmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "Doctor_Email", err: fmt.Errorf("ent: validator failed for field \"Doctor_Email\": %w", err)}
		}
	}
	if v, ok := duo.mutation.DoctorTel(); ok {
		if err := doctor.DoctorTelValidator(v); err != nil {
			return nil, &ValidationError{Name: "Doctor_tel", err: fmt.Errorf("ent: validator failed for field \"Doctor_tel\": %w", err)}
		}
	}

	var (
		err  error
		node *Doctor
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorMutation)
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
func (duo *DoctorUpdateOne) SaveX(ctx context.Context) *Doctor {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DoctorUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DoctorUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DoctorUpdateOne) sqlSave(ctx context.Context) (d *Doctor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Doctor.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.DoctorName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorName,
		})
	}
	if value, ok := duo.mutation.DoctorPassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorPassword,
		})
	}
	if value, ok := duo.mutation.DoctorEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorEmail,
		})
	}
	if value, ok := duo.mutation.DoctorTel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorTel,
		})
	}
	if nodes := duo.mutation.RemovedDoctorDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorDiagnoseTable,
			Columns: []string{doctor.DoctorDiagnoseColumn},
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
	if nodes := duo.mutation.DoctorDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorDiagnoseTable,
			Columns: []string{doctor.DoctorDiagnoseColumn},
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
	d = &Doctor{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}

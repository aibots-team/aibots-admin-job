// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-job/ent/task"
	"github.com/suyuan32/simple-admin-job/ent/tasklog"
)

// TaskLogCreate is the builder for creating a TaskLog entity.
type TaskLogCreate struct {
	config
	mutation *TaskLogMutation
	hooks    []Hook
}

// SetStartedAt sets the "started_at" field.
func (tlc *TaskLogCreate) SetStartedAt(t time.Time) *TaskLogCreate {
	tlc.mutation.SetStartedAt(t)
	return tlc
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (tlc *TaskLogCreate) SetNillableStartedAt(t *time.Time) *TaskLogCreate {
	if t != nil {
		tlc.SetStartedAt(*t)
	}
	return tlc
}

// SetFinishedAt sets the "finished_at" field.
func (tlc *TaskLogCreate) SetFinishedAt(t time.Time) *TaskLogCreate {
	tlc.mutation.SetFinishedAt(t)
	return tlc
}

// SetResult sets the "result" field.
func (tlc *TaskLogCreate) SetResult(u uint8) *TaskLogCreate {
	tlc.mutation.SetResult(u)
	return tlc
}

// SetID sets the "id" field.
func (tlc *TaskLogCreate) SetID(u uint64) *TaskLogCreate {
	tlc.mutation.SetID(u)
	return tlc
}

// SetTasksID sets the "tasks" edge to the Task entity by ID.
func (tlc *TaskLogCreate) SetTasksID(id uint64) *TaskLogCreate {
	tlc.mutation.SetTasksID(id)
	return tlc
}

// SetNillableTasksID sets the "tasks" edge to the Task entity by ID if the given value is not nil.
func (tlc *TaskLogCreate) SetNillableTasksID(id *uint64) *TaskLogCreate {
	if id != nil {
		tlc = tlc.SetTasksID(*id)
	}
	return tlc
}

// SetTasks sets the "tasks" edge to the Task entity.
func (tlc *TaskLogCreate) SetTasks(t *Task) *TaskLogCreate {
	return tlc.SetTasksID(t.ID)
}

// Mutation returns the TaskLogMutation object of the builder.
func (tlc *TaskLogCreate) Mutation() *TaskLogMutation {
	return tlc.mutation
}

// Save creates the TaskLog in the database.
func (tlc *TaskLogCreate) Save(ctx context.Context) (*TaskLog, error) {
	tlc.defaults()
	return withHooks[*TaskLog, TaskLogMutation](ctx, tlc.sqlSave, tlc.mutation, tlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tlc *TaskLogCreate) SaveX(ctx context.Context) *TaskLog {
	v, err := tlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tlc *TaskLogCreate) Exec(ctx context.Context) error {
	_, err := tlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlc *TaskLogCreate) ExecX(ctx context.Context) {
	if err := tlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tlc *TaskLogCreate) defaults() {
	if _, ok := tlc.mutation.StartedAt(); !ok {
		v := tasklog.DefaultStartedAt()
		tlc.mutation.SetStartedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tlc *TaskLogCreate) check() error {
	if _, ok := tlc.mutation.StartedAt(); !ok {
		return &ValidationError{Name: "started_at", err: errors.New(`ent: missing required field "TaskLog.started_at"`)}
	}
	if _, ok := tlc.mutation.FinishedAt(); !ok {
		return &ValidationError{Name: "finished_at", err: errors.New(`ent: missing required field "TaskLog.finished_at"`)}
	}
	if _, ok := tlc.mutation.Result(); !ok {
		return &ValidationError{Name: "result", err: errors.New(`ent: missing required field "TaskLog.result"`)}
	}
	return nil
}

func (tlc *TaskLogCreate) sqlSave(ctx context.Context) (*TaskLog, error) {
	if err := tlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	tlc.mutation.id = &_node.ID
	tlc.mutation.done = true
	return _node, nil
}

func (tlc *TaskLogCreate) createSpec() (*TaskLog, *sqlgraph.CreateSpec) {
	var (
		_node = &TaskLog{config: tlc.config}
		_spec = sqlgraph.NewCreateSpec(tasklog.Table, sqlgraph.NewFieldSpec(tasklog.FieldID, field.TypeUint64))
	)
	if id, ok := tlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tlc.mutation.StartedAt(); ok {
		_spec.SetField(tasklog.FieldStartedAt, field.TypeTime, value)
		_node.StartedAt = value
	}
	if value, ok := tlc.mutation.FinishedAt(); ok {
		_spec.SetField(tasklog.FieldFinishedAt, field.TypeTime, value)
		_node.FinishedAt = value
	}
	if value, ok := tlc.mutation.Result(); ok {
		_spec.SetField(tasklog.FieldResult, field.TypeUint8, value)
		_node.Result = value
	}
	if nodes := tlc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasklog.TasksTable,
			Columns: []string{tasklog.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_task_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskLogCreateBulk is the builder for creating many TaskLog entities in bulk.
type TaskLogCreateBulk struct {
	config
	builders []*TaskLogCreate
}

// Save creates the TaskLog entities in the database.
func (tlcb *TaskLogCreateBulk) Save(ctx context.Context) ([]*TaskLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tlcb.builders))
	nodes := make([]*TaskLog, len(tlcb.builders))
	mutators := make([]Mutator, len(tlcb.builders))
	for i := range tlcb.builders {
		func(i int, root context.Context) {
			builder := tlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tlcb *TaskLogCreateBulk) SaveX(ctx context.Context) []*TaskLog {
	v, err := tlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tlcb *TaskLogCreateBulk) Exec(ctx context.Context) error {
	_, err := tlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlcb *TaskLogCreateBulk) ExecX(ctx context.Context) {
	if err := tlcb.Exec(ctx); err != nil {
		panic(err)
	}
}

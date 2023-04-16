// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-job/ent/task"
	"github.com/suyuan32/simple-admin-job/ent/tasklog"
)

// TaskLog is the model entity for the TaskLog schema.
type TaskLog struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Task Started Time | 任务启动时间
	StartedAt time.Time `json:"started_at,omitempty"`
	// Task Finished Time | 任务完成时间
	FinishedAt time.Time `json:"finished_at,omitempty"`
	// The Task Process Result | 任务执行结果
	Result uint8 `json:"result,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskLogQuery when eager-loading is set.
	Edges          TaskLogEdges `json:"edges"`
	task_task_logs *uint64
	selectValues   sql.SelectValues
}

// TaskLogEdges holds the relations/edges for other nodes in the graph.
type TaskLogEdges struct {
	// Tasks holds the value of the tasks edge.
	Tasks *Task `json:"tasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskLogEdges) TasksOrErr() (*Task, error) {
	if e.loadedTypes[0] {
		if e.Tasks == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: task.Label}
		}
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaskLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tasklog.FieldID, tasklog.FieldResult:
			values[i] = new(sql.NullInt64)
		case tasklog.FieldStartedAt, tasklog.FieldFinishedAt:
			values[i] = new(sql.NullTime)
		case tasklog.ForeignKeys[0]: // task_task_logs
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaskLog fields.
func (tl *TaskLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tasklog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tl.ID = uint64(value.Int64)
		case tasklog.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				tl.StartedAt = value.Time
			}
		case tasklog.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				tl.FinishedAt = value.Time
			}
		case tasklog.FieldResult:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				tl.Result = uint8(value.Int64)
			}
		case tasklog.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field task_task_logs", value)
			} else if value.Valid {
				tl.task_task_logs = new(uint64)
				*tl.task_task_logs = uint64(value.Int64)
			}
		default:
			tl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TaskLog.
// This includes values selected through modifiers, order, etc.
func (tl *TaskLog) Value(name string) (ent.Value, error) {
	return tl.selectValues.Get(name)
}

// QueryTasks queries the "tasks" edge of the TaskLog entity.
func (tl *TaskLog) QueryTasks() *TaskQuery {
	return NewTaskLogClient(tl.config).QueryTasks(tl)
}

// Update returns a builder for updating this TaskLog.
// Note that you need to call TaskLog.Unwrap() before calling this method if this TaskLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (tl *TaskLog) Update() *TaskLogUpdateOne {
	return NewTaskLogClient(tl.config).UpdateOne(tl)
}

// Unwrap unwraps the TaskLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tl *TaskLog) Unwrap() *TaskLog {
	_tx, ok := tl.config.driver.(*txDriver)
	if !ok {
		panic("ent: TaskLog is not a transactional entity")
	}
	tl.config.driver = _tx.drv
	return tl
}

// String implements the fmt.Stringer.
func (tl *TaskLog) String() string {
	var builder strings.Builder
	builder.WriteString("TaskLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tl.ID))
	builder.WriteString("started_at=")
	builder.WriteString(tl.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("finished_at=")
	builder.WriteString(tl.FinishedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(fmt.Sprintf("%v", tl.Result))
	builder.WriteByte(')')
	return builder.String()
}

// TaskLogs is a parsable slice of TaskLog.
type TaskLogs []*TaskLog

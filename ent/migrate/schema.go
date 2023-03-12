// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysTasksColumns holds the columns for the "sys_tasks" table.
	SysTasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "task_group", Type: field.TypeString},
		{Name: "cron_expression", Type: field.TypeString},
		{Name: "pattern", Type: field.TypeString},
		{Name: "payload", Type: field.TypeString},
	}
	// SysTasksTable holds the schema information for the "sys_tasks" table.
	SysTasksTable = &schema.Table{
		Name:       "sys_tasks",
		Columns:    SysTasksColumns,
		PrimaryKey: []*schema.Column{SysTasksColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysTasksTable,
	}
)

func init() {
	SysTasksTable.Annotation = &entsql.Annotation{
		Table: "sys_tasks",
	}
}

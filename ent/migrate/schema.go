// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DefautlDbsColumns holds the columns for the "defautl_dbs" table.
	DefautlDbsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "logs", Type: field.TypeString},
		{Name: "accept", Type: field.TypeBool, Default: false},
		{Name: "time_at", Type: field.TypeInt64, Default: 1697492379},
	}
	// DefautlDbsTable holds the schema information for the "defautl_dbs" table.
	DefautlDbsTable = &schema.Table{
		Name:       "defautl_dbs",
		Columns:    DefautlDbsColumns,
		PrimaryKey: []*schema.Column{DefautlDbsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "roles", Type: field.TypeJSON},
		{Name: "manager", Type: field.TypeBool, Default: false},
		{Name: "update_at", Type: field.TypeInt64},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UsersLogsColumns holds the columns for the "users_logs" table.
	UsersLogsColumns = []*schema.Column{
		{Name: "users_id", Type: field.TypeInt},
		{Name: "defautl_db_id", Type: field.TypeInt},
	}
	// UsersLogsTable holds the schema information for the "users_logs" table.
	UsersLogsTable = &schema.Table{
		Name:       "users_logs",
		Columns:    UsersLogsColumns,
		PrimaryKey: []*schema.Column{UsersLogsColumns[0], UsersLogsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_logs_users_id",
				Columns:    []*schema.Column{UsersLogsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "users_logs_defautl_db_id",
				Columns:    []*schema.Column{UsersLogsColumns[1]},
				RefColumns: []*schema.Column{DefautlDbsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DefautlDbsTable,
		UsersTable,
		UsersLogsTable,
	}
)

func init() {
	UsersLogsTable.ForeignKeys[0].RefTable = UsersTable
	UsersLogsTable.ForeignKeys[1].RefTable = DefautlDbsTable
}
// Code generated by ent, DO NOT EDIT.

package defautldb

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the defautldb type in the database.
	Label = "defautl_db"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the userid field in the database.
	FieldUserID = "user_id"
	// FieldLogs holds the string denoting the logs field in the database.
	FieldLogs = "logs"
	// FieldAccept holds the string denoting the accept field in the database.
	FieldAccept = "accept"
	// FieldTimeAt holds the string denoting the timeat field in the database.
	FieldTimeAt = "time_at"
	// EdgeDefaultdb holds the string denoting the defaultdb edge name in mutations.
	EdgeDefaultdb = "defaultdb"
	// Table holds the table name of the defautldb in the database.
	Table = "defautl_dbs"
	// DefaultdbTable is the table that holds the defaultdb relation/edge. The primary key declared below.
	DefaultdbTable = "users_logs"
	// DefaultdbInverseTable is the table name for the Users entity.
	// It exists in this package in order to avoid circular dependency with the "users" package.
	DefaultdbInverseTable = "users"
)

// Columns holds all SQL columns for defautldb fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldLogs,
	FieldAccept,
	FieldTimeAt,
}

var (
	// DefaultdbPrimaryKey and DefaultdbColumn2 are the table columns denoting the
	// primary key for the defaultdb relation (M2M).
	DefaultdbPrimaryKey = []string{"users_id", "defautl_db_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAccept holds the default value on creation for the "accept" field.
	DefaultAccept bool
	// DefaultTimeAt holds the default value on creation for the "timeAt" field.
	DefaultTimeAt int64
)

// OrderOption defines the ordering options for the DefautlDB queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the userID field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByLogs orders the results by the logs field.
func ByLogs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogs, opts...).ToFunc()
}

// ByAccept orders the results by the accept field.
func ByAccept(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccept, opts...).ToFunc()
}

// ByTimeAt orders the results by the timeAt field.
func ByTimeAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeAt, opts...).ToFunc()
}

// ByDefaultdbCount orders the results by defaultdb count.
func ByDefaultdbCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDefaultdbStep(), opts...)
	}
}

// ByDefaultdb orders the results by defaultdb terms.
func ByDefaultdb(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDefaultdbStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDefaultdbStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DefaultdbInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, DefaultdbTable, DefaultdbPrimaryKey...),
	)
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// DefautlDB holds the schema definition for the DefautlDB entity.
type DefautlDB struct {
	ent.Schema
}

// Fields of the DefautlDB.
func (DefautlDB) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("userID"),
		field.String("logs"),
		field.Bool("accept").Default(false),
		field.Int64("timeAt").Default(time.Now().Unix()),
	}
	//return nil
}

// Edges of the DefautlDB.
func (DefautlDB) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("defaultdb", Users.Type).Ref("logs"),
	}
	//return nil
}

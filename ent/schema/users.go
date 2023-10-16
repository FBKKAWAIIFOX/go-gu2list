package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

type Roles []uint64

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("userID"),
		field.JSON("roles", &Roles{}),
		field.Bool("Manager").Default(false),
		field.Int64("updateAt"),
	}
	//return nil
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("logs", DefautlDB.Type),
		//edge.To(),
	}
	//return nil
}

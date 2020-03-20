package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Track holds the schema definition for the Track entity.
type Track struct {
	ent.Schema
}

// Fields of the Track.
func (Track) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive(),
		field.Time("date"),
	}
}

// Edges of the Track.
func (Track) Edges() []ent.Edge {
	return nil
}

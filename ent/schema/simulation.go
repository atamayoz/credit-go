package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Simulation holds the schema definition for the Simulation entity.
type Simulation struct {
	ent.Schema
}

// Fields of the Simulation.
func (Simulation) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount"),
		field.Float("interest"),
		field.Float("periods"),
		field.Float("monthly_payment"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Simulation.
func (Simulation) Edges() []ent.Edge {
	return nil
}

package config

// Relationship defines a relationship between two entities.
type Relationship struct {
	Target Entity
}

// Property is a attribute on an object.
type Property struct {
	PType       string
	Constraints map[string]string
}

// Entity is a the schema of a table, and an object within the system.
// an entity will be created as a table, have a http endpoint and
// can have events on actions.
type Entity struct {
	Relations  []Relationship
	Properties []Property
}

// Config is a definition of the Database, ORM and API layer.
type Config struct {
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"todo/ent/schema"
	"todo/ent/todo"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescTitle is the schema descriptor for title field.
	todoDescTitle := todoFields[0].Descriptor()
	// todo.DefaultTitle holds the default value on creation for the title field.
	todo.DefaultTitle = todoDescTitle.Default.(string)
	// todoDescCompleted is the schema descriptor for completed field.
	todoDescCompleted := todoFields[1].Descriptor()
	// todo.DefaultCompleted holds the default value on creation for the completed field.
	todo.DefaultCompleted = todoDescCompleted.Default.(bool)
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[2].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescPriority is the schema descriptor for priority field.
	todoDescPriority := todoFields[3].Descriptor()
	// todo.PriorityValidator is a validator for the "priority" field. It is called by the builders before save.
	todo.PriorityValidator = todoDescPriority.Validators[0].(func(int) error)
}

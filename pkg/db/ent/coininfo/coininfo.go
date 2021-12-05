// Code generated by entc, DO NOT EDIT.

package coininfo

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the coininfo type in the database.
	Label = "coin_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUnit holds the string denoting the unit field in the database.
	FieldUnit = "unit"
	// FieldPreSale holds the string denoting the pre_sale field in the database.
	FieldPreSale = "pre_sale"
	// FieldLogo holds the string denoting the logo field in the database.
	FieldLogo = "logo"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the coininfo in the database.
	Table = "coin_infos"
)

// Columns holds all SQL columns for coininfo fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldUnit,
	FieldPreSale,
	FieldLogo,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultUnit holds the default value on creation for the "unit" field.
	DefaultUnit string
	// UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	UnitValidator func(string) error
	// DefaultPreSale holds the default value on creation for the "pre_sale" field.
	DefaultPreSale bool
	// DefaultLogo holds the default value on creation for the "logo" field.
	DefaultLogo string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

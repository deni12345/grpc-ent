// Code generated by entc, DO NOT EDIT.

package car

const (
	// Label holds the string label denoting the car type in the database.
	Label = "car"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldManufactureDate holds the string denoting the manufacture_date field in the database.
	FieldManufactureDate = "manufacture_date"
	// FieldCost holds the string denoting the cost field in the database.
	FieldCost = "cost"
	// Table holds the table name of the car in the database.
	Table = "cars"
)

// Columns holds all SQL columns for car fields.
var Columns = []string{
	FieldID,
	FieldModel,
	FieldManufactureDate,
	FieldCost,
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

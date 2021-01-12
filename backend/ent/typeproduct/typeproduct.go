// Code generated by entc, DO NOT EDIT.

package typeproduct

const (
	// Label holds the string label denoting the typeproduct type in the database.
	Label = "typeproduct"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTypeproductname holds the string denoting the typeproductname field in the database.
	FieldTypeproductname = "typeproductname"

	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"

	// Table holds the table name of the typeproduct in the database.
	Table = "typeproducts"
	// ProductTable is the table the holds the product relation/edge.
	ProductTable = "products"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "Typeproduct"
)

// Columns holds all SQL columns for typeproduct fields.
var Columns = []string{
	FieldID,
	FieldTypeproductname,
}

// Code generated by entc, DO NOT EDIT.

package receipt

import (
	"time"
)

const (
	// Label holds the string label denoting the receipt type in the database.
	Label = "receipt"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDateTime holds the string denoting the date_time field in the database.
	FieldDateTime = "date_time"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldProductname holds the string denoting the productname field in the database.
	FieldProductname = "productname"
	// FieldReceiptcode holds the string denoting the receiptcode field in the database.
	FieldReceiptcode = "receiptcode"

	// EdgePaymenttype holds the string denoting the paymenttype edge name in mutations.
	EdgePaymenttype = "paymenttype"
	// EdgeAdminrepair holds the string denoting the adminrepair edge name in mutations.
	EdgeAdminrepair = "adminrepair"
	// EdgePersonal holds the string denoting the personal edge name in mutations.
	EdgePersonal = "personal"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"

	// Table holds the table name of the receipt in the database.
	Table = "receipts"
	// PaymenttypeTable is the table the holds the paymenttype relation/edge.
	PaymenttypeTable = "receipts"
	// PaymenttypeInverseTable is the table name for the PaymentType entity.
	// It exists in this package in order to avoid circular dependency with the "paymenttype" package.
	PaymenttypeInverseTable = "payment_types"
	// PaymenttypeColumn is the table column denoting the paymenttype relation/edge.
	PaymenttypeColumn = "paymenttype_id"
	// AdminrepairTable is the table the holds the adminrepair relation/edge.
	AdminrepairTable = "receipts"
	// AdminrepairInverseTable is the table name for the Adminrepair entity.
	// It exists in this package in order to avoid circular dependency with the "adminrepair" package.
	AdminrepairInverseTable = "adminrepairs"
	// AdminrepairColumn is the table column denoting the adminrepair relation/edge.
	AdminrepairColumn = "adminrepair_id"
	// PersonalTable is the table the holds the personal relation/edge.
	PersonalTable = "receipts"
	// PersonalInverseTable is the table name for the Personal entity.
	// It exists in this package in order to avoid circular dependency with the "personal" package.
	PersonalInverseTable = "personals"
	// PersonalColumn is the table column denoting the personal relation/edge.
	PersonalColumn = "personal_id"
	// CustomerTable is the table the holds the customer relation/edge.
	CustomerTable = "receipts"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_id"
	// ProductTable is the table the holds the product relation/edge.
	ProductTable = "receipts"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "product_id"
)

// Columns holds all SQL columns for receipt fields.
var Columns = []string{
	FieldID,
	FieldDateTime,
	FieldAddress,
	FieldProductname,
	FieldReceiptcode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Receipt type.
var ForeignKeys = []string{
	"adminrepair_id",
	"customer_id",
	"paymenttype_id",
	"personal_id",
	"product_id",
}

var (
	// DefaultDateTime holds the default value on creation for the date_time field.
	DefaultDateTime func() time.Time
	// AddressValidator is a validator for the "Address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// ProductnameValidator is a validator for the "Productname" field. It is called by the builders before save.
	ProductnameValidator func(string) error
	// ReceiptcodeValidator is a validator for the "Receiptcode" field. It is called by the builders before save.
	ReceiptcodeValidator func(string) error
)

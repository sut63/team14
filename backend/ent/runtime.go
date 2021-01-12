// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/tanapon395/playlist-video/ent/product"
	"github.com/tanapon395/playlist-video/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescProductname is the schema descriptor for Productname field.
	productDescProductname := productFields[0].Descriptor()
	// product.ProductnameValidator is a validator for the "Productname" field. It is called by the builders before save.
	product.ProductnameValidator = productDescProductname.Validators[0].(func(string) error)
	// productDescNumberofproduct is the schema descriptor for Numberofproduct field.
	productDescNumberofproduct := productFields[1].Descriptor()
	// product.NumberofproductValidator is a validator for the "Numberofproduct" field. It is called by the builders before save.
	product.NumberofproductValidator = productDescNumberofproduct.Validators[0].(func(string) error)
	// productDescPrice is the schema descriptor for Price field.
	productDescPrice := productFields[2].Descriptor()
	// product.PriceValidator is a validator for the "Price" field. It is called by the builders before save.
	product.PriceValidator = productDescPrice.Validators[0].(func(string) error)
}

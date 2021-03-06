// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/adminrepair"
	"github.com/tanapon395/playlist-video/ent/fix"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/product"
)

// Adminrepair is the model entity for the Adminrepair schema.
type Adminrepair struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Numberrepair holds the value of the "numberrepair" field.
	Numberrepair string `json:"numberrepair,omitempty"`
	// Equipmentdamate holds the value of the "equipmentdamate" field.
	Equipmentdamate string `json:"equipmentdamate,omitempty"`
	// Repairinformation holds the value of the "repairinformation" field.
	Repairinformation string `json:"repairinformation,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdminrepairQuery when eager-loading is set.
	Edges       AdminrepairEdges `json:"edges"`
	fix_id      *int
	personal_id *int
	product_id  *int
}

// AdminrepairEdges holds the relations/edges for other nodes in the graph.
type AdminrepairEdges struct {
	// Receipt holds the value of the receipt edge.
	Receipt []*Receipt
	// AdminrepairPersonal holds the value of the AdminrepairPersonal edge.
	AdminrepairPersonal *Personal
	// AdminrepairFix holds the value of the AdminrepairFix edge.
	AdminrepairFix *Fix
	// AdminrepairProduct holds the value of the AdminrepairProduct edge.
	AdminrepairProduct *Product
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ReceiptOrErr returns the Receipt value or an error if the edge
// was not loaded in eager-loading.
func (e AdminrepairEdges) ReceiptOrErr() ([]*Receipt, error) {
	if e.loadedTypes[0] {
		return e.Receipt, nil
	}
	return nil, &NotLoadedError{edge: "receipt"}
}

// AdminrepairPersonalOrErr returns the AdminrepairPersonal value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdminrepairEdges) AdminrepairPersonalOrErr() (*Personal, error) {
	if e.loadedTypes[1] {
		if e.AdminrepairPersonal == nil {
			// The edge AdminrepairPersonal was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: personal.Label}
		}
		return e.AdminrepairPersonal, nil
	}
	return nil, &NotLoadedError{edge: "AdminrepairPersonal"}
}

// AdminrepairFixOrErr returns the AdminrepairFix value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdminrepairEdges) AdminrepairFixOrErr() (*Fix, error) {
	if e.loadedTypes[2] {
		if e.AdminrepairFix == nil {
			// The edge AdminrepairFix was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: fix.Label}
		}
		return e.AdminrepairFix, nil
	}
	return nil, &NotLoadedError{edge: "AdminrepairFix"}
}

// AdminrepairProductOrErr returns the AdminrepairProduct value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdminrepairEdges) AdminrepairProductOrErr() (*Product, error) {
	if e.loadedTypes[3] {
		if e.AdminrepairProduct == nil {
			// The edge AdminrepairProduct was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.AdminrepairProduct, nil
	}
	return nil, &NotLoadedError{edge: "AdminrepairProduct"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Adminrepair) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // numberrepair
		&sql.NullString{}, // equipmentdamate
		&sql.NullString{}, // repairinformation
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Adminrepair) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // fix_id
		&sql.NullInt64{}, // personal_id
		&sql.NullInt64{}, // product_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Adminrepair fields.
func (a *Adminrepair) assignValues(values ...interface{}) error {
	if m, n := len(values), len(adminrepair.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field numberrepair", values[0])
	} else if value.Valid {
		a.Numberrepair = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field equipmentdamate", values[1])
	} else if value.Valid {
		a.Equipmentdamate = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field repairinformation", values[2])
	} else if value.Valid {
		a.Repairinformation = value.String
	}
	values = values[3:]
	if len(values) == len(adminrepair.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field fix_id", value)
		} else if value.Valid {
			a.fix_id = new(int)
			*a.fix_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field personal_id", value)
		} else if value.Valid {
			a.personal_id = new(int)
			*a.personal_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field product_id", value)
		} else if value.Valid {
			a.product_id = new(int)
			*a.product_id = int(value.Int64)
		}
	}
	return nil
}

// QueryReceipt queries the receipt edge of the Adminrepair.
func (a *Adminrepair) QueryReceipt() *ReceiptQuery {
	return (&AdminrepairClient{config: a.config}).QueryReceipt(a)
}

// QueryAdminrepairPersonal queries the AdminrepairPersonal edge of the Adminrepair.
func (a *Adminrepair) QueryAdminrepairPersonal() *PersonalQuery {
	return (&AdminrepairClient{config: a.config}).QueryAdminrepairPersonal(a)
}

// QueryAdminrepairFix queries the AdminrepairFix edge of the Adminrepair.
func (a *Adminrepair) QueryAdminrepairFix() *FixQuery {
	return (&AdminrepairClient{config: a.config}).QueryAdminrepairFix(a)
}

// QueryAdminrepairProduct queries the AdminrepairProduct edge of the Adminrepair.
func (a *Adminrepair) QueryAdminrepairProduct() *ProductQuery {
	return (&AdminrepairClient{config: a.config}).QueryAdminrepairProduct(a)
}

// Update returns a builder for updating this Adminrepair.
// Note that, you need to call Adminrepair.Unwrap() before calling this method, if this Adminrepair
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Adminrepair) Update() *AdminrepairUpdateOne {
	return (&AdminrepairClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Adminrepair) Unwrap() *Adminrepair {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Adminrepair is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Adminrepair) String() string {
	var builder strings.Builder
	builder.WriteString("Adminrepair(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", numberrepair=")
	builder.WriteString(a.Numberrepair)
	builder.WriteString(", equipmentdamate=")
	builder.WriteString(a.Equipmentdamate)
	builder.WriteString(", repairinformation=")
	builder.WriteString(a.Repairinformation)
	builder.WriteByte(')')
	return builder.String()
}

// Adminrepairs is a parsable slice of Adminrepair.
type Adminrepairs []*Adminrepair

func (a Adminrepairs) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}

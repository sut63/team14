// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/adminrepair"
	"github.com/tanapon395/playlist-video/ent/customer"
	"github.com/tanapon395/playlist-video/ent/paymenttype"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/product"
	"github.com/tanapon395/playlist-video/ent/receipt"
)

// Receipt is the model entity for the Receipt schema.
type Receipt struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AddedTime holds the value of the "added_time" field.
	AddedTime time.Time `json:"added_time,omitempty"`
	// Serviceprovider holds the value of the "Serviceprovider" field.
	Serviceprovider string `json:"Serviceprovider,omitempty"`
	// Address holds the value of the "Address" field.
	Address string `json:"Address,omitempty"`
	// Productname holds the value of the "Productname" field.
	Productname string `json:"Productname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReceiptQuery when eager-loading is set.
	Edges          ReceiptEdges `json:"edges"`
	adminrepair_id *int
	customer_id    *int
	paymenttype_id *int
	personal_id    *int
	product_id     *int
}

// ReceiptEdges holds the relations/edges for other nodes in the graph.
type ReceiptEdges struct {
	// Paymenttype holds the value of the paymenttype edge.
	Paymenttype *PaymentType
	// Adminrepair holds the value of the adminrepair edge.
	Adminrepair *Adminrepair
	// Personal holds the value of the personal edge.
	Personal *Personal
	// Customer holds the value of the customer edge.
	Customer *Customer
	// Product holds the value of the product edge.
	Product *Product
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// PaymenttypeOrErr returns the Paymenttype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReceiptEdges) PaymenttypeOrErr() (*PaymentType, error) {
	if e.loadedTypes[0] {
		if e.Paymenttype == nil {
			// The edge paymenttype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: paymenttype.Label}
		}
		return e.Paymenttype, nil
	}
	return nil, &NotLoadedError{edge: "paymenttype"}
}

// AdminrepairOrErr returns the Adminrepair value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReceiptEdges) AdminrepairOrErr() (*Adminrepair, error) {
	if e.loadedTypes[1] {
		if e.Adminrepair == nil {
			// The edge adminrepair was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: adminrepair.Label}
		}
		return e.Adminrepair, nil
	}
	return nil, &NotLoadedError{edge: "adminrepair"}
}

// PersonalOrErr returns the Personal value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReceiptEdges) PersonalOrErr() (*Personal, error) {
	if e.loadedTypes[2] {
		if e.Personal == nil {
			// The edge personal was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: personal.Label}
		}
		return e.Personal, nil
	}
	return nil, &NotLoadedError{edge: "personal"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReceiptEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[3] {
		if e.Customer == nil {
			// The edge customer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReceiptEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[4] {
		if e.Product == nil {
			// The edge product was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Receipt) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // added_time
		&sql.NullString{}, // Serviceprovider
		&sql.NullString{}, // Address
		&sql.NullString{}, // Productname
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Receipt) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // adminrepair_id
		&sql.NullInt64{}, // customer_id
		&sql.NullInt64{}, // paymenttype_id
		&sql.NullInt64{}, // personal_id
		&sql.NullInt64{}, // product_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Receipt fields.
func (r *Receipt) assignValues(values ...interface{}) error {
	if m, n := len(values), len(receipt.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field added_time", values[0])
	} else if value.Valid {
		r.AddedTime = value.Time
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Serviceprovider", values[1])
	} else if value.Valid {
		r.Serviceprovider = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Address", values[2])
	} else if value.Valid {
		r.Address = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Productname", values[3])
	} else if value.Valid {
		r.Productname = value.String
	}
	values = values[4:]
	if len(values) == len(receipt.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field adminrepair_id", value)
		} else if value.Valid {
			r.adminrepair_id = new(int)
			*r.adminrepair_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field customer_id", value)
		} else if value.Valid {
			r.customer_id = new(int)
			*r.customer_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field paymenttype_id", value)
		} else if value.Valid {
			r.paymenttype_id = new(int)
			*r.paymenttype_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field personal_id", value)
		} else if value.Valid {
			r.personal_id = new(int)
			*r.personal_id = int(value.Int64)
		}
		if value, ok := values[4].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field product_id", value)
		} else if value.Valid {
			r.product_id = new(int)
			*r.product_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPaymenttype queries the paymenttype edge of the Receipt.
func (r *Receipt) QueryPaymenttype() *PaymentTypeQuery {
	return (&ReceiptClient{config: r.config}).QueryPaymenttype(r)
}

// QueryAdminrepair queries the adminrepair edge of the Receipt.
func (r *Receipt) QueryAdminrepair() *AdminrepairQuery {
	return (&ReceiptClient{config: r.config}).QueryAdminrepair(r)
}

// QueryPersonal queries the personal edge of the Receipt.
func (r *Receipt) QueryPersonal() *PersonalQuery {
	return (&ReceiptClient{config: r.config}).QueryPersonal(r)
}

// QueryCustomer queries the customer edge of the Receipt.
func (r *Receipt) QueryCustomer() *CustomerQuery {
	return (&ReceiptClient{config: r.config}).QueryCustomer(r)
}

// QueryProduct queries the product edge of the Receipt.
func (r *Receipt) QueryProduct() *ProductQuery {
	return (&ReceiptClient{config: r.config}).QueryProduct(r)
}

// Update returns a builder for updating this Receipt.
// Note that, you need to call Receipt.Unwrap() before calling this method, if this Receipt
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Receipt) Update() *ReceiptUpdateOne {
	return (&ReceiptClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Receipt) Unwrap() *Receipt {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Receipt is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Receipt) String() string {
	var builder strings.Builder
	builder.WriteString("Receipt(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", added_time=")
	builder.WriteString(r.AddedTime.Format(time.ANSIC))
	builder.WriteString(", Serviceprovider=")
	builder.WriteString(r.Serviceprovider)
	builder.WriteString(", Address=")
	builder.WriteString(r.Address)
	builder.WriteString(", Productname=")
	builder.WriteString(r.Productname)
	builder.WriteByte(')')
	return builder.String()
}

// Receipts is a parsable slice of Receipt.
type Receipts []*Receipt

func (r Receipts) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}

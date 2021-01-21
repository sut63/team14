package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Adminrepair holds the schema definition for the Adminrepair entity.
type Adminrepair struct {
	ent.Schema
}

// Fields of the Adminrepair.
func (Adminrepair) Fields() []ent.Field {
	return []ent.Field{
		field.String("numberrepair").Unique().MaxLen(6).Validate(func(amp string) error {
			match, _ := regexp.MatchString("[A]+[M]+[P]+[-]+[0-9]+[0-9]", amp)
			if !match {
				return errors.New("รูปแบบรหัสบัตรคิวของ บันทึกซ่อมแซมคอมพิวเตอร์ของพนักงงานไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("equipmentdamate").MaxLen(75).MinLen(5),
		field.String("repairinformation").MaxLen(500).MinLen(10),
	}
}

// Edges of the Adminrepair.
func (Adminrepair) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("receipt", Receipt.Type).StorageKey(edge.Column("adminrepair_id")),
		edge.From("AdminrepairPersonal", Personal.Type).Ref("personal").Unique(),
		edge.From("AdminrepairFix", Fix.Type).Ref("fix").Unique(),
		edge.From("AdminrepairProduct", Product.Type).Ref("product").Unique(),
	}
}

// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/tanapon395/playlist-video/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type decisionCtxKey struct{}

// DecisionContext creates a decision context.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

func decisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AdminrepairQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AdminrepairQueryRuleFunc func(context.Context, *ent.AdminrepairQuery) error

// EvalQuery return f(ctx, q).
func (f AdminrepairQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AdminrepairQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AdminrepairQuery", q)
}

// The AdminrepairMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AdminrepairMutationRuleFunc func(context.Context, *ent.AdminrepairMutation) error

// EvalMutation calls f(ctx, m).
func (f AdminrepairMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AdminrepairMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AdminrepairMutation", m)
}

// The BrandQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BrandQueryRuleFunc func(context.Context, *ent.BrandQuery) error

// EvalQuery return f(ctx, q).
func (f BrandQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BrandQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BrandQuery", q)
}

// The BrandMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BrandMutationRuleFunc func(context.Context, *ent.BrandMutation) error

// EvalMutation calls f(ctx, m).
func (f BrandMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BrandMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BrandMutation", m)
}

// The CustomerQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CustomerQueryRuleFunc func(context.Context, *ent.CustomerQuery) error

// EvalQuery return f(ctx, q).
func (f CustomerQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CustomerQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CustomerQuery", q)
}

// The CustomerMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CustomerMutationRuleFunc func(context.Context, *ent.CustomerMutation) error

// EvalMutation calls f(ctx, m).
func (f CustomerMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CustomerMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CustomerMutation", m)
}

// The DepartmentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DepartmentQueryRuleFunc func(context.Context, *ent.DepartmentQuery) error

// EvalQuery return f(ctx, q).
func (f DepartmentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DepartmentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DepartmentQuery", q)
}

// The DepartmentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DepartmentMutationRuleFunc func(context.Context, *ent.DepartmentMutation) error

// EvalMutation calls f(ctx, m).
func (f DepartmentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DepartmentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DepartmentMutation", m)
}

// The FixQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FixQueryRuleFunc func(context.Context, *ent.FixQuery) error

// EvalQuery return f(ctx, q).
func (f FixQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FixQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FixQuery", q)
}

// The FixMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FixMutationRuleFunc func(context.Context, *ent.FixMutation) error

// EvalMutation calls f(ctx, m).
func (f FixMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FixMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FixMutation", m)
}

// The GenderQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GenderQueryRuleFunc func(context.Context, *ent.GenderQuery) error

// EvalQuery return f(ctx, q).
func (f GenderQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GenderQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GenderQuery", q)
}

// The GenderMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GenderMutationRuleFunc func(context.Context, *ent.GenderMutation) error

// EvalMutation calls f(ctx, m).
func (f GenderMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GenderMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GenderMutation", m)
}

// The PersonalQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PersonalQueryRuleFunc func(context.Context, *ent.PersonalQuery) error

// EvalQuery return f(ctx, q).
func (f PersonalQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PersonalQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PersonalQuery", q)
}

// The PersonalMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PersonalMutationRuleFunc func(context.Context, *ent.PersonalMutation) error

// EvalMutation calls f(ctx, m).
func (f PersonalMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PersonalMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PersonalMutation", m)
}

// The ProductQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProductQueryRuleFunc func(context.Context, *ent.ProductQuery) error

// EvalQuery return f(ctx, q).
func (f ProductQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProductQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProductQuery", q)
}

// The ProductMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProductMutationRuleFunc func(context.Context, *ent.ProductMutation) error

// EvalMutation calls f(ctx, m).
func (f ProductMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProductMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProductMutation", m)
}

// The ReceiptQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ReceiptQueryRuleFunc func(context.Context, *ent.ReceiptQuery) error

// EvalQuery return f(ctx, q).
func (f ReceiptQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ReceiptQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ReceiptQuery", q)
}

// The ReceiptMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ReceiptMutationRuleFunc func(context.Context, *ent.ReceiptMutation) error

// EvalMutation calls f(ctx, m).
func (f ReceiptMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ReceiptMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ReceiptMutation", m)
}

// The TitleQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TitleQueryRuleFunc func(context.Context, *ent.TitleQuery) error

// EvalQuery return f(ctx, q).
func (f TitleQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TitleQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TitleQuery", q)
}

// The TitleMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TitleMutationRuleFunc func(context.Context, *ent.TitleMutation) error

// EvalMutation calls f(ctx, m).
func (f TitleMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TitleMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TitleMutation", m)
}

// The TypeproductQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TypeproductQueryRuleFunc func(context.Context, *ent.TypeproductQuery) error

// EvalQuery return f(ctx, q).
func (f TypeproductQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TypeproductQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TypeproductQuery", q)
}

// The TypeproductMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TypeproductMutationRuleFunc func(context.Context, *ent.TypeproductMutation) error

// EvalMutation calls f(ctx, m).
func (f TypeproductMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TypeproductMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TypeproductMutation", m)
}

// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package asterank

import "fmt"

// Parameter defines a query filter for a single field.
type Parameter struct {
	Field   string      // Name of the field to filter on.
	Compare string      // Type of comparison to perform.
	Value   interface{} // Scalar value to compare to.
}

func newParameter(field, op string, value interface{}) Parameter {
	return Parameter{field, op, value}
}

// LT matches values that are less than a specified value.
func LT(field string, value interface{}) Parameter {
	return newParameter(field, "$lt", value)
}

// LTE matches values that are less than or equal to a specified value.
func LTE(field string, value interface{}) Parameter {
	return newParameter(field, "$lte", value)
}

// GT matches values that are greater than a specified value.
func GT(field string, value interface{}) Parameter {
	return newParameter(field, "$gt", value)
}

// GTE matches values that are greater than or equal to a specified value.
func GTE(field string, value interface{}) Parameter {
	return newParameter(field, "$gte", value)
}

// EQ matches values that are equal to a specified value.
func EQ(field string, value interface{}) Parameter {
	return newParameter(field, "$eq", value)
}

// NE matches all values that are not equal to a specified value.
func NE(field string, value interface{}) Parameter {
	return newParameter(field, "$ne", value)
}

func (p *Parameter) String() string {
	switch v := p.Value.(type) {
	case string:
		return fmt.Sprintf("%q:{%q:%q}", p.Field, p.Compare, v)
	case int:
		return fmt.Sprintf("%q:{%q:%v}", p.Field, p.Compare, float64(v))
	case int64:
		return fmt.Sprintf("%q:{%q:%v}", p.Field, p.Compare, float64(v))
	case float64:
		return fmt.Sprintf("%q:{%q:%.f}", p.Field, p.Compare, v)
	}

	return fmt.Sprintf("%q:{%q:%v}", p.Field, p.Compare, p.Value)
}

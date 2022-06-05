package decimal

import (
	"fmt"
	"io"
)

// Implement interfaces to add support for automatic
// type conversion between GraphQL <-> Go Types

// interfaces are implicit in Go. Removed to avoid
// graphql package dependency here.
//var _ graphql.Marshaler = (*Decimal)(nil)
//var _ graphql.Unmarshaler = (*Decimal)(nil)

func (d Decimal) MarshalGQL(w io.Writer) {
	w.Write([]byte(d.String()))
}

func (d *Decimal) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("type Decimal value must be a string in GraphQL, got: %+v", v)
	}
	dec, err := NewFromString(value)
	if err != nil {
		return fmt.Errorf("error parsing Decimal in GraphQL: %+v", err)
	}
	*d = dec

	return nil
}

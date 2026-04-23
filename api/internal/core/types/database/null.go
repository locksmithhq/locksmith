package database

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Null struct {
	value any
	valid bool
}

func (n Null) IsNull() bool {
	return !n.valid
}

func (n *Null) Scan(value any) error {
	if value == nil {
		n.value = nil
		n.valid = false
		return nil
	}
	n.value = value
	n.valid = true
	return nil
}

func (n Null) MarshalJSON() ([]byte, error) {
	if !n.valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.value)
}

func ParseNull(value string) Null {
	return Null{
		value: value,
		valid: value != "",
	}
}

// Value implements the driver.Valuer interface.
func (n Null) Value() (driver.Value, error) {
	if !n.valid {
		return nil, nil
	}
	return n.value, nil
}

// String returns the underlying value as a string.
// Returns empty string if the value is null.
func (n Null) String() string {
	if !n.valid || n.value == nil {
		return ""
	}
	switch v := n.value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}



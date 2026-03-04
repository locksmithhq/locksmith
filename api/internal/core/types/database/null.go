package database

import (
	"database/sql/driver"
	"encoding/json"
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



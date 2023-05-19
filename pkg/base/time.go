package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

// MarshalJSON implements the json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, t.UTC().Format(time.RFC3339))), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" || s == "null" {
		t.Time = time.Time{}
		return nil
	}
	parsedTime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return errors.New("invalid RFC3339 format")
	}
	t.Time = parsedTime
	return nil
}

func (c *Time) UnmarshalValue(value interface{}) error {
	g, ok := value.(*gvar.Var)
	if !ok {
		return nil
	}
	c.Time = g.Time()
	return nil
}

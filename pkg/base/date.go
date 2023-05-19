package base

import (
	"github.com/gogf/gf/v2/container/gvar"
	"time"
)

type Date struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
func (t Date) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Format(time.DateOnly) + `"`), nil
}

func (t Date) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.DateOnly)
}

// UnmarshalJSON parses a JSON string into the Time struct
func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		d.Time = time.Time{}
		return nil
	}

	newTime, err := time.Parse(time.DateOnly, string(data))
	if err != nil {
		return err
	}

	d.Time = newTime
	return nil
}

func (c *Date) UnmarshalValue(value interface{}) error {
	g, ok := value.(*gvar.Var)
	if !ok {
		return nil
	}
	c.Time = g.Time()
	return nil
}

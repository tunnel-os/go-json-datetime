package jsondt

import (
	"encoding/json"
	"fmt"
	"time"
)

var _ json.Marshaler = (*Date)(nil)
var _ json.Unmarshaler = (*Date)(nil)

type Date struct {
	time.Time
}

func (t Date) String() string {
	return t.Time.Format("2006-01-02")
}

func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *Date) UnmarshalJSON(b []byte) error {
	if isEmpty(b) {
		return nil
	}

	if len(b) != 12 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("%w: failed to unmarshal non-string value %q as an YYYY-MM-dd", ErrJSONDateTime, b)
	}

	now := time.Now()
	tm, err := time.ParseInLocation("2006-01-02", string(b[1:len(b)-1]), now.Location())
	if err != nil {
		return err
	}
	*t = Date{tm}
	return nil
}

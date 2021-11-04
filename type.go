package int64option

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Type struct {
	hasValue bool
	value    int64
}

var _ fmt.GoStringer = Type{}
var _ fmt.Stringer = Type{}
var _ json.Marshaler = Type{}
var _ json.Unmarshaler = &Type{}
var _ driver.Valuer = Type{}
var _ sql.Scanner = &Type{}

func Nothing() Type {
	return Type{}
}

func Something(value int64) Type {
	return Type{true, value}
}

func (receiver Type) Return() (int64, error) {
	if !receiver.hasValue {
		return 0, errors.New("there are no values")
	}
	return receiver.value, nil
}

func (receiver Type) GoString() string {
	if !receiver.hasValue {
		return "int64option.Nothing()"
	}

	return fmt.Sprintf("int64option.Something(%d)", receiver.value)
}

func (receiver Type) String() string {
	if !receiver.hasValue {
		return "⧼nothing⧽"
	}

	return fmt.Sprintf("int64option.Something(%d)", receiver.value)
}

func (receiver Type) toString() string {
	if !receiver.hasValue {
		return "nothing()"
	}
	return fmt.Sprintf("something(%d)", receiver.value)
}

func (receiver Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.toString())
}

var somethingRegExp = regexp.MustCompile(`^something\(\d+\)$`)
var nothingRegExp = regexp.MustCompile(`^nothing\(\)$`)
var numRegExp = regexp.MustCompile(`\d+`)

func parseString(str string) (Type, error) {
	if somethingRegExp.MatchString(str) {
		v, err := strconv.Atoi(numRegExp.FindString(str))
		if err != nil {
			return Type{}, err
		}
		return Something(int64(v)), nil
	} else if nothingRegExp.MatchString(str) {
		return Nothing(), nil
	}
	return Type{}, fmt.Errorf("invalid type string %q", str)
}

func (receiver *Type) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	t, err := parseString(str)
	if err != nil {
		return err
	}
	*receiver = t
	return nil
}

func (receiver Type) Value() (driver.Value, error) {
	return receiver.toString(), nil
}

func (receiver *Type) Scan(src interface{}) error {
	switch casted := src.(type) {
	case string:
		t, err := parseString(casted)
		if err != nil {
			return err
		}
		*receiver = t
		return nil
	case []byte:
		t, err := parseString(string(casted))
		if err != nil {
			return err
		}
		*receiver = t
		return nil
	default:
		return errors.New("failed to scan item. Not a valid type")
	}
}

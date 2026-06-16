package athenahealth

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FlexBool is a boolean that tolerates athenahealth returning the same field as
// either a JSON bool (true/false) or a string ("", "true", "false"). athenahealth
// is inconsistent about this for fields like "documentationonly", so a plain bool
// or string fails to unmarshal half the responses. An empty or null value is
// treated as false.
type FlexBool bool

func (b *FlexBool) UnmarshalJSON(data []byte) error {
	var aux interface{}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch v := aux.(type) {
	case bool:
		*b = FlexBool(v)
	case string:
		if v == "" {
			*b = false
			return nil
		}
		parsed, err := strconv.ParseBool(v)
		if err != nil {
			return fmt.Errorf("cannot parse %q as bool: %w", v, err)
		}
		*b = FlexBool(parsed)
	case nil:
		*b = false
	default:
		return fmt.Errorf("unknown type: %T", v)
	}

	return nil
}

type NumberString string

func (n *NumberString) UnmarshalJSON(data []byte) error {
	var aux interface{}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch v := aux.(type) {
	case string:
		*n = NumberString(v)
	case int:
		*n = NumberString(strconv.Itoa(v))
	case float64:
		*n = NumberString(strconv.FormatFloat(v, 'f', -1, 64))
	default:
		return fmt.Errorf("unknown type: %T", v)
	}

	return nil
}

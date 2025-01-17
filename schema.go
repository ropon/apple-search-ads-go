package asa

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

const (
	dateFormat          = "2006-01-02 15"
	ReqDateFormat       = "2006-01-02"
	customISO8601Format = "2006-01-02T15:04:05.000"
	emailRegexString    = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
)

var (
	emailRegex = regexp.MustCompile(emailRegexString)
)

// ErrInvalidEmail occurs when the value does not conform to the library author's understanding of
// what constitutes a valid email address, and cannot be marshaled or unmarshaled into JSON.
type ErrInvalidEmail struct {
	Value string
}

func (e ErrInvalidEmail) Error() string {
	return fmt.Sprintf("email: %s failed to pass regex validation", e.Value)
}

// Date represents a date with no time component.
type Date struct {
	time.Time
}

// MarshalJSON is a custom marshaller for time-less dates.
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(dateFormat))
}

// UnmarshalJSON is a custom unmarshaller for time-less dates.
func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string

	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}

	parsed, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return err
	}

	d.Time = parsed

	return nil
}

// ReqDate represents a date with no time component.
type ReqDate struct {
	time.Time
}

// MarshalJSON is a custom marshaller for time-less dates.
func (d ReqDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(ReqDateFormat))
}

// UnmarshalJSON is a custom unmarshaller for time-less dates.
func (d *ReqDate) UnmarshalJSON(data []byte) error {
	var dateStr string

	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}

	parsed, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return err
	}

	d.Time = parsed

	return nil
}

// DateTime represents a date with an ISO8601-like date-time.
type DateTime struct {
	time.Time
}

// MarshalJSON is a custom marshaller for date-times.
func (d DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(d.Time.Format(customISO8601Format))
}

// UnmarshalJSON is a custom unmarshaller for date-times.
func (d *DateTime) UnmarshalJSON(data []byte) error {
	var dateTimeStr string

	err := json.Unmarshal(data, &dateTimeStr)
	if err != nil {
		return err
	}

	parsed, err := time.Parse(time.RFC3339, dateTimeStr)
	if err != nil {
		parsed, err = time.Parse(customISO8601Format, dateTimeStr)
		if err != nil {
			return err
		}
	}

	d.Time = parsed

	return nil
}

// Email is a validated email address string.
type Email string

// MarshalJSON is a custom marshaler for email addresses.
func (e Email) MarshalJSON() ([]byte, error) {
	emailString := string(e)
	if !emailRegex.MatchString(emailString) {
		return nil, ErrInvalidEmail{Value: emailString}
	}

	return json.Marshal(string(e))
}

// UnmarshalJSON is a custom unmarshaller for email addresses.
func (e *Email) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if !emailRegex.MatchString(s) {
		return ErrInvalidEmail{Value: s}
	}

	*e = Email(s)

	return nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	return &v
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int {
	return &v
}

// Float is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float(v float64) *float64 {
	return &v
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	return &v
}

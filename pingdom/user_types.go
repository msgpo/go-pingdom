package pingdom

import (
	"fmt"
)

// UserSms represents the sms contact object for a User.
type UserSms struct {
	Severity    string `json:"severity"`
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
	Provider    string `json:"provider"`
}

// UserEmail represents the email contact object for a User.
type UserEmail struct {
	Severity string `json:"severity"`
	Address  string `json:"address"`
}

// Contact represents a Pingdom contact target.
type Contact struct {
	Severity    string `json:"severitylevel"`
	CountryCode string `json:"countrycode"`
	Number      string `json:"number"`
	Provider    string `json:"provider"`
	Email       string `json:"email"`
}

// User represents a Pingdom User or Contact.
type User struct {
	Paused   string              `json:"paused,omitempty"`
	Username string              `json:"name,omitempty"`
	Primary  string              `json:"primary,omitempty"`
	Sms      []UserSmsResponse   `json:"sms,omitempty"`
	Email    []UserEmailResponse `json:"email,omitempty"`
}

// ValidUser determines whether a User contains valid fields.
func (u *User) ValidUser() error {
	if u.Username == "" {
		return fmt.Errorf("Invalid value for `Username`.  Must contain non-empty string")
	}

	return nil
}

// ValidContact determines whether a Contact contains valid fields.
func (c *Contact) ValidContact() error {
	if c.Email == "" && c.Number == "" {
		return fmt.Errorf("you must provide either an Email or a Phone Number to create a contact target")
	}

	if c.Number != "" && c.CountryCode == "" {
		return fmt.Errorf("you must provide a Country Code if providing a phone number")
	}

	if c.Provider != "" && (c.Number == "" || c.CountryCode == "") {
		return fmt.Errorf("you must provide CountryCode and Number if Provider is provided")
	}

	return nil
}

// PostParams returns a map of params that are sent with an HTTP POST request for a User.
func (u *User) PostParams() map[string]string {
	m := map[string]string{
		"name": u.Username,
	}

	return m
}

// PostContactParams returns a map of params that are sent with an HTTP POST request for a Contact.
func (c *Contact) PostContactParams() map[string]string {
	m := map[string]string{}

	// Ignore if not defined
	if c.Email != "" {
		m["email"] = c.Email
	}

	if c.Number != "" {
		m["number"] = c.Number
	}

	if c.CountryCode != "" {
		m["countrycode"] = c.CountryCode
	}

	if c.Severity != "" {
		m["severitylevel"] = c.Severity
	}

	if c.Provider != "" {
		m["provider"] = c.Provider
	}

	return m
}

// PutParams returns a map of params that are sent with an HTTP PUT request for a User.
func (u *User) PutParams() map[string]string {
	m := map[string]string{
		"name": u.Username,
	}

	if u.Primary != "" {
		m["primary"] = u.Primary
	}

	if u.Paused != "" {
		m["paused"] = u.Paused
	}

	return m
}

// PutContactParams returns a map of params that are sent with an HTTP PUT request for a Contact.
func (c *Contact) PutContactParams() map[string]string {
	return c.PostContactParams()
}

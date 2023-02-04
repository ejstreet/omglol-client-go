package omglol

import (
	"os"
	"testing"
)

var testEmail = os.Getenv("OMGLOL_USER_EMAIL")
var testKey   = os.Getenv("OMGLOL_API_KEY")
var testName  = os.Getenv("OMGLOL_USERNAME")


func TestGetAccountInfo(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountInfo()

	if a.Message != "Here is the account info that you requested." {
		t.Errorf(err.Error())
	}
}

func TestGetAddresses(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountAddresses()

	t.Logf("%+v\n", *a)
}

func TestSetAccountName(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.SetAccountName(testName)

	t.Logf("%+v\n", a)

	if a.Name != testName {
		t.Errorf(err.Error())
	}
}

func TestGetAccountName(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountName()

	t.Logf("%+v\n", a)

	if a.Name != testName {
		t.Errorf(err.Error())
	}
}

func TestGetActiveSessions(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetActiveSessions()

	t.Logf("%+v\n", a)

	if !a.Request.Success {
		t.Errorf(err.Error())
	}
}

// This test cannot currently be run automatically
// func TestDeleteActiveSession(t *testing.T) {
// 	c, err := NewClient(testEmail, testKey)

// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// 	a, err := c.DeleteActiveSession("1be17d138f202b0fba996192f22cc249")

// 	t.Logf("%+v\n", a)

// 	if !a.Request.Success {
// 		t.Errorf(err.Error())
// 	}
// }

func TestGetAccountSettings(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountSettings()

	if a.Settings.Owner == "" {
		t.Errorf(err.Error())
	}
}

func TestSetAccountSettings(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.SetAccountSettings(map[string]string{"communication": "email_ok"})

	t.Logf("%+v\n", a)

	if !a.Request.Success {
		t.Errorf(err.Error())
	}
}
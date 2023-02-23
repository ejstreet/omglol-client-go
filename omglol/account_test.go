package omglol

import (
	"os"
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	t.Parallel()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", *a)
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountAddresses()
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", *a)
}

func TestSetAccountName(t *testing.T) {
	t.Parallel()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	if err := c.SetAccountName(testName); err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetAccountName(t *testing.T) {
	t.Parallel()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	name, err := c.GetAccountName()
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", *name)

	if *name != testName {
		t.Errorf("Expected %s, got %s", testName, *name)
	}
}

func TestGetActiveSessions(t *testing.T) {
	t.Parallel()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetActiveSessions()
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", a)
}

// This test cannot currently be run automatically
func TestDeleteActiveSession(t *testing.T) {
	t.Parallel()
	sessionID := os.Getenv("OMGLOL_DELETABLE_SESSION_ID")

	if sessionID == "" {
		t.Skip()
	}

	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	if err := c.DeleteActiveSession(sessionID); err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetAccountSettings(t *testing.T) {
	t.Parallel()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	s, err := c.GetAccountSettings()
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", s)
}

func TestSetAccountSettings(t *testing.T) {
	t.Parallel()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	err = c.SetAccountSettings(map[string]string{"communication": "email_ok"})
	if err != nil {
		t.Errorf(err.Error())
	}
}

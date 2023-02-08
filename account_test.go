package client

import (
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
	}

	if !a.Request.Success {
		t.Errorf(err.Error())
	}
}

func TestGetAddresses(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountAddresses()
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", *a)

	if !a.Request.Success {
		t.Errorf(err.Error())
	}
}

func TestSetAccountName(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.SetAccountName(testName)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", a)

	if !a.Request.Success {
		t.Errorf(err.Error())
	}
}

func TestGetAccountName(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountName()
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", a)

	if !a.Request.Success {
		t.Errorf(err.Error())
	}
}

func TestGetActiveSessions(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetActiveSessions()
	if err != nil {
		t.Errorf(err.Error())
	}

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
//  if err != nil {
// 	  t.Errorf(err.Error())
//  }

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
	if err != nil {
		t.Errorf(err.Error())
	}

	if !a.Request.Success {
		t.Errorf(err.Error())
	}
}

func TestSetAccountSettings(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.SetAccountSettings(map[string]string{"communication": "email_ok"})
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", a)

	if !a.Request.Success {
		t.Errorf(err.Error())
	}
}

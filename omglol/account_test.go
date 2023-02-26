package omglol

import (
	"net"
	"os"
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
	}

	if a != nil {
		t.Logf("%+v\n", *a)

		if a.APIKey != c.Auth.ApiKey {
			t.Error("Incorrect API key")
		}
		if a.Email != c.Auth.Email {
			t.Errorf("Incorrect email: %s", a.Email)
		}
		if a.Settings.Communication != nil && !isOneOf(*a.Settings.Communication, []string{"email_ok", "email_not_ok"}) {
			t.Errorf("Invalid communication setting: %s", *a.Settings.Communication)
		}
		if a.Settings.DateFormat != nil && !isOneOf(*a.Settings.DateFormat, []string{"iso_8601", "dmy", "mdy"}) {
			t.Errorf("Invalid date format setting: %s", *a.Settings.DateFormat)
		}
		if a.Settings.Owner != testEmail {
			t.Errorf("Settings.Owner %s did not match expected %s.", a.Settings.Owner, testEmail)
		}
		testTimestamps(t, a.Created.UnixEpochTime, a.Created.Iso8601Time, a.Created.Rfc2822Time, a.Created.RelativeTime)
	} else {
		t.Error("Account Info returned 'nil'.")
	}

}

func TestGetAddresses(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAccountAddresses()
	if err != nil {
		t.Errorf(err.Error())
	}

	if a != nil {
		t.Logf("%+v\n", *a)

		for _, addr := range *a {
			if len(addr.Address) <= 0 {
				t.Errorf("Invalid address")
			}
			testTimestamps(t, addr.Expiration.UnixEpochTime, addr.Expiration.Iso8601Time, addr.Expiration.Rfc2822Time, addr.Expiration.RelativeTime)
			testTimestamps(t, addr.Registration.UnixEpochTime, addr.Registration.Iso8601Time, addr.Registration.Rfc2822Time, addr.Registration.RelativeTime)
		}
	} else {
		t.Error("Addresses returned 'nil'.")
	}

}

func TestSetAccountName(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	if err := c.SetAccountName(testName); err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetAccountName(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	name, err := c.GetAccountName()
	if err != nil {
		t.Errorf(err.Error())
	}

	if name != nil {
		t.Logf("%+v\n", *name)

		if *name != testName {
			t.Errorf("Expected %s, got %s", testName, *name)
		}
	} else {
		t.Error("Account name returned 'nil'.")
	}

}

func TestGetActiveSessions(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	s, err := c.GetActiveSessions()
	if err != nil {
		t.Errorf(err.Error())
	}

	if s != nil {
		t.Logf("Active Sessions: %+v", *s)
		for _, x := range *s {
			if len(x.SessionID) != 32 {
				t.Errorf("Session ID %s is not expected length of 32, got length %d", x.SessionID, len(x.SessionID))
			}
			if len(x.UserAgent) <= 0 {
				t.Errorf("user_agent is empty")
			}
			if net.ParseIP(x.CreatedIP) == nil {
				t.Errorf("Invalid IP address: %s", x.CreatedIP)
			}
			if x.CreatedOn == 0 {
				t.Errorf("Created on timestamp is 0.")
			}
			if x.CreatedOn > x.ExpiresOn {
				t.Errorf("Create date: %d, is after expire date: %d", x.CreatedOn, x.ExpiresOn)
			}
		}
	} else {
		t.Error("Active sessions returned 'nil'.")
	}

}

// This test cannot currently be run automatically
func TestDeleteActiveSession(t *testing.T) {	
	sessionID, exists := os.LookupEnv("OMGLOL_DELETABLE_SESSION_ID")
	if !exists {
		t.Skip()
	}
	sleep()

	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	if err := c.DeleteActiveSession(sessionID); err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetAccountSettings(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	s, err := c.GetAccountSettings()
	if err != nil {
		t.Errorf(err.Error())
	}

	if s != nil {
		t.Logf("%+v\n", *s)

		if s.Communication != nil && !isOneOf(*s.Communication, []string{"email_ok", "email_not_ok"}) {
			t.Errorf("Invalid communication setting: %s", *s.Communication)
		}
		if s.DateFormat != nil && !isOneOf(*s.DateFormat, []string{"iso_8601", "dmy", "mdy"}) {
			t.Errorf("Invalid date format setting: %s", *s.DateFormat)
		}
		if s.Owner != testEmail {
			t.Errorf("Settings.Owner %s did not match expected %s.", s.Owner, testEmail)
		}
		// s.WebEditor appears to be depricated
	} else {
		t.Errorf("Account Settings returned 'nil'.")
	}

}

func TestSetAccountSettings(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	err = c.SetAccountSettings(map[string]string{"communication": "email_ok"})
	if err != nil {
		t.Errorf(err.Error())
	}
}

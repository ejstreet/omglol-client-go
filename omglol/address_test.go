package omglol

import (
	"strings"
	"testing"
)

func TestGetAddressAvailability(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	// Test unavailable address
	u, err := c.GetAddressAvailability(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if u != nil {
		t.Logf("%+v\n", *u)

		if u.Address != testOwnedDomain {
			t.Errorf("Returned address: %s did not match expected address %s.", u.Address, testOwnedDomain)
		}
		if strings.ToLower(u.Availability) != "unavailable" {
			t.Errorf("Returned availability: %s did not match expected 'unavailable'.", u.Availability)
		}
		if u.Available != false {
			t.Errorf("Return available field was %t, expected false.", u.Available)
		}
	} else {
		t.Error("Address Availability returned 'nil' when getting unavailable address.")
	}

	// Test available address by creating a random string 32 characters long.
	// If anyone *has* an address 32 characters long, there is a 1 in 6.33 x 10^49 chance that it will match and this test will fail
	availableDomain := randStringBytes(32)
	a, err := c.GetAddressAvailability(availableDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if a != nil {
		t.Logf("%+v\n", *a)

		if a.Address != availableDomain {
			t.Errorf("Returned address: %s did not match expected address %s.", a.Address, availableDomain)
		}
		if strings.ToLower(a.Availability) != "available" {
			t.Errorf("Returned availability: %s did not match expected 'available'.", u.Availability)
		}
		if a.Available != true {
			t.Errorf("Return available field was %t, expected true.", u.Available)
		}
	} else {
		t.Error("Address Availability returned 'nil' when getting available address.")
	}

}

func TestGetAddressExpiration(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	e, err := c.GetAddressExpiration(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if e != nil {
		t.Logf("Expiration for domain: '%s' returned: '%t'.", testOwnedDomain, *e)
		if *e != false {
			t.Errorf("Expected expiration to be false, got %t.", *e)
		}
	} else {
		t.Errorf("Expiration returned 'nil'.")
	}
}

func TestGetAddressInfo(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	i, err := c.GetAddressInfo(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if i != nil {
		t.Logf("Address Info: %+v\n", *i)

		if i.Address != testOwnedDomain {
			t.Errorf("Expected Address: %s, got %s", testOwnedDomain, i.Address)
		}
		if len(i.Message) <= 0 {
			t.Error("Expected Message to not be empty.")
		}
		testTimestamps(t, i.Registration.UnixEpochTime, i.Registration.Iso8601Time, i.Registration.Rfc2822Time, i.Registration.RelativeTime)
		testTimestamps(t, i.Expiration.UnixEpochTime, i.Expiration.Iso8601Time, i.Expiration.Rfc2822Time, i.Expiration.RelativeTime)
		if i.Verification.Verified {
			if i.Verification.Message != "This address has been verified." {
				t.Errorf("Unexpected message when Verified == `true`: %s", i.Verification.Message)
			}
		} else {
			if i.Verification.Message != "This address has not been verified." {
				t.Errorf("Unexpected message when Verified == `false`: %s", i.Verification.Message)
			}
		}
		if len(i.Owner) <= 0 {
			t.Error("Unexpected empty Owner.")
		}
	} else {
		t.Error("Address info returned 'nil'.")
	}
}

func TestGetAddressDirectory(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	d, err := c.GetAddressDirectory()
	if err != nil {
		t.Errorf(err.Error())
	}

	if d != nil {
		t.Logf("Address Directory: %+v\n", *d)
		if len(d.Directory) <= 0 {
			t.Error("Directory returned empty.")
		}
	} else {
		t.Error("Address directory returned 'nil'.")
	}
}

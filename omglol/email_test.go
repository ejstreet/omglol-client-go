package omglol

import (
	"testing"
)

func TestGetClearAndSetEmails(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	originalDestinations, err := c.GetEmails(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if originalDestinations != nil {
		t.Logf("Original Destinations: %+v\n", originalDestinations)

	} else {
		t.Error("Account destination returned 'nil'.")
	}

	sleep()

	clear := []string{}
	if err := c.SetEmails(testOwnedDomain, clear); err != nil {
		t.Errorf(err.Error())
	}

	sleep()

	if err != nil {
		t.Errorf(err.Error())
	}

	destination, err := c.GetEmails(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if destination != nil {
		t.Logf("Destinations once cleared: %+v\n", destination)

		if len(destination) != 0 {
			t.Errorf("Expected %d, got %d", 0, len(destination))
		}
	} else {
		t.Error("Account destination returned 'nil'.")
	}

	if err := c.SetEmails(testOwnedDomain, originalDestinations); err != nil {
		t.Errorf(err.Error())
	}

	finalDestinations, err := c.GetEmails(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if finalDestinations != nil {
		t.Logf("Final Destinations: %+v\n", finalDestinations)

	} else {
		t.Error("Account destination returned 'nil'.")
	}

	if !listsHaveSameElements(t, originalDestinations, finalDestinations) {
		t.Errorf("Original: %s does not match final: %s.", originalDestinations, finalDestinations)
	}
}

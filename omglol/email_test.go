package omglol

import (
	"testing"
)

func TestSetAndGetEmail(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	destinations := []string{testEmail}
	if err := c.SetEmail(testOwnedDomain, destinations); err != nil {
		t.Errorf(err.Error())
	}

	sleep()

	if err != nil {
		t.Errorf(err.Error())
	}

	destination, err := c.GetEmail(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if destination != nil {
		t.Logf("%+v\n", destination)

		expected := []string{testEmail}
		if destination[0] != testEmail {
			t.Errorf("Expected %s, got %s", expected, destination)
		}
	} else {
		t.Error("Account destination returned 'nil'.")
	}
}

func TestClearAndGetEmail(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	destinations := []string{}
	if err := c.SetEmail(testOwnedDomain, destinations); err != nil {
		t.Errorf(err.Error())
	}

	sleep()

	if err != nil {
		t.Errorf(err.Error())
	}

	destination, err := c.GetEmail(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if destination != nil {
		t.Logf("%+v\n", destination)

		if len(destination) != 0 {
			t.Errorf("Expected %d, got %d", 0, len(destination))
		}
	} else {
		t.Error("Account destination returned 'nil'.")
	}
}

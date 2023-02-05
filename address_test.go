package omglol

import (
	"testing"
)

func TestGetAddressAvailability(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAddressAvailability("test")

	t.Logf("%+v\n", *a)

	if a.Message == "" {
		t.Errorf(err.Error())
	}
}

func TestGetAddressExpiration(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAddressExpiration("test")

	t.Logf("%+v\n", *a)

	if a.Message == "" {
		t.Errorf(err.Error())
	}
}

func TestGetAddressInfo(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	i, err := c.GetAddressInfo("test")

	t.Logf("%+v\n", *i)

	if i.Message == "" {
		t.Errorf(err.Error())
	}
}
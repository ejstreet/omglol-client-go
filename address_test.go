package client

import (
	"testing"
)

func TestGetAddressAvailability(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAddressAvailability("test")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", *a)
}

func TestGetAddressExpiration(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetAddressExpiration("test")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", *a)
}

func TestGetAddressInfo(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	i, err := c.GetAddressInfo("terraform")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", *i)
}

func TestGetAddressDirectory(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	d, err := c.GetAddressDirectory()
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", *d)
}

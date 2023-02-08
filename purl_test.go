package client

import (
	"testing"
)

func TestGetPersistentURL(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	a, err := c.GetPersistentURL(testOwnedDomain, "testget")
	if err != nil {
		t.Errorf(err.Error())
	}

	if !a.Request.Success {
		t.Errorf(err.Error())
	}
}

func TestListPersistentURLs(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	d, err := c.ListPersistentURLs(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", *d)

	if d.Response.Message == "" {
		t.Errorf(err.Error())
	}
}

func TestCreateAndDeletePersistentURL(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	purlName := "test"

	r, err := c.CreatePersistentURL(testOwnedDomain, purlName, "https://example.com")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", r)

	if !r.Request.Success {
		t.Errorf(err.Error())
	}

	m, err := c.DeletePersistentURL(testOwnedDomain, purlName)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", m)

	if !m.Request.Success {
		t.Errorf(err.Error())
	}
}

package omglol

import (
	"testing"
)

func TestGetPersistentURL(t *testing.T) {
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	p, err := c.GetPersistentURL(testOwnedDomain, "testget")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(p.ToString())
}

func TestListPersistentURLs(t *testing.T) {
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	l, err := c.ListPersistentURLs(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, p := range *l {
		t.Logf(p.ToString() + "\n")
	}
}

func TestCreateAndDeletePersistentURL(t *testing.T) {
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	name := "test" + RunUID

	p := NewPersistentURL(name, "https://example.com", false)

	err = c.CreatePersistentURL(testOwnedDomain, *p)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = c.DeletePersistentURL(testOwnedDomain, name)
	if err != nil {
		t.Errorf(err.Error())
	}
}

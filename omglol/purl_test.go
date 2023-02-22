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

	name1 := "testunlisted" + RunUID

	unlisted := NewPersistentURL(name1, "https://example.com", false)

	err = c.CreatePersistentURL(testOwnedDomain, *unlisted)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = c.GetPersistentURL(testOwnedDomain, name1)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = c.DeletePersistentURL(testOwnedDomain, name1)
	if err != nil {
		t.Errorf(err.Error())
	}

	name2 := "testlisted" + RunUID

	listed := NewPersistentURL(name2, "https://example.com", true)

	err = c.CreatePersistentURL(testOwnedDomain, *listed)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = c.GetPersistentURL(testOwnedDomain, name2)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = c.DeletePersistentURL(testOwnedDomain, name2)
	if err != nil {
		t.Errorf(err.Error())
	}
}

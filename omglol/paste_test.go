package omglol

import (
	"testing"
)

func validatePaste(t *testing.T, p Paste) {
	if len(p.Title) <= 0 {
		t.Error("Paste Title is empty.")
	}
	if *p.ModifiedOn <= 0 {
		t.Error("ModifiedOn time is unset.")
	}
}

func TestGetPaste(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	p, err := c.GetPaste(testOwnedDomain, "testget")
	if err != nil {
		t.Logf("This test will fail if a paste named 'testget' does not exist.")
		t.Errorf(err.Error())
	}

	if p != nil {
		t.Logf(p.String())
		validatePaste(t, *p)
	} else {
		t.Error("GetPaste returned 'nil'.")
	}
}

func TestListPastes(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	l, err := c.ListPastes(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if l != nil {
		for _, p := range *l {
			t.Logf(p.String() + "\n")
			validatePaste(t, p)
		}
	} else {
		t.Error("ListPastes returned 'nil'.")
	}
}

func TestCreateAndDeletePaste(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	unlistedTitle := "unlistedtest" + RunUID

	unlistedPaste := NewPaste(unlistedTitle, "example paste content", false)

	err = c.CreatePaste(testOwnedDomain, *unlistedPaste)
	if err != nil {
		t.Errorf(err.Error())
	}

	u, err := c.GetPaste(testOwnedDomain, unlistedTitle)
	if err != nil {
		t.Errorf(err.Error())
	}

	if u != nil {
		t.Log(u.String())
		validatePaste(t, *u)
		if u.Listed != false {
			t.Error("Unlisted paste should have Listed value 'false'.")
		}
	} else {
		t.Error("GetPaste returned 'nil' when retrieving unlisted paste.")
	}

	err = c.DeletePaste(testOwnedDomain, unlistedTitle)
	if err != nil {
		t.Errorf(err.Error())
	}
	sleep()
	listedTitle := "listedtest" + RunUID
	listedPaste := NewPaste(listedTitle, "example paste content", true)

	err = c.CreatePaste(testOwnedDomain, *listedPaste)
	if err != nil {
		t.Errorf(err.Error())
	}

	l, err := c.GetPaste(testOwnedDomain, listedTitle)
	if err != nil {
		t.Errorf(err.Error())
	}

	if l != nil {
		t.Log(l.String())
		validatePaste(t, *l)
		if l.Listed != true {
			t.Error("Listed paste should have Listed value 'true'.")
		}
	} else {
		t.Error("GetPaste returned 'nil' when retrieving listed paste.")
	}

	err = c.DeletePaste(testOwnedDomain, listedTitle)
	if err != nil {
		t.Errorf(err.Error())
	}
}

package omglol

import (
	"testing"
)

func validatePaste(t *testing.T, p Paste) {
	if len(p.Title) <= 0 {
		t.Error("Paste Title is empty.")
	}
	if len(p.Content) <= 0 {
		t.Error("Paste Content is empty.")
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

	title := "test" + RunUID

	paste := NewPaste(title, "example paste content", 0)

	err = c.CreatePaste(testOwnedDomain, *paste)
	if err != nil {
		t.Errorf(err.Error())
	}

	u, err := c.GetPaste(testOwnedDomain, title)
	if err != nil {
		t.Errorf(err.Error())
	}

	if u != nil {
		t.Log(u.String())
		validatePaste(t, *u)
	} else {
		t.Error("GetPaste returned 'nil' when retrieving paste.")
	}

	err = c.DeletePaste(testOwnedDomain, title)
	if err != nil {
		t.Errorf(err.Error())
	}
}

package omglol

import (
	"testing"
)

func validatePersistentURL(t *testing.T, p PersistentURL) {
	if len(p.Name) <= 0 {
		t.Error("PURL Name is empty.")
	}
	if len(p.URL) <= 0 {
		t.Error("PURL URL is empty.")
	}
	if p.Counter == nil {
		t.Error("PURL Counter should not be 'nil'.")
	}
}

func TestGetPersistentURL(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	p, err := c.GetPersistentURL(testOwnedDomain, "testget")
	if err != nil {
		t.Logf("This test will fail if a purl named 'testget' does not exist.")
		t.Errorf(err.Error())
	}

	if p != nil {
		t.Logf(p.String())
		validatePersistentURL(t, *p)
	} else {
		t.Error("GetPersistentURL returned 'nil'.")
	}
}

func TestListPersistentURLs(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	l, err := c.ListPersistentURLs(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if l != nil {
		for _, p := range *l {
			t.Logf(p.String() + "\n")
			validatePersistentURL(t, p)
		}
	} else {
		t.Error("ListPersistentURLs returned 'nil'.")
	}

}

func TestCreateGetAndDeletePersistentURL(t *testing.T) {
	sleep()
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

	u, err := c.GetPersistentURL(testOwnedDomain, name1)
	if err != nil {
		t.Errorf(err.Error())
	}

	if u != nil {
		t.Log(u.String())
		validatePersistentURL(t, *u)
		if u.Listed != false {
			t.Error("Unlisted PURL should not have listed set to 'true'.")
		}
	} else {
		t.Error("GetPersistentURL returned 'nil' when retrieving unlisted PURL.")
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

	l, err := c.GetPersistentURL(testOwnedDomain, name2)
	if err != nil {
		t.Errorf(err.Error())
	}

	if l != nil {
		t.Log(l.String())
		validatePersistentURL(t, *l)
		if l.Listed != true {
			t.Error("Listed PURL should not have listed set to 'false'.")
		}
	} else {
		t.Error("GetPersistentURL returned 'nil' when retrieving listed PURL.")
	}

	err = c.DeletePersistentURL(testOwnedDomain, name2)
	if err != nil {
		t.Errorf(err.Error())
	}
}

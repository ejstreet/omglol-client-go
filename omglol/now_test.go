package omglol

import (
	"testing"
)

func TestGetSetRestoreNow(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	w, err := c.GetNow(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if w != nil {
		t.Logf("Previous content: %s", w.Content)
	} else {
		t.Logf("Previous content was empty.")
	}

	previousContent := w.Content
	newContent := "this is a now page\n\nright now I am running tests\n"

	sleep()
	err = c.SetNow(testOwnedDomain, []byte(newContent), false)
	if err != nil {
		t.Errorf(err.Error())
	}

	sleep()
	w, err = c.GetNow(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if w != nil {
		t.Logf("New content: %s", w.Content)
		if w.Content != newContent {
			t.Errorf("Expected content '%s', got '%s'", newContent, w.Content)
		}
		if string(w.ContentBytes) != newContent {
			t.Errorf("Expected ContentBytes '%s', got '%s'", newContent, string(w.Content))
		}
	} else {
		t.Errorf("GetNow returned 'nil'.")
	}

	sleep()
	err = c.SetNow(testOwnedDomain, []byte(previousContent), true)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetNowGarden(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	g, err := c.GetNowGarden()
	if err != nil {
		t.Errorf(err.Error())
	}

	if g != nil {
		t.Logf("Now Garden: %+v\n", *g)
		if len(*g) <= 0 {
			t.Error("Garden returned empty.")
		}
	} else {
		t.Error("Now garden returned 'nil'.")
	}
}

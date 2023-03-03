package omglol

import (
	"testing"
)

func TestGetSetRestoreWeb(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	w, err := c.GetWeb(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	previousContent := w.Content
	newContent := "this is the new content\n\nit has newlines\n"

	sleep()
	published, err := c.SetWeb(testOwnedDomain, []byte(newContent), false)
	if err != nil {
		t.Errorf(err.Error())
	}
	if published {
		t.Errorf("Expected published to be false, got true")
	}

	sleep()
	w, err = c.GetWeb(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}
	if w.Content != newContent {
		t.Errorf("Expected content '%s', got '%s'", newContent, w.Content)
	}
	if string(w.ContentBytes) != newContent {
		t.Errorf("Expected ContentBytes '%s', got '%s'", newContent, string(w.Content))
	}

	sleep()
	published, err = c.SetWeb(testOwnedDomain, []byte(previousContent), true)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !published {
		t.Errorf("Expected published to be true, got false")
	}
}

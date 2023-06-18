package omglol

import (
	"testing"
)

func TestGetServiceInfo(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	s, err := c.GetServiceInfo()
	if err != nil {
		t.Errorf(err.Error())
	}

	if s != nil {
		t.Logf("%+v\n", *s)
		if s.Members < 1 {
			t.Errorf("Invalid member count: %d", s.Members)
		}
		if s.Addresses < 1 {
			t.Errorf("Invalid address count: %d", s.Addresses)
		}
		if s.Profiles < 1 {
			t.Errorf("Invalid profile count: %d", s.Profiles)
		}
	} else {
		t.Error("Service Info returned 'nil'.")
	}
}

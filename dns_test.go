package omglol

import (
	"testing"
)

func TestGetDNSRecords(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	d, err := c.GetDNSRecords(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", *d)

	if d.Message == "" {
		t.Errorf(err.Error())
	}
}

// There is currently no test for the Update method, as it does not work at time of writing, see https://github.com/neatnik/omg.lol/issues/584
func TestCreateAndDeleteDNSRecord(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	record := map[string]string{
		"type": "CNAME",
		"name": "test",
		"data": "example.com",
		"ttl":  "300",
	}

	r, err := c.CreateDNSRecord(testOwnedDomain, record)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", r)

	if r.ResponseReceived.Data.ID <= 0 {
		t.Errorf(err.Error())
	}

	m, err := c.DeleteDNSRecord(testOwnedDomain, r.ResponseReceived.Data.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", m)

	if !m.Request.Success {
		t.Errorf(err.Error())
	}
}

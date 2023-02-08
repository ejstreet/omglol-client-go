package client

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

	if d.Response.Message == "" {
		t.Errorf(err.Error())
	}
}

func TestFilterDNSRecords(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	criteria := map[string]interface{}{
		"ID":        "41923511",
		"Type":      "TXT",
		"Name":      "testdns.terraform",
		"Data":      "test=true",
		"TTL":       "300",
		"CreatedAt": "2023-02-05T21:17:51Z",
		"UpdatedAt": "2023-02-05T21:17:51Z",
	}

	d, err := c.FilterDNSRecord(testOwnedDomain, criteria)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(*d.ID, *d.Type, *d.Name, *d.Data, *d.CreatedAt)

	t.Logf("%+v\n", *d)
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

	if !r.Request.Success {
		t.Errorf(err.Error())
	}

	m, err := c.DeleteDNSRecord(testOwnedDomain, r.Response.ResponseReceived.Data.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", m)

	if !m.Request.Success {
		t.Errorf(err.Error())
	}
}

func TestCreateReplaceDeleteDNSRecord(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	record1 := map[string]string{
		"type": "TXT",
		"name": "testcreate",
		"data": "example.com",
		"ttl":  "300",
	}

	record2 := map[string]string{
		"type": "TXT",
		"name": "testreplace",
		"data": "example.com",
		"ttl":  "300",
	}

	create, err := c.CreateDNSRecord(testOwnedDomain, record1)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", create)

	if !create.Request.Success {
		t.Errorf(err.Error())
	}

	replace, err := c.ReplaceDNSRecord(testOwnedDomain, record2, create.Response.ResponseReceived.Data.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", replace)

	if !replace.Request.Success {
		t.Errorf(err.Error())
	}

	delete, err := c.DeleteDNSRecord(testOwnedDomain, replace.Response.ResponseReceived.Data.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", delete)

	if !delete.Request.Success {
		t.Errorf(err.Error())
	}
}

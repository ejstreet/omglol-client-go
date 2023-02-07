package client

import (
	"fmt"
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

func TestFilterDNSRecords(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	criteria := map[string]interface{}{
		"ID": "41923511",
		"Type": "TXT",
		"Name": "testdns.terraform",
		"Data": "test=true",
		"TTL": "300",
		"CreatedAt": "2023-02-05T21:17:51Z",
		"UpdatedAt": "2023-02-05T21:17:51Z",
	}

	d, err := c.FilterDNSRecord(testOwnedDomain, criteria)

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf(*d.ID, *d.Type, *d.Name, *d.Data, *d.CreatedAt)

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

func TestCreateReplaceDeleteDNSRecord(t *testing.T) {
	c, err := NewClient(testEmail, testKey)

	if err != nil {
		t.Errorf(err.Error())
	}

	record1 := map[string]string{
		"type": "CNAME",
		"name": "foo",
		"data": "example.com",
		"ttl":  "300",
	}

	record2 := map[string]string{
		"type": "CNAME",
		"name": "bar",
		"data": "example.com",
		"ttl":  "300",
	}

	create, err := c.CreateDNSRecord(testOwnedDomain, record1)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", create)

	if create.ResponseReceived.Data.ID <= 0 {
		t.Errorf(err.Error())
	}

	replace, err := c.ReplaceDNSRecord(testOwnedDomain, record2, create.ResponseReceived.Data.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", replace)

	if replace.ResponseReceived.Data.ID <= 0 {
		t.Errorf(err.Error())
	}

	delete, err := c.DeleteDNSRecord(testOwnedDomain, replace.ResponseReceived.Data.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("%+v\n", delete)

	if !delete.Request.Success {
		t.Errorf(err.Error())
	}
}

package omglol

import (
	"testing"
)

func TestListDNSRecords(t *testing.T) {
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	l, err := c.ListDNSRecords(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, d := range *l {
		t.Logf(d.ToString() + "\n")
	}
}

func TestFilterDNSRecords(t *testing.T) {
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	criteria := map[string]interface{}{
		"Name": "testdns.terraform",
		"Type": "TXT",
		"Data": "test=true",
		"TTL":  300,
	}

	d, err := c.FilterDNSRecord(testOwnedDomain, criteria)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(d.ToString())
}

// There is currently no test for the Update method, as it does not work at time of writing, see https://github.com/neatnik/omg.lol/issues/584
func TestCreateAndDeleteDNSRecord(t *testing.T) {
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	record := NewDNSEntry("TXT", "testcreatetodelete"+RunUID, "test=true", 300)

	r, err := c.CreateDNSRecord(testOwnedDomain, *record)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(r.ToString())

	err = c.DeleteDNSRecord(testOwnedDomain, *r.ID)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestCreateReplaceDeleteDNSRecord(t *testing.T) {
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	record1 := NewDNSEntry("TXT", "testcreate"+RunUID, "test=true", 300)

	record2 := NewDNSEntry("TXT", "testreplace"+RunUID, "test=true", 300)

	create, err := c.CreateDNSRecord(testOwnedDomain, *record1)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(create.ToString())

	replace, err := c.ReplaceDNSRecord(testOwnedDomain, *record2, *create.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(replace.ToString())

	err = c.DeleteDNSRecord(testOwnedDomain, *replace.ID)
	if err != nil {
		t.Errorf(err.Error())
	}
}

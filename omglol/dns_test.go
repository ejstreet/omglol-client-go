package omglol

import (
	"testing"
	"time"
)

// Local test function to check validity
func validateRecord(t *testing.T, r DNSRecord) {
	if r.ID <= 0 {
		t.Errorf("Record ID is invalid.")
	}
	if !isOneOf(r.Type, []string{"A", "AAAA", "CAA", "CNAME", "MX", "NS", "SRV", "TXT"}) {
		t.Errorf("Unexpected record type: %s.", r.Type)
	}
	if len(r.Name) <= 0 {
		t.Errorf("Name is empty.")
	}
	if len(r.Data) <= 0 {
		t.Errorf("Data is empty.")
	}
	if r.Type == "MX" && r.Priority == nil {
		t.Errorf("Priority cannot be nil for record type 'MX'.")
	}
	if r.Type != "MX" && r.Priority != nil {
		t.Errorf("Record of type '%s' should not have a priority not equal to 'nil'.", r.Type)
	}
	if r.TTL <= 0 {
		t.Errorf("TTL must be greater than 0, got %d", r.TTL)
	}
	created, err := time.Parse(time.RFC3339, r.CreatedAt)
	if err != nil {
		t.Error(err.Error())
	}
	if created.Unix() <= 0 {
		t.Error("Invalid CreatedAt timestamp.")
	}
	updated, err := time.Parse(time.RFC3339, r.UpdatedAt)
	if err != nil {
		t.Error(err.Error())
	}
	if updated.Unix() <= 0 {
		t.Error("Invalid UpdatedAt timestamp.")
	}
	if updated.Unix() < created.Unix() {
		t.Errorf("Updated: %s, cannot be before Created: %s.", r.UpdatedAt, r.CreatedAt)
	}
}

func TestListDNSRecords(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	l, err := c.ListDNSRecords(testOwnedDomain)
	if err != nil {
		t.Errorf(err.Error())
	}

	if l != nil {
		for _, d := range *l {
			t.Logf(d.String() + "\n")
			validateRecord(t, d)
		}
	} else {
		t.Error("ListDNSRecords returned 'nil'.")
	}

}

func TestFilterDNSRecords(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)
	if err != nil {
		t.Errorf(err.Error())
	}

	// TXT Record
	criteria1 := map[string]any{
		"Name":     "testlistdns." + testOwnedDomain,
		"Type":     "TXT",
		"TTL":      300,
		"Priority": nil,
	}

	d1, err := c.FilterDNSRecord(testOwnedDomain, criteria1)
	if err != nil {
		t.Errorf(err.Error())
	}

	if d1 != nil {
		t.Logf(d1.String())
		validateRecord(t, *d1)
	} else {
		t.Logf("This test will fail if a TXT record named 'testget' with TTL 300 does not exist.")
		t.Errorf("FilterDNSRecord returned nil.")
	}

	// MX Record
	criteria2 := map[string]any{
		"Type":     "MX",
		"Name":     "mail." + testOwnedDomain,
		"Priority": int64(20),
	}

	d2, err := c.FilterDNSRecord(testOwnedDomain, criteria2)
	if err != nil {
		t.Logf("This test will fail if a MX record named 'mail' with priority 20 does not exist.")
		t.Errorf(err.Error())
	}

	if d2 != nil {
		t.Logf(d2.String())
		validateRecord(t, *d2)
	} else {
		t.Errorf("FilterDNSRecord returned nil.")
	}
}

func TestCreateAndDeleteDNSRecord(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	record := NewDNSEntry("TXT", "testcreatetodelete"+RunUID, "test=true", 300)

	r, err := c.CreateDNSRecord(testOwnedDomain, *record)
	if err != nil {
		t.Errorf(err.Error())
	}

	if r != nil {
		t.Logf(r.String())
		validateRecord(t, *r)
	} else {
		t.Error("CreateDNSRecord returned 'nil'.")
	}

	err = c.DeleteDNSRecord(testOwnedDomain, r.ID)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestCreateUpdateDeleteDNSRecord(t *testing.T) {
	sleep()
	c, err := NewClient(testEmail, testKey, testHostURL)

	if err != nil {
		t.Errorf(err.Error())
	}

	record1 := NewDNSEntry("TXT", "testcreate"+RunUID, "test=true", 300)

	record2 := NewDNSEntry("TXT", "testupdate"+RunUID, "test=true", 300)

	create, err := c.CreateDNSRecord(testOwnedDomain, *record1)
	if err != nil {
		t.Errorf(err.Error())
	}

	if create != nil {
		t.Logf(create.String())
		validateRecord(t, *create)
	} else {
		t.Error("CreateDNSRecord returned 'nil'.")
	}
	sleep()
	update, err := c.UpdateDNSRecord(testOwnedDomain, *record2, create.ID)
	if err != nil {
		t.Errorf(err.Error())
	}

	if update != nil {
		t.Logf(update.String())
		validateRecord(t, *update)
	} else {
		t.Error("UpdateDNSRecord returned 'nil'.")
	}

	sleep()
	err = c.DeleteDNSRecord(testOwnedDomain, update.ID)
	if err != nil {
		t.Errorf(err.Error())
	}
}

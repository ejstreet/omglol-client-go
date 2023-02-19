package omglol

import (
	"encoding/json"
)

type request struct {
	StatusCode int64 `json:"status_code"`
	Success    bool  `json:"success"`
}

type apiResponse struct {
	Request  request         `json:"request"`
	Response json.RawMessage `json:"response"`
}

type Registration struct {
	Message       string `json:"message,omitempty"`
	UnixEpochTime string `json:"unix_epoch_time,omitempty"`
	Iso8601Time   string `json:"iso_8601_time,omitempty"`
	Rfc2822Time   string `json:"rfc_2822_time,omitempty"`
	RelativeTime  string `json:"relative_time,omitempty"`
}

type AccountSettings struct {
	Owner         string `json:"owner,omitempty"`
	Communication string `json:"communication,omitempty"`
	DateFormat    string `json:"date_format,omitempty"`
	WebEditor     string `json:"web_editor,omitempty"`
}

type Account struct {
	Message string `json:"message,omitempty"`
	Email   string `json:"email,omitempty"`
	Name    string `json:"name,omitempty"`
	Created struct {
		UnixEpochTime string `json:"unix_epoch_time,omitempty"`
		Iso8601Time   string `json:"iso_8601_time,omitempty"`
		Rfc2822Time   string `json:"rfc_2822_time,omitempty"`
		RelativeTime  string `json:"relative_time,omitempty"`
	} `json:"created,omitempty"`
	APIKey   string          `json:"api_key,omitempty"`
	Settings AccountSettings `json:"settings,omitempty"`
}

type ActiveSession struct {
	SessionID string `json:"session_id"`
	UserAgent string `json:"user_agent"`
	CreatedIP string `json:"created_ip"`
	CreatedOn string `json:"created_on"`
	ExpiresOn string `json:"expires_on"`
}

type Address struct {
	Address      string            `json:"address"`
	Message      string            `json:"message,omitempty"`
	Punycode     string            `json:"punycode,omitempty"`
	SeeAlso      string            `json:"see-also,omitempty"`
	Registration Registration      `json:"registration"`
	Expiration   AddressExpiration `json:"expiration"`
	Owner        string            `json:"owner,omitempty"`
}

type AddressAvailability struct {
	Message      string `json:"message"`
	Address      string `json:"address"`
	Available    bool   `json:"available"`
	Availability string `json:"availability"`
}

type AddressExpiration struct {
	Message       string `json:"message,omitempty"`
	Expired       bool   `json:"expired,omitempty"`
	WillExpire    bool   `json:"will_expire,omitempty"`
	UnixEpochTime string `json:"unix_epoch_time,omitempty"`
	Iso8601Time   string `json:"iso_8601_time,omitempty"`
	Rfc2822Time   string `json:"rfc_2822_time,omitempty"`
	RelativeTime  string `json:"relative_time,omitempty"`
}

type AddressInfo struct {
	Address      string            `json:"address"`
	Message      string            `json:"message"`
	Registration Registration      `json:"registration"`
	Expiration   AddressExpiration `json:"expiration"`
	Verification struct {
		Message  string `json:"message"`
		Verified bool   `json:"verified"`
	} `json:"verification"`
	Owner string `json:"owner"`
}

type AddressDirectory struct {
	Message   string   `json:"message"`
	URL       string   `json:"url"`
	Directory []string `json:"directory"`
}

// Used to create or modify a DNS record
type DNSEntry struct {
	Type     *string `json:"type"`
	Name     *string `json:"name"`
	Data     *string `json:"data"`
	Priority *int64  `json:"priority"`
	TTL      *int64  `json:"ttl"`
}

// Return type for DNS related methods
type DNSRecord struct {
	ID        *int64  `json:"id"`
	Type      *string `json:"type"`
	Name      *string `json:"name"`
	Data      *string `json:"data"`
	Priority  *int64  `json:"priority"`
	TTL       *int64  `json:"ttl"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type dnsRecordContent struct {
	ID        *int64  `json:"id"`
	Type      *string `json:"type"`
	Name      *string `json:"name"`
	Content   *string `json:"content"`
	Priority  *int64  `json:"priority"`
	TTL       *int64  `json:"ttl"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type dnsChangeResponse struct {
	Request  request `json:"request"`
	Response struct {
		Message  string `json:"message"`
		DataSent struct {
			Type     string `json:"type"`
			Priority *int64 `json:"priority"`
			TTL      *int64 `json:"ttl"`
			Name     string `json:"name"`
			Content  string `json:"content"`
		} `json:"data_sent"`
		ResponseReceived struct {
			Data dnsRecordContent `json:"data"`
		} `json:"response_received"`
	} `json:"response"`
}

type PersistentURL struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Counter *int64 `json:"counter"`
	Listed  *bool  `json:"listed"`
}

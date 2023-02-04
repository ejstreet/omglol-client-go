package omglol

import (
	"encoding/json"
	"time"	
)

type Response struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response json.RawMessage `json:"response"`
}

type MessageResponse struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct{
		Message   string    `json:"message,omitempty"`
	} `json:"response"`
}

type Timestamp struct {
	UnixEpochTime string `json:"unix_epoch_time,omitempty"`
	Iso8601Time   string `json:"iso_8601_time,omitempty"`
	Rfc2822Time   string `json:"rfc_2822_time,omitempty"`
	RelativeTime  string `json:"relative_time,omitempty"`
}

type Registration struct {
	Message   string    `json:"message,omitempty"`
	Timestamp Timestamp `json:"registration"`
}

type Expiration struct {
	Expired    bool `json:"expired"`
	WillExpire bool `json:"will_expire"`
}

type Account struct {
	Message  string           `json:"message,omitempty"`
	Email    string           `json:"email,omitempty"`
	Name     string           `json:"name,omitempty"`
	Created  Timestamp        `json:"created,omitempty"`
	APIKey   string           `json:"api_key,omitempty"`
	Settings struct { 
		Owner         string `json:"owner,omitempty"`
		Communication string `json:"communication,omitempty"`
		DateFormat    string `json:"date_format,omitempty"`
		WebEditor     string `json:"web_editor,omitempty"`
	} `json:"settings,omitempty"`
}

type AccountSettings struct {
	Message  string `json:"message,omitempty"`
	Settings struct { 
		Owner         string `json:"owner,omitempty"`
		Communication string `json:"communication,omitempty"`
		DateFormat    string `json:"date_format,omitempty"`
		WebEditor     string `json:"web_editor,omitempty"`
	} `json:"settings,omitempty"`
}

type AccountName struct {
	Message  string           `json:"message,omitempty"`
	Name     string           `json:"name,omitempty"`
}

type ActiveSessions struct {
	Request struct {
		StatusCode int    `json:"status_code"`
		Success    bool   `json:"success"`
	} `json:"request"`
	Response []struct {
		SessionID   string `json:"session_id"`
		UserAgent   string `json:"user_agent"`
		CreatedIP   string `json:"created_ip"`
		CreatedOn   string `json:"created_on"`
		ExpiresOn   string `json:"expires_on"`
	} `json:"response"`
}

type Address struct {
	Address      string       `json:"address"`
	Message      string       `json:"message,omitempty"`
	Registration Registration `json:"registration"`
	Expiration   Expiration   `json:"expiration"`
	Owner        string       `json:"owner,omitempty"`
}

type Addresses struct {
	Addresses []Address `json:"response"`
}

type DNSRecord struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Name      string      `json:"name"`
	Data      string      `json:"data"`
	Priority  interface{} `json:"priority"`
	TTL       string      `json:"ttl"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type DNSRecords struct {
	Message string      `json:"message,omitempty"`
	DNS     []DNSRecord `json:"dns"`
}

type PersistantURL struct {
	Message string `json:"message"`
	Purl    struct {
		Name    string      `json:"name"`
		Url     string      `json:"url"`
		Counter interface{} `json:"counter"`
	} `json:"purl"`
}

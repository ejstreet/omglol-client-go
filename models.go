package client

// import (
// 	"encoding/json"
// )

type Request struct {
	StatusCode int  `json:"status_code"`
	Success    bool `json:"success"`
}

type MessageResponse struct {
	Request  Request `json:"request"`
	Response struct {
		Message string `json:"message,omitempty"`
	} `json:"response"`
}

type Registration struct {
	Message       string `json:"message,omitempty"`
	UnixEpochTime string `json:"unix_epoch_time,omitempty"`
	Iso8601Time   string `json:"iso_8601_time,omitempty"`
	Rfc2822Time   string `json:"rfc_2822_time,omitempty"`
	RelativeTime  string `json:"relative_time,omitempty"`
}

type Account struct {
	Request  Request `json:"request"`
	Response struct {
		Message string `json:"message,omitempty"`
		Email   string `json:"email,omitempty"`
		Name    string `json:"name,omitempty"`
		Created struct {
			UnixEpochTime string `json:"unix_epoch_time,omitempty"`
			Iso8601Time   string `json:"iso_8601_time,omitempty"`
			Rfc2822Time   string `json:"rfc_2822_time,omitempty"`
			RelativeTime  string `json:"relative_time,omitempty"`
		} `json:"created,omitempty"`
		APIKey   string `json:"api_key,omitempty"`
		Settings struct {
			Owner         string `json:"owner,omitempty"`
			Communication string `json:"communication,omitempty"`
			DateFormat    string `json:"date_format,omitempty"`
			WebEditor     string `json:"web_editor,omitempty"`
		} `json:"settings,omitempty"`
	} `json:"response"`
}

type AccountSettings struct {
	Request  Request `json:"request"`
	Response struct {
		Message  string `json:"message,omitempty"`
		Settings struct {
			Owner         string `json:"owner,omitempty"`
			Communication string `json:"communication,omitempty"`
			DateFormat    string `json:"date_format,omitempty"`
			WebEditor     string `json:"web_editor,omitempty"`
		} `json:"settings,omitempty"`
	} `json:"response"`
}

type AccountName struct {
	Request  Request `json:"request"`
	Response struct {
		Message string `json:"message,omitempty"`
		Name    string `json:"name,omitempty"`
	} `json:"response"`
}

type ActiveSessions struct {
	Request  Request `json:"request"`
	Response []struct {
		SessionID string `json:"session_id"`
		UserAgent string `json:"user_agent"`
		CreatedIP string `json:"created_ip"`
		CreatedOn string `json:"created_on"`
		ExpiresOn string `json:"expires_on"`
	} `json:"response"`
}

type Address struct {
	Address      string       `json:"address"`
	Message      string       `json:"message,omitempty"`
	Punycode     string       `json:"punycode,omitempty"`
	SeeAlso      string       `json:"see-also,omitempty"`
	Registration Registration `json:"registration"`
	Expiration   struct {
		Expired    bool `json:"expired"`
		WillExpire bool `json:"will_expire"`
	} `json:"expiration"`
	Owner string `json:"owner,omitempty"`
}

type Addresses struct {
	Request  Request   `json:"request"`
	Response []Address `json:"response"`
}

type AddressAvailability struct {
	Request  Request `json:"request"`
	Response struct {
		Message      string `json:"message"`
		Address      string `json:"address"`
		Available    bool   `json:"available"`
		Availability string `json:"availability"`
	} `json:"response"`
}

type AddressExpiration struct {
	Request  Request `json:"request"`
	Response struct {
		Message string `json:"message,omitempty"`
		Expired bool   `json:"expired,omitempty"`
	} `json:"response"`
}

type AddressInfo struct {
	Request  Request `json:"request"`
	Response struct {
		Address      string `json:"address"`
		Message      string `json:"message"`
		Registration struct {
			Message       string `json:"message"`
			UnixEpochTime string `json:"unix_epoch_time"`
			Iso8601Time   string `json:"iso_8601_time"`
			Rfc2822Time   string `json:"rfc_2822_time"`
			RelativeTime  string `json:"relative_time"`
		} `json:"registration"`
		Expiration struct {
			Message       string `json:"message"`
			Expired       bool   `json:"expired"`
			WillExpire    bool   `json:"will_expire"`
			UnixEpochTime string `json:"unix_epoch_time"`
			Iso8601Time   string `json:"iso_8601_time"`
			Rfc2822Time   string `json:"rfc_2822_time"`
			RelativeTime  string `json:"relative_time"`
		} `json:"expiration"`
		Verification struct {
			Message  string `json:"message"`
			Verified bool   `json:"verified"`
		} `json:"verification"`
		Owner string `json:"owner"`
	} `json:"response"`
}

type AddressDirectory struct {
	Request  Request `json:"request"`
	Response struct {
		Message   string   `json:"message"`
		URL       string   `json:"url"`
		Directory []string `json:"directory"`
	} `json:"response"`
}

type DNSRecord struct {
	ID        *string      `json:"id"`
	Type      *string      `json:"type"`
	Name      *string      `json:"name"`
	Data      *string      `json:"data"`
	Priority  *interface{} `json:"priority"`
	TTL       *string      `json:"ttl"`
	CreatedAt *string      `json:"created_at"`
	UpdatedAt *string      `json:"updated_at"`
}

type DNSRecords struct {
	Request  Request `json:"request"`
	Response struct {
		Message string      `json:"message,omitempty"`
		DNS     []DNSRecord `json:"dns"`
	} `json:"response"`
}

type DNSChangeResponse struct {
	Request  Request `json:"request"`
	Response struct {
		Message  string `json:"message"`
		DataSent struct {
			Type     string  `json:"type"`
			Priority *string `json:"priority"`
			TTL      *string `json:"ttl"`
			Name     string  `json:"name"`
			Content  string  `json:"content"`
		} `json:"data_sent"`
		ResponseReceived struct {
			Data struct {
				ID        int    `json:"id"`
				Name      string `json:"name"`
				Content   string `json:"content"`
				TTL       int    `json:"ttl"`
				Priority  *int   `json:"priority"`
				Type      string `json:"type"`
				CreatedAt string `json:"created_at"`
				UpdatedAt string `json:"updated_at"`
			} `json:"data"`
		} `json:"response_received"`
	} `json:"response"`
}

type PersistantURL struct {
	Request  Request `json:"request"`
	Response struct {
		Message string `json:"message"`
		Purl    struct {
			Name    string      `json:"name"`
			Url     string      `json:"url"`
			Counter interface{} `json:"counter"`
		} `json:"purl"`
	} `json:"response"`
}

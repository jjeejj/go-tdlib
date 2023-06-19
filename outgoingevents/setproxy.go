package outgoingevents

type SetProxy struct {
	Id           int32       `json:"id"`
	Server       string      `json:"server"` // Proxy server IP or Host address.
	Port         int32       `json:"reply_to_message_id"`
	LastUsedDate int32       `json:"last_used_date"` // Point in time (Unix timestamp) when the proxy was last used; 0 if never.
	IsEnabled    bool        `json:"is_enabled"`     // 是否默认开启
	Proxy        interface{} `json:"type"`           // Type of the proxy.
}

func (s SetProxy) Type() string {
	return "proxy"
}

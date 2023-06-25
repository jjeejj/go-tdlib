package incomingevents

import "github.com/jjeejj/go-tdlib/entities"

type Proxy struct {
	Event
	ID           int32                  `json:"id"`             // Unique identifier of the proxy
	Server       string                 `json:"server"`         // Proxy server IP address
	Port         int32                  `json:"port"`           // Proxy server port
	LastUsedDate int32                  `json:"last_used_date"` // Point in time (Unix timestamp) when the proxy was last used; 0 if never
	IsEnabled    bool                   `json:"is_enabled"`     // True, if the proxy is enabled now
	Type         entities.ProxyTypeEnum `json:"type"`           // Type of the proxy
}

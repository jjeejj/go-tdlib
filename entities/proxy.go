package entities

// ProxyTypeEnum Alias for abstract ProxyType 'Sub-Classes', used as constant-enum here
type ProxyTypeEnum string

// ProxyType enums
const (
	ProxyTypeSocks5Type  ProxyTypeEnum = "proxyTypeSocks5"
	ProxyTypeHttpType    ProxyTypeEnum = "proxyTypeHttp"
	ProxyTypeMtprotoType ProxyTypeEnum = "proxyTypeMtproto"
)

type HttpProxy struct {
	Username string `json:"username"`
	Password string `json:"password"`
	HttpOnly bool   `json:"http_only"` // Pass true if the proxy supports only HTTP requests and doesn't support transparent TCP connections via HTTP CONNECT method
}

type Socks5Proxy struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

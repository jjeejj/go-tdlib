package entities

type ProxyType string

const (
	HttpProxyType   ProxyType = "http_proxy_type"
	Socks5ProxyType ProxyType = "socks5_proxy_type"
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

type Proxy struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Type     ProxyType `json:"type"`
}

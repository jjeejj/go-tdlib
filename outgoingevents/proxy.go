package outgoingevents

import "github.com/jjeejj/go-tdlib/entities"

// ReqProxyInfo 请求有关代理得信息
type ReqProxyInfo struct {
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Type     entities.ProxyTypeEnum `json:"type"`
}

// 新增代理
type AddProxy struct {
	Server    string      `json:"server"` // Proxy server IP or Host address.
	Port      int32       `json:"port"`
	IsEnabled bool        `json:"is_enabled"` // 是否默认开启
	Proxy     interface{} `json:"type"`       // Type of the proxy.
}

func (p AddProxy) Type() string {
	return "addProxy"
}

// EnableProxy 开启代理
type EnableProxy struct {
	ProxyId int32 `json:"proxy_id"` // Type of the proxy.
}

func (p EnableProxy) Type() string {
	return "enableProxy"
}

// DisableProxy 关闭当前开启的代理
type DisableProxy struct {
}

func (p DisableProxy) Type() string {
	return "disableProxy"
}

type GetProxies struct {
}

func (p GetProxies) Type() string {
	return "getProxies"
}

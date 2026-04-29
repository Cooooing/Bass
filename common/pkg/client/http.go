package client

import (
	"net"
	"net/http"
	"time"
)

// NewHTTPClient 创建一个单例 HTTP 客户端。
func NewHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			// 连接池：限制总量，防止单个目标打满
			MaxIdleConns:        200,              // 全局空闲连接上限
			MaxIdleConnsPerHost: 50,               // 每个 host 的空闲连接
			MaxConnsPerHost:     100,              // 每个 host 的总连接上限（活跃+空闲）
			IdleConnTimeout:     90 * time.Second, // 空闲连接回收

			// 超时：精细控制每个阶段
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,  // TCP 建连超时
				KeepAlive: 30 * time.Second, // TCP keepalive 探活间隔
			}).DialContext,
			TLSHandshakeTimeout:   5 * time.Second,  // TLS 握手超时
			ResponseHeaderTimeout: 10 * time.Second, // 等待响应头超时
			ExpectContinueTimeout: 1 * time.Second,  // 100-continue 超时

			// 性能
			DisableCompression: false,
			ForceAttemptHTTP2:  true, // 启用 HTTP/2 多路复用
		},
		// 不设全局 Timeout，由调用方通过 context 控制
	}
}

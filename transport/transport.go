package transport

import (
	"crypto/tls"
	"github.com/fabiolb/fabio/config"
	"net"
	"net/http"
)

var (
	cfg *config.Config = &config.Config{}
)

func NewTransport(tlscfg *tls.Config) *http.Transport {
	return &http.Transport{
		ResponseHeaderTimeout: cfg.Proxy.ResponseHeaderTimeout,
		IdleConnTimeout:       cfg.Proxy.IdleConnTimeout,
		MaxIdleConnsPerHost:   cfg.Proxy.MaxConn,
		TLSHandshakeTimeout:   cfg.Proxy.DialTimeout,
		DialContext: (&net.Dialer{
			Timeout:   cfg.Proxy.DialTimeout,
			KeepAlive: cfg.Proxy.KeepAliveTimeout,
		}).DialContext,
		DialTLSContext: (&net.Dialer{
			Timeout:   cfg.Proxy.DialTimeout,
			KeepAlive: cfg.Proxy.KeepAliveTimeout,
		}).DialContext,
		TLSClientConfig: tlscfg,
	}
}

func SetConfig(cfg *config.Config) {
	cfg = cfg
}

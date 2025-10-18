package model

import (
	"google.golang.org/protobuf/types/known/durationpb"
)

type EtcdConf struct {
	Endpoints []string             `json:"endpoints,omitempty"`
	Username  string               `json:"username,omitempty"`
	Password  string               `json:"password,omitempty"`
	Timeout   *durationpb.Duration `json:"timeout,omitempty"`
}

type RedisConf struct {
	Addr         string               `json:"addr,omitempty"`
	Password     string               `json:"password,omitempty"`
	Db           int32                `json:"db,omitempty"`
	ReadTimeout  *durationpb.Duration `json:"readTimeout,omitempty"`
	WriteTimeout *durationpb.Duration `json:"writeTimeout,omitempty"`
}

type RabbitmqConf struct {
	Url            string               `json:"url,omitempty"`
	Heartbeat      *durationpb.Duration `json:"heartbeat,omitempty"`
	DialTimeout    *durationpb.Duration `json:"dialTimeout,omitempty"`
	PrefetchCount  int32                `json:"prefetchCount,omitempty"`
	PrefetchGlobal bool                 `json:"prefetchGlobal,omitempty"`
	DeliveryMode   int32                `json:"deliveryMode,omitempty"`
	AutoAck        bool                 `json:"autoAck,omitempty"`
}

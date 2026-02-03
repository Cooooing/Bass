package model

import (
	"google.golang.org/protobuf/types/known/durationpb"
)

type EtcdConf struct {
	Endpoints            []string             `json:"endpoints,omitempty"`
	Username             string               `json:"username,omitempty"`
	Password             string               `json:"password,omitempty"`
	Timeout              *durationpb.Duration `json:"timeout,omitempty"`
	AutoSyncInterval     *durationpb.Duration `json:"auto_sync_interval,omitempty"`
	DialKeepAliveTime    *durationpb.Duration `json:"dial_keep_alive_time,omitempty"`
	DialKeepAliveTimeout *durationpb.Duration `json:"dial_keep_alive_timeout,omitempty"`
	PermitWithoutStream  bool                 `json:"permit_without_stream,omitempty"`
}

type RedisConf struct {
	Addr            string               `json:"addr,omitempty"`
	Password        string               `json:"password,omitempty"`
	Db              int32                `json:"db,omitempty"`
	DialTimeout     *durationpb.Duration `json:"dial_timeout,omitempty"`
	ReadTimeout     *durationpb.Duration `json:"read_timeout,omitempty"`
	WriteTimeout    *durationpb.Duration `json:"write_timeout,omitempty"`
	PoolSize        int32                `json:"pool_size,omitempty"`
	MinIdleConns    int32                `json:"min_idle_conns,omitempty"`
	PoolTimeout     *durationpb.Duration `json:"pool_timeout,omitempty"`
	ConnMaxIdleTime *durationpb.Duration `json:"conn_max_idle_time,omitempty"`
	ConnMaxLifetime *durationpb.Duration `json:"conn_max_lifetime,omitempty"`
}

type RabbitmqConf struct {
	Url            string               `json:"url,omitempty"`
	Heartbeat      *durationpb.Duration `json:"heartbeat,omitempty"`
	DialTimeout    *durationpb.Duration `json:"dial_timeout,omitempty"`
	PrefetchCount  int32                `json:"prefetch_count,omitempty"`
	PrefetchGlobal bool                 `json:"prefetch_global,omitempty"`
	DeliveryMode   int32                `json:"delivery_mode,omitempty"`
	AutoAck        bool                 `json:"auto_ack,omitempty"`
}

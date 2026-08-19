package client

import (
	"economy/internal/biz/base"
	"fmt"
	"hash/crc32"
	"os"

	"github.com/sony/sonyflake/v2"
)

type SonyflakeTransactionNoGenerator struct {
	sonyflake *sonyflake.Sonyflake
}

func NewTransactionNoGenerator() (base.TransactionNoGenerator, error) {
	sonyflake, err := sonyflake.New(sonyflake.Settings{
		MachineID: func() (int, error) {
			hostname, err := os.Hostname()
			if err != nil {
				return 0, err
			}
			return int(crc32.ChecksumIEEE([]byte(hostname)) % 65536), nil
		},
	})
	if err != nil {
		return nil, err
	}
	return &SonyflakeTransactionNoGenerator{sonyflake: sonyflake}, nil
}

func (g *SonyflakeTransactionNoGenerator) NewTransactionNo() (string, error) {
	id, err := g.sonyflake.NextID()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("eco_%d", id), nil
}

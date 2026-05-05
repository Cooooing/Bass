package main

import (
	"common/api/gen/common"
	"common/pkg/client"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hashicorp/consul/api"
	"google.golang.org/protobuf/types/known/durationpb"
)

var services = []string{"gateway", "user", "content", "notify", "im", "signal", "connector"}

func main() {
	// 初始化 Consul 客户端
	c, f, err := client.NewConsulClient(log.GetLogger(), &common.Consul{
		Address:     fmt.Sprintf("%s:%s", os.Getenv("CONSUL_HOST"), os.Getenv("CONSUL_PORT")),
		Datacenter:  "dc1",
		Token:       os.Getenv("CONSUL_ACL_MASTER_TOKEN"),
		DialTimeout: durationpb.New(5 * time.Second),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer f()

	for _, service := range services {
		// 设定 Consul 中的路径
		consulKey := fmt.Sprintf("config/%s/config.yaml", service)

		// 读取本地原始 YAML 文件
		filePath := filepath.Join("..", "app", service, "configs", "config.yaml")
		rawBytes, err := os.ReadFile(filePath)
		if err != nil {
			log.Errorf("读取文件失败 %s: %v", filePath, err)
			continue
		}

		// 核心步骤：替换环境变量
		// os.ExpandEnv 会寻找字符串里的 ${VAR} 或 $VAR 并用本地环境变量替换
		// 如果环境变量不存在，默认会替换为空字符串
		content := os.Expand(string(rawBytes), func(str string) string {
			parts := strings.SplitN(str, ":", 2)
			key := parts[0]
			val := os.Getenv(key)
			if val == "" && len(parts) > 1 {
				return parts[1]
			}
			return val
		})

		// 将处理后的字符串写入 Consul KV
		p := &api.KVPair{
			Key:   consulKey,
			Value: []byte(content), // 存入替换后的内容
		}

		_, err = c.Client.KV().Put(p, nil)
		if err != nil {
			log.Errorf("上传配置到 Consul 失败 [%s]: %v", consulKey, err)
			continue
		}

		fmt.Printf("服务 [%s] 配置已注入环境变量并上传至: %s\n", service, consulKey)
	}
}

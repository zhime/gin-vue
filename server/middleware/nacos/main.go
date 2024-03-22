package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

func main() {

	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "console1.nacos.io",
			ContextPath: "/nacos",
			Port:        80,
			Scheme:      "http",
		},
		{
			IpAddr:      "console2.nacos.io",
			ContextPath: "/nacos",
			Port:        80,
			Scheme:      "http",
		},
	}

	fmt.Println(clientConfig)
	fmt.Println(serverConfigs)
}

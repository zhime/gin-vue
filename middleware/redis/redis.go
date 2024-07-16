package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	// 创建 Redis 集群客户端
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"192.168.34.33:6379",
			"192.168.34.34:6379",
			"192.168.34.35:6379",
		},
		Password: "password",
	})

	// 测试连接
	err := rdb.Ping(ctx).Err()
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	fmt.Println("连接成功")

	// 设置键值
	//err = rdb.Set(ctx, "key", "value", 0).Err()
	//if err != nil {
	//	fmt.Println("设置键值失败:", err)
	//	return
	//}

	// 获取键值
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("获取键值失败:", err)
		return
	}
	fmt.Println("获取到的值:", val)
}

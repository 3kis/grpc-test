package main

import (
	"context"
	"log"
	"time"

	pb "grpc-test/proto" // 导入生成的 Go proto 文件，现在与模块名一致

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接到 gRPC 服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("无法连接到 gRPC 服务器: %v", err)
	}
	defer conn.Close()

	// 创建 UserService 客户端
	c := pb.NewUserServiceClient(conn)

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用 GetUserInfo RPC (请求1: 张三, Age: 35)
	r, err := c.GetUserInfo(ctx, &pb.UserRequest{Name: "张三", Age: 1})
	if err != nil {
		log.Fatalf("无法获取用户信息 (请求1): %v", err)
	}

	// 打印服务器响应 (请求1)
	log.Printf("收到服务器响应 (请求1):")
	log.Printf("性别: %s", r.GetGender().String())
	log.Printf("消息: %s", r.GetMessage())
	log.Printf("项目: %v", r.GetItems())
	if r.GetJob() != nil {
		log.Printf("职位: %s", r.GetJob().GetItems()) // 打印 Job 信息
	}

	// 调用 GetUserInfo RPC (请求2: 李四, Age: 25)
	r2, err := c.GetUserInfo(ctx, &pb.UserRequest{Name: "李四", Age: 2})
	if err != nil {
		log.Fatalf("无法获取用户信息 (请求2): %v", err)
	}
	log.Printf("\n收到服务器响应 (请求2):")
	log.Printf("性别: %s", r2.GetGender().String())
	log.Printf("消息: %s", r2.GetMessage())
	log.Printf("项目: %v", r2.GetItems())
	if r2.GetJob() != nil {
		log.Printf("职位: %s", r2.GetJob().GetItems()) // 打印 Job 信息
	}

	// 发送一个 nil 请求 (请求3)
	log.Printf("\n发送一个 nil 请求 (请求3)...")
	r3, err := c.GetUserInfo(ctx, nil) // 服务器会收到 age=0
	if err != nil {
		log.Fatalf("无法获取用户信息 (nil 请求): %v", err)
	}
	log.Printf("\n收到服务器对 nil 请求的响应 (请求3):")
	log.Printf("性别: %s", r3.GetGender().String())
	log.Printf("消息: %s", r3.GetMessage())
	log.Printf("项目: %v", r3.GetItems())
	if r3.GetJob() != nil {
		log.Printf("职位: %s", r3.GetJob().GetItems()) // 打印 Job 信息
	}
}

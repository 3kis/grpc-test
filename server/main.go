package main

import (
	"context"
	"log"
	"net"

	pb "grpc-test/proto" // 导入生成的 Go proto 文件，现在与模块名一致

	"google.golang.org/grpc"
)

// server 结构体用于实现 UserService 服务接口
type server struct {
	pb.UnimplementedUserServiceServer // 嵌入 UnimplementedUserServiceServer 以确保向前兼容
}

// GetUserInfo 实现 UserService 服务中的 GetUserInfo 方法
func (s *server) GetUserInfo(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("收到用户请求: Name: %s, Age: %d", in.GetName(), in.GetAge())

	var gender pb.Gender
	var message string
	var items []string
	var job *pb.Job // 声明 Job 字段

	// 根据年龄设置不同的响应
	if in.GetAge() == 0 {
		gender = pb.Gender_GENDER_MALE
		message = "您是一位经验丰富的用户！"
		items = []string{"高级会员", "专属服务"}
		job = &pb.Job{Items: "实习工程师"} // 设置 Job 信息
	} else if in.GetAge() == 1 {
		gender = pb.Gender_GENDER_FEMALE // 这里只是示例，实际业务逻辑会更复杂
		message = "欢迎新用户！"
		items = nil
		job = &pb.Job{Items: "初级工程师"} // 设置 Job 信息
	} else if in.GetAge() == 2 {
		gender = pb.Gender_GENDER_UNSPECIFIED
		message = "您是一位经验丰富的用户！"
		items = []string{}
		job = &pb.Job{Items: "高级工程师"} // 设置 Job 信息
	}

	return &pb.UserResponse{
		Gender:  gender,
		Items:   items,
		Message: message,
		Job:     job, // 包含 Job 字段
	}, nil
}

func main() {
	// 监听 TCP 端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("无法监听端口: %v", err)
	}

	// 创建一个新的 gRPC 服务器实例
	s := grpc.NewServer()

	// 将我们的 server 结构体注册到 gRPC 服务器
	pb.RegisterUserServiceServer(s, &server{})

	log.Printf("服务器正在监听 %v", lis.Addr())
	// 启动 gRPC 服务器并开始处理请求
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

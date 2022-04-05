package service

import (
	"context"
	pb "grpc-go/pb"
	"log"
)

type server struct {
	pb.UnimplementedTestServer
}

func (*server) GrpcTest(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	// Getxxxメソッドも自動に作成されています。
	result := in.Text + "　を受け取ったよ！"
	log.Printf("Received " + in.Text)
	// 受け取ったメッセージを連結したレスポンスを返します。
	return &pb.Response{
		Text: result,
	}, nil
}

//初期化用メソッド
func NewTestServer() *server {
	return &server{}
}

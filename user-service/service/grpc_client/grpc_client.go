package grpcClient

import (
	"fmt"
	pb "github.com/baxromumarov/my-services/user-service/genproto"
	"github.com/baxromumarov/my-services/user-service/config"
	"google.golang.org/grpc"
)

//GrpcClientI ...
type GrpcClientI interface {
	PostSevice() pb.PostServiceClient
}

//GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	postService pb.PostServiceClient
}

//New ...
func New(cfg config.Config) (*GrpcClient, error) {
  connPost, err := grpc.Dial(
    fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
    grpc.WithInsecure())
  if err != nil {
    return nil, fmt.Errorf("post service dial host: %s port: %d",
      cfg.PostServiceHost, cfg.PostServicePort)
  }

  return &GrpcClient{
    cfg:         cfg,
    postService: pb.NewPostServiceClient(connPost),
  }, nil
}

//PostService ...
func (s *GrpcClient) PostSevice() pb.PostServiceClient {
	return s.postService
}

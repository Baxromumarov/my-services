package service

import (
	"context"
	// "fmt"
	"time"

	pb "github.com/baxromumarov/my-services/post-service/genproto"
	l "github.com/baxromumarov/my-services/post-service/pkg/logger"
	cl "github.com/baxromumarov/my-services/post-service/service/grpc_client"
	"github.com/baxromumarov/my-services/post-service/storage"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client cl.GrpcClientI
}

//NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger,client cl.GrpcClientI) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client: client,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	id, err := uuid.NewV4()
	crtime := time.Now()
	
	if err != nil {
		s.logger.Error("Error while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while generating uuid")
	}

	req.Id = id.String()
	req.CreatedAt = crtime.Format(time.RFC3339)
	post, err := s.storage.Post().CreatePost(req)
	if err != nil {
		s.logger.Error("Error while inserting post", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while inserting post")
	}

	return post, nil
}

func (s *PostService) GetByIdPost(ctx context.Context, id *pb.ByIdPost) (*pb.Post, error) {
	post, err := s.storage.Post().GetByIdPost(id.Id)
	if err != nil {
		s.logger.Error("Error while getting post", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting post")
	}

	return post, nil
}

func (s *PostService) GetAllUserPosts(ctx context.Context, req *pb.ByUserIdPost) (*pb.GetUserPosts, error) {
	posts, err := s.storage.Post().GetAllUserPosts(req.UserId)
	if err != nil {
		s.logger.Error("failed get all user posts", l.Error(err))
		return nil, status.Error(codes.Internal, "failed get all user posts")
	}
	// _, err = s.client.UserService().GetById(
	// 	context.Background(),
	// 	&pb.ById{
	// 		Id:req.UserId,
	// 	},
		
	// )
	// if err != nil {
	// 	s.logger.Error("failed get user by id", l.Error(err))
	// 	return nil, status.Error(codes.Internal, "failed get user by id")
	// }
	
	
	return &pb.GetUserPosts{
		Posts: posts,
		
	}, err
}
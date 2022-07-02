package service

import (
	"context"
	"fmt"
	pb "github.com/baxromumarov/my-services/user-service/genproto"
	// "github.com/baxromumarov/my-services/user-service/pkg/logger"
	l "github.com/baxromumarov/my-services/user-service/pkg/logger"
	cl "github.com/baxromumarov/my-services/user-service/service/grpc_client"
	"github.com/baxromumarov/my-services/user-service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  cl.GrpcClientI
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client cl.GrpcClientI) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {

	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("Error while creating user", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while creating user")
	}
	return user, nil
}

func (s *UserService) CreateAd(ctx context.Context, cad *pb.Address) (*pb.Address, error) {

	cred, err := s.storage.User().CreateAd(cad)
	if err != nil {
		s.logger.Error("Error while creating address", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while creating address")
	}
	return cred, nil
}

func (s *UserService) Insert(ctx context.Context, req1 *pb.User) (*pb.User, error) {
	user, err := s.storage.User().Insert(req1)
	if err != nil {
		s.logger.Error("Error while inserting user", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while inserting user")
	}
	if req1.Post != nil {
		for _, post := range req1.Post {
			post.UserId = user.Id
			createdPost, err := s.client.PostSevice().CreatePost(context.Background(), post)
			if err != nil {
				s.logger.Error("Error while inserting post", l.Error(err))
				return nil, status.Error(codes.Internal, "Error while inserting post")
			}
			fmt.Println(createdPost)
		}

	}
	return user, nil

}

func (s *UserService) InsertAd(ctx context.Context, add *pb.Address) (*pb.Address, error) {
	address, err := s.storage.User().InsertAd(add)
	if err != nil {
		s.logger.Error("Error while inserting address", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while inserting address")
	}
	return address, nil
}

func (s *UserService) Delete(ctx context.Context, id *pb.ById) (*pb.UserInfo, error) {
	user, err := s.storage.User().Delete(id)
	if err != nil {
		s.logger.Error("Error while deleting user", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while deleting user")
	}
	return user, nil
}

func (s *UserService) GetById(ctx context.Context, id *pb.ById) (*pb.User, error) {
	user, err := s.storage.User().GetById(id)
	if err != nil {
		s.logger.Error("Error while getting all users", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting all users")
	}
	return user, nil
}

func (s *UserService) GetAll(ctx context.Context, req *pb.Empty) (*pb.UserResp, error) {
	users, err := s.storage.User().GetAll()
	if err != nil {
		s.logger.Error("Error while getting all users1", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting all users1")
	}

	for _, user := range users {
		posts, err := s.client.PostSevice().GetAllUserPosts(
			ctx,
			&pb.ByUserIdPost{
				UserId: user.Id,
			},
		)

		if err != nil {
			s.logger.Error("Error while getting all users", l.Error(err))
			return nil, status.Error(codes.Internal, "Error while getting all users")
		}
		user.Post = posts.Posts

	}
	return &pb.UserResp{
		User: users,
	}, err

}

func (s *UserService) ListUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, count, err := s.storage.User().GetUserList(req.Limit, req.Page)
	if err != nil {
		s.logger.Error("failed while getting all users", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting all users")
	}
	for _, user := range users {
		post, err := s.client.PostSevice().GetAllUserPosts(
			ctx,
			&pb.ByUserIdPost{UserId: user.Id})

		if err != nil {
			s.logger.Error("failed while getting user posts", l.Error(err))
			return nil, status.Error(codes.Internal, "failed while getting user posts")
		}
		user.Post = post.Posts
	}
	return &pb.GetUsersResponse{
		Users: users,
		Count: count,
	}, nil

}

func (s *UserService) UserList(ctx context.Context, req *pb.UserListRequest) (*pb.UserListResponse, error) {
	users, count, err := s.storage.User().UserList(req.Limit, req.Page)

	if err != nil {
		s.logger.Error("failed while getting list of users", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting list of users")
	}

	for _, user := range users {
		posts, err := s.client.PostSevice().GetAllUserPosts(
			ctx,
			&pb.ByUserIdPost{
				UserId: user.Id,
			})

		if err != nil {
			s.logger.Error("failed while getting list of user postd", l.Error(err))
			return nil, status.Error(codes.Internal, "failed while getting list of user posts")
		}

		user.Post = posts.Posts
	}

	return &pb.UserListResponse{
		User:  users,
		Count: count,
	}, nil
}

func (s *UserService) CheckField(ctx context.Context, req *pb.UserCheckRequest) (*pb.UserCheckResponse, error) {

	bl, err := s.storage.User().CheckFeild(req.Field, req.Value)

	if err != nil {
		s.logger.Error("CheckFeild FUNC ERROR", l.Error(err))
		return nil, status.Error(codes.Internal, "CheckFeild FUNC ERROR")
	}

	return &pb.UserCheckResponse{Response: bl}, nil
}

func (s *UserService) GetByEmail(ctx context.Context, req *pb.Email) (*pb.UserData, error) {
	email, err := s.storage.User().GetByEmail(req)
	if err != nil {
		s.logger.Error("Error while getting email information", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting email information")
	}
	return email, nil
}
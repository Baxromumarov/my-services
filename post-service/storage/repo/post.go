package repo

import (
	pb "github.com/baxromumarov/my-services/post-service/genproto"
)

//PostStorageI ...
type PostStorageI interface {
	CreatePost(*pb.Post) (*pb.Post, error)
	GetByIdPost(id string) (*pb.Post, error)
	GetAllUserPosts(Id string) ([]*pb.Post, error)
}

package storage

import (
	pb "github.com/baxromumarov/my-services/api-gateway/genproto"
	"github.com/baxromumarov/my-services/api-gateway/storage/postgres"
	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	User() UserStorageI
	Train() TrainStorageI
}

// UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	CreateAd(*pb.Address) (*pb.Address, error)
	Insert(*pb.User) (*pb.User, error)
	InsertAd(*pb.Address) (*pb.Address, error)
	//Update(id, firstName, lastName *pb.User) (*pb.UserInfo, error)
	Delete(id *pb.ById) (*pb.UserInfo, error)
	GetById(*pb.ById) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	GetUserList(limit, page int64) ([]*pb.User, int64, error)

	UserList(limit, page int64) ([]*pb.User, int64, error)
	CheckFeild(field, value string) (bool, error)
}

type TrainStorageI interface {
	GetJson()
}

type storagePg struct {
	db        *sqlx.DB
	userRepo  UserStorageI
	trainRepo TrainStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:        db,
		userRepo:  postgres.NewUserRepo(db),
		trainRepo: postgres.NewTrainRepo(db),
	}
}

func (s storagePg) Train() TrainStorageI {
	return s.trainRepo
}

func (s storagePg) User() UserStorageI {
	return s.userRepo
}

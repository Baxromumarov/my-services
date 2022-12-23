package repo

type RedisRepositoryStorage interface {
	Set(key, value string) error
	SetWithTTL(key, value string, second int64) error
	Get(key string) (interface{}, error)
}

// // UserStorageI ...
// type UserStorageI interface {
// 	Create(*pb.User) (*pb.User, error)
// 	CreateAd(*pb.Address) (*pb.Address, error)
// 	Insert(*pb.User) (*pb.User, error)
// 	InsertAd(*pb.Address) (*pb.Address, error)
// 	//Update(id, firstName, lastName *pb.User) (*pb.UserInfo, error)
// 	Delete(id *pb.ById) (*pb.UserInfo, error)
// 	GetById(*pb.ById) (*pb.User, error)
// 	GetAll() ([]*pb.User, error)
// 	GetUserList(limit, page int64) ([]*pb.User, int64, error)

// 	UserList(limit, page int64) ([]*pb.User, int64, error)
// 	CheckFeild(field, value string) (bool, error)
// }

// type TrainStorageI interface {
// 	GetJson() (models.JsonFile, error)
// }
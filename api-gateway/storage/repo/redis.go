package response

type RedisRepositoryStorage interface {
	Set(key, value string) error
	SetWithTTL(key, value string, second int64) error
	Get(key string) (interface{}, error)
}

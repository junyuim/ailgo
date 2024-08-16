package core_cache

type ICacheHandler interface {
	Keys(prefix string) ([]string, error)

	Has(key string) (bool, error)

	Get(key string, out any) (bool, error)

	Set(key string, value any, expire int64) error

	Del(key string) (bool, error)
}

package iface

type Cache interface {
	Get(key string) (any, error)
	Set(key string, value any) error
	Delete(key string) error
}

package config

type IConfig interface {
	Key() string
	Validate() error
}

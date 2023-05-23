package utils_gorm

type Config struct {
	Dsn          string
	Echo         bool
	UseApm       bool
	MaxOpenConns int
}

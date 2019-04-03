package config

type Config struct {
	Addr string `json:"address"`
	Db   db     `json:"db"`
}

type db struct {
	Name       string `json:"name"`
	Addr       string `json:"address"`
	User       string `json:"user"`
	Password   string `json:"password"`
	DbPoolSize int    `json:"pool_size"`
}

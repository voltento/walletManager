package config

type Config struct {
	// Address where start listening for ms
	Addr string `json:"address"`
	// Database config
	Db db `json:"db"`
}

type db struct {
	// Database name
	Name string `json:"name"`

	// Database address
	Addr string `json:"address"`

	// User name
	User string `json:"user"`

	// Password
	Password string `json:"password"`

	// Size of database clients pool
	DbPoolSize int `json:"pool_size"`
}
